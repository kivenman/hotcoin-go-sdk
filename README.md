# HOTCOIN Go SDK

这是 HOTCOIN 永续合约 API 的官方 Go 语言 SDK，提供了完整的API接口封装，支持行情查询、账户管理、交易下单、持仓管理、WebSocket实时数据推送等功能。

## 功能特性

- ✅ 完整的REST API支持
- ✅ WebSocket实时数据推送
- ✅ HmacSHA256签名认证
- ✅ 自动重连和心跳机制
- ✅ 详细的错误处理
- ✅ 类型安全的数据结构
- ✅ 丰富的示例代码
- ✅ 完整的单元测试

## 支持的接口

### 行情接口
- 获取合约列表
- 获取K线数据
- 获取深度信息
- 获取交易记录
- 获取指数价格
- 获取资金费率
- 获取24小时行情统计

### 账户接口
- 获取账户信息
- 获取账户余额
- 获取持仓信息
- 设置杠杆倍数
- 设置保证金模式
- 获取费率信息
- 获取财务记录

### 交易接口
- 下单
- 批量下单
- 撤销订单
- 批量撤销
- 获取订单信息
- 获取订单明细
- 获取当前委托
- 获取历史委托
- 计划委托下单
- 获取成交记录

### 持仓接口
- 获取持仓信息
- 获取子账户持仓
- 一键平仓

### 通用接口
- 获取系统状态
- 获取服务器时间
- 获取合约要素
- 获取保险基金
- 获取强平订单
- 母子账户划转

### WebSocket接口
- 行情数据推送
- 订单状态推送
- 持仓变化推送
- 账户变化推送

## 安装

```bash
go get github.com/kivenman/hotcoin-go-sdk
```

## 快速开始

### 基础用法

```go
package main

import (
    "fmt"
    "log"
    
    hotcoin "github.com/kivenman/hotcoin-go-sdk"
)

func main() {
    // 创建客户端
    client := hotcoin.NewClient("your_api_key", "your_secret_key")
    
    // 启用调试模式
    client.SetDebug(true)
    
    // 获取合约列表
    contracts, err := client.Market.GetContracts("")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("获取到 %d 个合约\n", len(contracts))
}
```

### 行情数据

```go
// 获取K线数据
klines, err := client.Market.GetKline("BTC-USDT", "1min", 100)
if err != nil {
    log.Fatal(err)
}

// 获取深度数据
depth, err := client.Market.GetDepth("BTC-USDT", "step0")
if err != nil {
    log.Fatal(err)
}

// 获取24小时行情统计
tickers, err := client.Market.GetTicker("BTC-USDT")
if err != nil {
    log.Fatal(err)
}
```

### 账户管理

```go
// 获取账户信息
accountInfo, err := client.Account.GetAccountInfo("USDT")
if err != nil {
    log.Fatal(err)
}

// 获取持仓信息
positions, err := client.Position.GetPositions("")
if err != nil {
    log.Fatal(err)
}

// 设置杠杆倍数
err = client.Account.SetLeverage("BTC-USDT", 10)
if err != nil {
    log.Fatal(err)
}
```

### 交易操作

```go
// 下限价单
orderReq := &hotcoin.OrderPlaceRequest{
    Symbol:         "BTC-USDT",
    ContractType:   "swap",
    Direction:      "buy",
    Offset:         "open",
    Volume:         "1",
    Price:          "50000",
    LeverRate:      10,
    OrderPriceType: "limit",
}

order, err := client.Trading.PlaceOrder(orderReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("下单成功，订单ID: %s\n", order.OrderID)

// 撤销订单
cancelReq := &hotcoin.OrderCancelRequest{
    Symbol:  "BTC-USDT",
    OrderID: order.OrderID,
}

_, err = client.Trading.CancelOrder(cancelReq)
if err != nil {
    log.Fatal(err)
}
```

### WebSocket实时数据

