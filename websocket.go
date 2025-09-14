package hotcoin

import (
	"compress/gzip"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketService WebSocket服务
type WebSocketService struct {
	client      *Client
	conn        *websocket.Conn
	config      *WSConfig
	isAuth      bool
	isConnected bool
	mutex       sync.RWMutex

	// 事件处理器
	onConnected    ConnectionHandler
	onDisconnected ConnectionHandler
	onError        ErrorHandler
	onMessage      EventHandler

	// 订阅管理
	subscriptions map[string]bool
	subMutex      sync.RWMutex

	// 控制通道
	stopChan chan struct{}
	done     chan struct{}
}

// NewWebSocketService 创建WebSocket服务
func NewWebSocketService(client *Client) *WebSocketService {
	return &WebSocketService{
		client:        client,
		config:        DefaultWSConfig(),
		subscriptions: make(map[string]bool),
		stopChan:      make(chan struct{}),
		done:          make(chan struct{}),
	}
}

// SetConfig 设置WebSocket配置
func (ws *WebSocketService) SetConfig(config *WSConfig) {
	ws.config = config
}

// OnConnected 设置连接成功回调
func (ws *WebSocketService) OnConnected(handler ConnectionHandler) {
	ws.onConnected = handler
}

// OnDisconnected 设置断开连接回调
func (ws *WebSocketService) OnDisconnected(handler ConnectionHandler) {
	ws.onDisconnected = handler
}

// OnError 设置错误回调
func (ws *WebSocketService) OnError(handler ErrorHandler) {
	ws.onError = handler
}

// OnMessage 设置消息回调
func (ws *WebSocketService) OnMessage(handler EventHandler) {
	ws.onMessage = handler
}

// Connect 连接WebSocket
func (ws *WebSocketService) Connect() error {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	if ws.isConnected {
		return fmt.Errorf("already connected")
	}

	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	conn, _, err := dialer.Dial(ws.config.URL, nil)
	if err != nil {
		return fmt.Errorf("websocket dial failed: %w", err)
	}

	ws.conn = conn
	ws.isConnected = true

	// 启动消息处理协程
	go ws.readMessages()

	// 启动心跳协程
	if ws.config.EnableHeartbeat {
		go ws.heartbeat()
	}

	if ws.onConnected != nil {
		ws.onConnected()
	}

	return nil
}

// Disconnect 断开连接
func (ws *WebSocketService) Disconnect() error {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	if !ws.isConnected {
		return nil
	}

	close(ws.stopChan)
	<-ws.done

	err := ws.conn.Close()
	ws.conn = nil
	ws.isConnected = false
	ws.isAuth = false

	// 清空订阅
	ws.subMutex.Lock()
	ws.subscriptions = make(map[string]bool)
	ws.subMutex.Unlock()

	if ws.onDisconnected != nil {
		ws.onDisconnected()
	}

	return err
}

// IsConnected 检查是否已连接
func (ws *WebSocketService) IsConnected() bool {
	ws.mutex.RLock()
	defer ws.mutex.RUnlock()
	return ws.isConnected
}

// Auth 认证
func (ws *WebSocketService) Auth() error {
	if !ws.IsConnected() {
		return fmt.Errorf("not connected")
	}

	if ws.client.config.APIKey == "" || ws.client.config.SecretKey == "" {
		return fmt.Errorf("api key and secret key are required for auth")
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")

	// 构建签名字符串
	signString := "GET\n" + "api-ct.hotcoin.fit\n" + "/api/v1/perpetual/notification\n" +
		"AccessKeyId=" + ws.client.config.APIKey +
		"&SignatureMethod=HmacSHA256" +
		"&SignatureVersion=2" +
		"&Timestamp=" + url.QueryEscape(timestamp)

	// 生成签名
	h := hmac.New(sha256.New, []byte(ws.client.config.SecretKey))
	h.Write([]byte(signString))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	authReq := AuthRequest{
		Op:               "auth",
		Type:             "api",
		AccessKeyID:      ws.client.config.APIKey,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
		Timestamp:        timestamp,
		Signature:        signature,
	}

	return ws.sendMessage(authReq)
}

// Subscribe 订阅主题
func (ws *WebSocketService) Subscribe(topic string) error {
	if !ws.IsConnected() {
		return fmt.Errorf("not connected")
	}

	ws.subMutex.Lock()
	if ws.subscriptions[topic] {
		ws.subMutex.Unlock()
		return fmt.Errorf("already subscribed to %s", topic)
	}
	ws.subscriptions[topic] = true
	ws.subMutex.Unlock()

	req := SubscribeRequest{
		Sub: topic,
		ID:  fmt.Sprintf("sub_%d", time.Now().UnixNano()),
	}

	return ws.sendMessage(req)
}

// Unsubscribe 取消订阅
func (ws *WebSocketService) Unsubscribe(topic string) error {
	if !ws.IsConnected() {
		return fmt.Errorf("not connected")
	}

	ws.subMutex.Lock()
	delete(ws.subscriptions, topic)
	ws.subMutex.Unlock()

	req := UnsubscribeRequest{
		Unsub: topic,
		ID:    fmt.Sprintf("unsub_%d", time.Now().UnixNano()),
	}

	return ws.sendMessage(req)
}

// SubscribeKline 订阅K线数据
func (ws *WebSocketService) SubscribeKline(symbol, period string) error {
	topic := fmt.Sprintf("market.%s.kline.%s", symbol, period)
	return ws.Subscribe(topic)
}

// SubscribeDepth 订阅深度数据
func (ws *WebSocketService) SubscribeDepth(symbol, depthType string) error {
	topic := fmt.Sprintf("market.%s.depth.%s", symbol, depthType)
	return ws.Subscribe(topic)
}

// SubscribeTrade 订阅交易数据
func (ws *WebSocketService) SubscribeTrade(symbol string) error {
	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	return ws.Subscribe(topic)
}

// SubscribeTicker 订阅行情数据
func (ws *WebSocketService) SubscribeTicker(symbol string) error {
	topic := fmt.Sprintf("market.%s.detail", symbol)
	return ws.Subscribe(topic)
}

// SubscribeOrders 订阅订单推送
func (ws *WebSocketService) SubscribeOrders(symbol string) error {
	if !ws.isAuth {
		return fmt.Errorf("authentication required")
	}
	topic := fmt.Sprintf("orders.%s", symbol)
	return ws.Subscribe(topic)
}

// SubscribePositions 订阅持仓推送
func (ws *WebSocketService) SubscribePositions(symbol string) error {
	if !ws.isAuth {
		return fmt.Errorf("authentication required")
	}
	topic := fmt.Sprintf("positions.%s", symbol)
	return ws.Subscribe(topic)
}

// SubscribeAccount 订阅账户推送
func (ws *WebSocketService) SubscribeAccount(symbol string) error {
	if !ws.isAuth {
		return fmt.Errorf("authentication required")
	}
	topic := fmt.Sprintf("accounts.%s", symbol)
	return ws.Subscribe(topic)
}

// sendMessage 发送消息
func (ws *WebSocketService) sendMessage(message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal message failed: %w", err)
	}

	ws.mutex.RLock()
	defer ws.mutex.RUnlock()

	if !ws.isConnected || ws.conn == nil {
		return fmt.Errorf("connection not available")
	}

	return ws.conn.WriteMessage(websocket.TextMessage, data)
}

// readMessages 读取消息
func (ws *WebSocketService) readMessages() {
	defer func() {
		close(ws.done)
	}()

	for {
		select {
		case <-ws.stopChan:
			return
		default:
			_, data, err := ws.conn.ReadMessage()
			if err != nil {
				if ws.onError != nil {
					ws.onError(fmt.Errorf("read message failed: %w", err))
				}
				return
			}

			// 解压gzip数据
			if len(data) > 2 && data[0] == 0x1f && data[1] == 0x8b {
				reader, err := gzip.NewReader(strings.NewReader(string(data)))
				if err != nil {
					if ws.onError != nil {
						ws.onError(fmt.Errorf("gzip decompress failed: %w", err))
					}
					continue
				}

				decompressed, err := io.ReadAll(reader)
				reader.Close()
				if err != nil {
					if ws.onError != nil {
						ws.onError(fmt.Errorf("read decompressed data failed: %w", err))
					}
					continue
				}
				data = decompressed
			}

			var message WebSocketMessage
			if err := json.Unmarshal(data, &message); err != nil {
				if ws.onError != nil {
					ws.onError(fmt.Errorf("unmarshal message failed: %w", err))
				}
				continue
			}

			ws.handleMessage(&message)
		}
	}
}

// handleMessage 处理消息
func (ws *WebSocketService) handleMessage(message *WebSocketMessage) {
	// 处理ping消息
	if message.Ping > 0 {
		pong := map[string]int64{"pong": message.Ping}
		if err := ws.sendMessage(pong); err != nil {
			if ws.onError != nil {
				ws.onError(fmt.Errorf("send pong failed: %w", err))
			}
		}
		return
	}

	// 处理认证响应
	if message.Op == "auth" {
		if message.ErrCode == 0 {
			ws.isAuth = true
			log.Println("WebSocket authentication successful")
		} else {
			if ws.onError != nil {
				ws.onError(fmt.Errorf("authentication failed: %s", message.ErrMsg))
			}
		}
		return
	}

	// 处理订阅确认
	if message.Subbed != "" {
		log.Printf("Subscribed to: %s", message.Subbed)
		return
	}

	// 处理取消订阅确认
	if message.Unsubbed != "" {
		log.Printf("Unsubscribed from: %s", message.Unsubbed)
		return
	}

	// 回调用户处理器
	if ws.onMessage != nil {
		ws.onMessage(message)
	}
}

// heartbeat 心跳
func (ws *WebSocketService) heartbeat() {
	ticker := time.NewTicker(time.Duration(ws.config.HeartbeatInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ws.stopChan:
			return
		case <-ticker.C:
			ping := map[string]int64{"ping": time.Now().UnixMilli()}
			if err := ws.sendMessage(ping); err != nil {
				if ws.onError != nil {
					ws.onError(fmt.Errorf("send ping failed: %w", err))
				}
				return
			}
		}
	}
}
