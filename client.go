package hotcoin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"time"
)

// Client HOTCOIN API客户端
type Client struct {
	config     *Config
	httpClient *http.Client
	signature  *Signature

	// API服务
	Market    *MarketService
	Account   *AccountService
	Trading   *TradingService
	Position  *PositionService
	Common    *CommonService
	WebSocket *WebSocketService
}

// NewClient 创建新的客户端实例
func NewClient(apiKey, secretKey string) *Client {
	config := DefaultConfig()
	config.APIKey = apiKey
	config.SecretKey = secretKey
	return NewClientWithConfig(config)
}

// NewClientWithConfig 使用配置创建客户端
func NewClientWithConfig(config *Config) *Client {
	httpClient := &http.Client{
		Timeout: config.Timeout,
	}

	signature := NewSignature(config.SecretKey)

	client := &Client{
		config:     config,
		httpClient: httpClient,
		signature:  signature,
	}

	// 初始化各个服务
	client.Market = &MarketService{client: client}
	client.Account = &AccountService{client: client}
	client.Trading = &TradingService{client: client}
	client.Position = &PositionService{client: client}
	client.Common = &CommonService{client: client}
	client.WebSocket = &WebSocketService{client: client}

	return client
}

// doRequest 执行HTTP请求
func (c *Client) doRequest(ctx context.Context, method, path string, params map[string]string, body interface{}, needAuth bool) (*Response, error) {
	var reqBody io.Reader

	// 处理请求体
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(jsonData)
	}

	// 构建URL
	var requestURL string
	var err error

	if needAuth && c.config.APIKey != "" && c.config.SecretKey != "" {
		// 需要认证的请求
		requestURL, err = c.signature.BuildAuthURL(method, c.config.BaseURL, path, c.config.APIKey, params)
		if err != nil {
			return nil, fmt.Errorf("build auth URL: %w", err)
		}
	} else {
		// 公开接口请求
		u, err := url.Parse(c.config.BaseURL + path)
		if err != nil {
			return nil, fmt.Errorf("parse URL: %w", err)
		}

		if len(params) > 0 {
			query := u.Query()
			for key, value := range params {
				query.Set(key, value)
			}
			u.RawQuery = query.Encode()
		}

		requestURL = u.String()
	}

	// 创建HTTP请求
	req, err := http.NewRequestWithContext(ctx, method, requestURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "hotcoin-go-sdk/1.0.0")

	if c.config.Debug {
		fmt.Printf("[DEBUG] Request: %s %s\n", method, requestURL)
		if body != nil {
			jsonData, _ := json.Marshal(body)
			fmt.Printf("[DEBUG] Request Body: %s\n", string(jsonData))
		}
	}

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if c.config.Debug {
		fmt.Printf("[DEBUG] Response Status: %d\n", resp.StatusCode)
		fmt.Printf("[DEBUG] Response Body: %s\n", string(respBody))
	}

	// 解析响应
	var response Response
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	// 检查业务错误
	if response.Code != 200 {
		return nil, &ErrorResponse{
			Code: response.Code,
			Msg:  response.Msg,
		}
	}

	return &response, nil
}

// get 发送GET请求
func (c *Client) get(ctx context.Context, path string, params map[string]string, needAuth bool) (*Response, error) {
	return c.doRequest(ctx, "GET", path, params, nil, needAuth)
}

// post 发送POST请求
func (c *Client) post(ctx context.Context, path string, body interface{}, needAuth bool) (*Response, error) {
	return c.doRequest(ctx, "POST", path, nil, body, needAuth)
}

// delete 发送DELETE请求
func (c *Client) delete(ctx context.Context, path string, params map[string]string, needAuth bool) (*Response, error) {
	return c.doRequest(ctx, "DELETE", path, params, nil, needAuth)
}

// SetDebug 设置调试模式
func (c *Client) SetDebug(debug bool) {
	c.config.Debug = debug
}

// SetTimeout 设置请求超时时间
func (c *Client) SetTimeout(timeout time.Duration) {
	c.config.Timeout = timeout
	c.httpClient.Timeout = timeout
}

// GetConfig 获取客户端配置
func (c *Client) GetConfig() *Config {
	return c.config
}
