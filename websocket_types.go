package hotcoin

// WebSocketMessage WebSocket消息结构
type WebSocketMessage struct {
	ID       string      `json:"id"`       // 消息ID
	Status   string      `json:"status"`   // 状态
	Subbed   string      `json:"subbed"`   // 订阅主题
	Unsubbed string      `json:"unsubbed"` // 取消订阅主题
	Ping     int64       `json:"ping"`     // ping时间戳
	Pong     int64       `json:"pong"`     // pong时间戳
	Rep      string      `json:"rep"`      // 响应请求标识
	Ch       string      `json:"ch"`       // 频道
	Ts       int64       `json:"ts"`       // 时间戳
	Tick     interface{} `json:"tick"`     // 数据
	Data     interface{} `json:"data"`     // 数据
	Op       string      `json:"op"`       // 操作类型
	ErrCode  int         `json:"err-code"` // 错误码
	ErrMsg   string      `json:"err-msg"`  // 错误信息
}

// SubscribeRequest 订阅请求
type SubscribeRequest struct {
	Sub string `json:"sub"` // 订阅主题
	ID  string `json:"id"`  // 请求ID
}

// UnsubscribeRequest 取消订阅请求
type UnsubscribeRequest struct {
	Unsub string `json:"unsub"` // 取消订阅主题
	ID    string `json:"id"`    // 请求ID
}

// AuthRequest 认证请求
type AuthRequest struct {
	Op               string `json:"op"`   // 操作类型，固定为auth
	Type             string `json:"type"` // 认证类型，api
	AccessKeyID      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
}

// WSKlineData WebSocket K线数据
type WSKlineData struct {
	ID     int64  `json:"id"`     // K线ID
	Open   string `json:"open"`   // 开盘价
	Close  string `json:"close"`  // 收盘价
	High   string `json:"high"`   // 最高价
	Low    string `json:"low"`    // 最低价
	Amount string `json:"amount"` // 成交量
	Vol    string `json:"vol"`    // 成交额
	Count  int    `json:"count"`  // 成交笔数
}

// WSDepthData WebSocket深度数据
type WSDepthData struct {
	Bids    [][]string `json:"bids"`    // 买盘
	Asks    [][]string `json:"asks"`    // 卖盘
	Version int64      `json:"version"` // 版本号
	Ts      int64      `json:"ts"`      // 时间戳
}

// WSTradeData WebSocket交易数据
type WSTradeData struct {
	ID   int64 `json:"id"` // 交易ID
	Ts   int64 `json:"ts"` // 时间戳
	Data []struct {
		Amount    string `json:"amount"`    // 成交量
		Ts        int64  `json:"ts"`        // 时间戳
		ID        int64  `json:"id"`        // 成交ID
		Price     string `json:"price"`     // 成交价格
		Direction string `json:"direction"` // 主动成交方向
	} `json:"data"`
}

// WSTickerData WebSocket行情数据
type WSTickerData struct {
	ID     int64    `json:"id"`     // 消息id
	Open   string   `json:"open"`   // 开盘价
	Close  string   `json:"close"`  // 收盘价，当K线为最晚的一根时，是最新成交价
	High   string   `json:"high"`   // 最高价
	Low    string   `json:"low"`    // 最低价
	Amount string   `json:"amount"` // 成交量
	Vol    string   `json:"vol"`    // 成交额
	Count  int      `json:"count"`  // 成交笔数
	Bid    []string `json:"bid"`    // [买1价,买1量]
	Ask    []string `json:"ask"`    // [卖1价,卖1量]
}

