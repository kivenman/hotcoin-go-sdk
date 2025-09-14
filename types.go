package hotcoin

import (
	"time"
)

// Config SDK配置
type Config struct {
	APIKey    string        // API访问密钥
	SecretKey string        // 签名密钥
	BaseURL   string        // API基础URL
	Timeout   time.Duration // 请求超时时间
	Debug     bool          // 是否开启调试模式
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		BaseURL: "https://api-ct.hotcoin.fit",
		Timeout: 30 * time.Second,
		Debug:   false,
	}
}

// Response 通用响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ErrorResponse) Error() string {
	return e.Msg
}

// OrderSide 订单方向
type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

// OrderType 订单类型
type OrderType string

const (
	OrderTypeMarket OrderType = "market"
	OrderTypeLimit  OrderType = "limit"
)

// OrderStatus 订单状态
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusFilled    OrderStatus = "filled"
	OrderStatusCancelled OrderStatus = "cancelled"
	OrderStatusPartial   OrderStatus = "partial"
)

// PositionSide 持仓方向
type PositionSide string

const (
	PositionSideLong  PositionSide = "long"
	PositionSideShort PositionSide = "short"
)

// MarginMode 保证金模式
type MarginMode string

const (
	MarginModeIsolated MarginMode = "isolated" // 逐仓
	MarginModeCrossed  MarginMode = "crossed"  // 全仓
)

// ContractDirection 合约方向
type ContractDirection int

const (
	ContractDirectionForward ContractDirection = 0 // 正向合约
	ContractDirectionReverse ContractDirection = 1 // 反向合约
)

// Environment 环境标识
type Environment int

const (
	EnvironmentProduction Environment = 0 // 线上环境
	EnvironmentTest       Environment = 1 // 测试环境
)