```go
// 创建WebSocket客户端
ws := hotcoin.NewWebSocketService(client)

// 设置事件处理器
ws.OnConnected(func() {
    fmt.Println("WebSocket连接成功")
})

ws.OnMessage(func(message *hotcoin.WebSocketMessage) {
    fmt.Printf("收到消息: %s\n", message.Ch)
})

ws.OnError(func(err error) {
    fmt.Printf("WebSocket错误: %v\n", err)
})

// 连接WebSocket
err := ws.Connect()
if err != nil {
    log.Fatal(err)
}

// 订阅K线数据
err = ws.SubscribeKline("BTC-USDT", "1min")
if err != nil {
    log.Fatal(err)
}

// 订阅深度数据
err = ws.SubscribeDepth("BTC-USDT", "step0")
if err != nil {
    log.Fatal(err)
}

// 认证（用于订阅私有数据）
err = ws.Auth()
if err != nil {
    log.Fatal(err)
}

// 订阅订单推送
err = ws.SubscribeOrders("BTC-USDT")
if err != nil {
    log.Fatal(err)
}
```

## 配置选项

```go
// 使用自定义配置
config := &hotcoin.Config{
    APIKey:    "your_api_key",
    SecretKey: "your_secret_key",
    BaseURL:   "https://api-ct.hotcoin.fit",
    Timeout:   30 * time.Second,
    Debug:     true,
}

client := hotcoin.NewClientWithConfig(config)
```

## 错误处理

SDK提供了详细的错误信息：

```go
orders, err := client.Trading.GetOpenOrders("BTC-USDT", 1, 20)
if err != nil {
    if apiErr, ok := err.(*hotcoin.ErrorResponse); ok {
        fmt.Printf("API错误: Code=%d, Msg=%s\n", apiErr.Code, apiErr.Msg)
    } else {
        fmt.Printf("其他错误: %v\n", err)
    }
    return
}
```

## 数据类型

### 订单方向
```go
hotcoin.OrderSideBuy    // 买入
hotcoin.OrderSideSell   // 卖出
```

### 订单类型
```go
hotcoin.OrderTypeLimit   // 限价单
hotcoin.OrderTypeMarket  // 市价单
```

### 订单状态
```go
hotcoin.OrderStatusPending   // 待成交
hotcoin.OrderStatusFilled    // 已成交
hotcoin.OrderStatusCancelled // 已撤销
hotcoin.OrderStatusPartial   // 部分成交
```

### 持仓方向
```go
hotcoin.PositionSideLong  // 多头
hotcoin.PositionSideShort // 空头
```

### 保证金模式
```go
hotcoin.MarginModeIsolated // 逐仓
hotcoin.MarginModeCrossed  // 全仓
```

## 示例代码

更多详细示例请查看 [examples](./examples) 目录：

- [基础示例](./examples/basic/main.go) - 基本API使用
- [WebSocket示例](./examples/websocket/main.go) - WebSocket实时数据
- [交易示例](./examples/trading/main.go) - 完整的交易流程

## API文档

完整的API文档请参考：[HOTCOIN API文档](https://docs.hotcoin.fit)

## 注意事项

1. **API密钥安全**: 请妥善保管您的API密钥，不要在代码中硬编码或提交到版本控制系统
2. **测试环境**: 建议在测试环境充分测试后再在生产环境使用
3. **频率限制**: 请遵守API的频率限制，避免过于频繁的请求
4. **错误处理**: 请做好错误处理，特别是网络异常和API错误
5. **WebSocket重连**: WebSocket连接可能会断开，请实现重连机制

## 贡献

欢迎提交Issue和Pull Request来改进这个SDK。

## 许可证

MIT License

## 联系我们

如有问题或建议，请联系：
- 官方网站：https://www.hotcoin.fit
- API文档：https://docs.hotcoin.fit
- 技术支持：api@hotcoin.fit