// WSOrderData WebSocket订单数据
type WSOrderData struct {
	Symbol         string `json:"symbol"`           // 交易对
	ContractCode   string `json:"contract_code"`    // 合约代码
	ContractType   string `json:"contract_type"`    // 合约类型
	Volume         string `json:"volume"`           // 委托数量
	Price          string `json:"price"`            // 委托价格
	OrderPriceType string `json:"order_price_type"` // 订单报价类型
	Direction      string `json:"direction"`        // 买卖方向
	Offset         string `json:"offset"`           // 开平方向
	Status         int    `json:"status"`           // 订单状态
	LeverRate      int    `json:"lever_rate"`       // 杠杆倍数
	OrderID        int64  `json:"order_id"`         // 订单ID
	OrderIDStr     string `json:"order_id_str"`     // 字符串格式的订单ID
	ClientOrderID  int64  `json:"client_order_id"`  // 客户订单ID
	OrderSource    string `json:"order_source"`     // 订单来源
	OrderType      int    `json:"order_type"`       // 订单类型
	CreatedAt      int64  `json:"created_at"`       // 订单创建时间
	TradeVolume    string `json:"trade_volume"`     // 成交数量
	TradeTurnover  string `json:"trade_turnover"`   // 成交总金额
	Fee            string `json:"fee"`              // 手续费
	TradeAvgPrice  string `json:"trade_avg_price"`  // 成交均价
	MarginFrozen   string `json:"margin_frozen"`    // 冻结保证金
	Profit         string `json:"profit"`           // 收益
}

// WSPositionData WebSocket持仓数据
type WSPositionData struct {
	Symbol         string `json:"symbol"`          // 交易对
	ContractCode   string `json:"contract_code"`   // 合约代码
	ContractType   string `json:"contract_type"`   // 合约类型
	Volume         string `json:"volume"`          // 持仓数量
	Available      string `json:"available"`       // 可平仓数量
	Frozen         string `json:"frozen"`          // 冻结数量
	CostOpen       string `json:"cost_open"`       // 开仓均价
	CostHold       string `json:"cost_hold"`       // 持仓均价
	ProfitUnreal   string `json:"profit_unreal"`   // 未实现盈亏
	ProfitRate     string `json:"profit_rate"`     // 收益率
	Profit         string `json:"profit"`          // 收益
	PositionMargin string `json:"position_margin"` // 持仓保证金
	LeverRate      int    `json:"lever_rate"`      // 杠杆倍数
	Direction      string `json:"direction"`       // 仓位方向
	LastPrice      string `json:"last_price"`      // 最新价格
}

// WSAccountData WebSocket账户数据
type WSAccountData struct {
	Symbol            string `json:"symbol"`             // 保证金币种
	MarginBalance     string `json:"margin_balance"`     // 账户权益
	MarginStatic      string `json:"margin_static"`      // 静态权益
	MarginPosition    string `json:"margin_position"`    // 持仓保证金
	MarginFrozen      string `json:"margin_frozen"`      // 冻结保证金
	MarginAvailable   string `json:"margin_available"`   // 可用保证金
	ProfitReal        string `json:"profit_real"`        // 已实现盈亏
	ProfitUnreal      string `json:"profit_unreal"`      // 未实现盈亏
	WithdrawAvailable string `json:"withdraw_available"` // 可提取数量
	RiskRate          string `json:"risk_rate"`          // 保证金率
	LiquidationPrice  string `json:"liquidation_price"`  // 预估强平价
	LeverRate         int    `json:"lever_rate"`         // 杠杆倍数
	AdjustFactor      string `json:"adjust_factor"`      // 调整系数
}

// EventHandler 事件处理器
type EventHandler func(message *WebSocketMessage)

// ErrorHandler 错误处理器
type ErrorHandler func(err error)

// ConnectionHandler 连接处理器
type ConnectionHandler func()

// WSConfig WebSocket配置
type WSConfig struct {
	URL               string // WebSocket URL
	EnableHeartbeat   bool   // 是否启用心跳
	HeartbeatInterval int    // 心跳间隔(秒)
}

// DefaultWSConfig 默认WebSocket配置
func DefaultWSConfig() *WSConfig {
	return &WSConfig{
		URL:               "wss://api-ct.hotcoin.fit/linear-swap-ws",
		EnableHeartbeat:   true,
		HeartbeatInterval: 20,
	}
}
