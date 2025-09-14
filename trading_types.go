package hotcoin

// Order 订单信息
type Order struct {
	OrderID        string `json:"order_id"`         // 订单ID
	OrderIDStr     string `json:"order_id_str"`     // 订单ID字符串
	Symbol         string `json:"symbol"`           // 交易对
	ContractCode   string `json:"contract_code"`    // 合约代码
	ContractType   string `json:"contract_type"`    // 合约类型
	Direction      string `json:"direction"`        // 买卖方向 buy-买 sell-卖
	Offset         string `json:"offset"`           // 开平方向 open-开 close-平
	Volume         string `json:"volume"`           // 委托数量
	Price          string `json:"price"`            // 委托价格
	CreateDate     int64  `json:"create_date"`      // 创建时间
	OrderSource    string `json:"order_source"`     // 订单来源
	OrderPriceType string `json:"order_price_type"` // 订单类型
	MarginFrozen   string `json:"margin_frozen"`    // 冻结保证金
	Profit         string `json:"profit"`           // 收益
	Instrument     string `json:"instrument"`       // 合约标识
	OrderType      int    `json:"order_type"`       // 订单类型 1-报单 2-撤单 3-强平 4-交割
	Status         int    `json:"status"`           // 订单状态
	LeverRate      int    `json:"lever_rate"`       // 杠杆倍数
	Fee            string `json:"fee"`              // 手续费
	FeeAsset       string `json:"fee_asset"`        // 手续费币种
	CanceledAt     int64  `json:"canceled_at"`      // 撤销时间
	TradeVolume    string `json:"trade_volume"`     // 成交数量
	TradeTurnover  string `json:"trade_turnover"`   // 成交金额
	TradeAvgPrice  string `json:"trade_avg_price"`  // 成交均价
}

// OrderPlaceRequest 下单请求
type OrderPlaceRequest struct {
	Symbol         string `json:"symbol"`           // 交易对
	ContractType   string `json:"contract_type"`    // 合约类型
	ContractCode   string `json:"contract_code"`    // 合约代码
	ClientOrderID  string `json:"client_order_id"`  // 客户自定义订单ID
	Price          string `json:"price"`            // 价格
	Volume         string `json:"volume"`           // 数量
	Direction      string `json:"direction"`        // 买卖方向
	Offset         string `json:"offset"`           // 开平方向
	LeverRate      int    `json:"lever_rate"`       // 杠杆倍数
	OrderPriceType string `json:"order_price_type"` // 订单类型
}

// OrderPlaceResponse 下单响应
type OrderPlaceResponse struct {
	OrderID       string `json:"order_id"`        // 订单ID
	OrderIDStr    string `json:"order_id_str"`    // 订单ID字符串
	ClientOrderID string `json:"client_order_id"` // 客户订单ID
}

// OrderCancelRequest 撤单请求
type OrderCancelRequest struct {
	OrderID       string `json:"order_id"`        // 订单ID，多个用逗号分隔
	ClientOrderID string `json:"client_order_id"` // 客户自定义订单ID，多个用逗号分隔
	Symbol        string `json:"symbol"`          // 交易对
}

// OrderCancelResponse 撤单响应
type OrderCancelResponse struct {
	Successes []string `json:"successes"` // 成功的订单ID列表
	Errors    []struct {
		OrderID string `json:"order_id"` // 订单ID
		ErrCode int    `json:"err_code"` // 错误码
		ErrMsg  string `json:"err_msg"`  // 错误信息
	} `json:"errors"` // 失败的订单列表
}

// OrderDetail 订单详情
type OrderDetail struct {
	Order
	Trades []TradeDetail `json:"trades"` // 成交明细
}

// TradeDetail 成交明细
type TradeDetail struct {
	TradeID       string `json:"trade_id"`       // 成交ID
	TradeVolume   string `json:"trade_volume"`   // 成交数量
	TradePrice    string `json:"trade_price"`    // 成交价格
	TradeFee      string `json:"trade_fee"`      // 成交手续费
	TradeTurnover string `json:"trade_turnover"` // 成交金额
	CreatedAt     int64  `json:"created_at"`     // 成交时间
	Role          string `json:"role"`           // 成交角色 maker/taker
}

// BatchOrderRequest 批量下单请求
type BatchOrderRequest struct {
	Orders []OrderPlaceRequest `json:"orders"` // 订单列表，最多10个
}

// BatchOrderResponse 批量下单响应
type BatchOrderResponse struct {
	Successes []OrderPlaceResponse `json:"successes"` // 成功的订单
	Errors    []struct {
		Index   int    `json:"index"`    // 订单索引
		ErrCode int    `json:"err_code"` // 错误码
		ErrMsg  string `json:"err_msg"`  // 错误信息
	} `json:"errors"` // 失败的订单
}

// OrderQueryRequest 查询订单请求
type OrderQueryRequest struct {
	Symbol        string `json:"symbol"`          // 交易对
	OrderID       string `json:"order_id"`        // 订单ID
	ClientOrderID string `json:"client_order_id"` // 客户订单ID
	OrderType     int    `json:"order_type"`      // 订单类型
	Status        int    `json:"status"`          // 订单状态
	CreateDate    int    `json:"create_date"`     // 创建日期
	PageIndex     int    `json:"page_index"`      // 页码
	PageSize      int    `json:"page_size"`       // 每页大小
}

// PlanOrder 计划委托订单
type PlanOrder struct {
	OrderID        string `json:"order_id"`         // 订单ID
	OrderIDStr     string `json:"order_id_str"`     // 订单ID字符串
	Symbol         string `json:"symbol"`           // 交易对
	ContractCode   string `json:"contract_code"`    // 合约代码
	ContractType   string `json:"contract_type"`    // 合约类型
	TriggerType    string `json:"trigger_type"`     // 触发类型
	Volume         string `json:"volume"`           // 委托数量
	OrderType      int    `json:"order_type"`       // 订单类型
	Direction      string `json:"direction"`        // 买卖方向
	Offset         string `json:"offset"`           // 开平方向
	LeverRate      int    `json:"lever_rate"`       // 杠杆倍数
	OrderPrice     string `json:"order_price"`      // 委托价格
	OrderPriceType string `json:"order_price_type"` // 订单价格类型
	TriggerPrice   string `json:"trigger_price"`    // 触发价格
	CreatedAt      int64  `json:"created_at"`       // 创建时间
	OrderSource    string `json:"order_source"`     // 订单来源
	Status         int    `json:"status"`           // 订单状态
	OrderOrigType  int    `json:"order_orig_type"`  // 订单原始类型
}

// PlanOrderRequest 计划委托下单请求
type PlanOrderRequest struct {
	Symbol         string `json:"symbol"`           // 交易对
	ContractType   string `json:"contract_type"`    // 合约类型
	ContractCode   string `json:"contract_code"`    // 合约代码
	TriggerType    string `json:"trigger_type"`     // 触发类型 ge-大于等于 le-小于等于
	TriggerPrice   string `json:"trigger_price"`    // 触发价格
	OrderPrice     string `json:"order_price"`      // 委托价格
	OrderPriceType string `json:"order_price_type"` // 订单类型
	Volume         string `json:"volume"`           // 数量
	Direction      string `json:"direction"`        // 买卖方向
	Offset         string `json:"offset"`           // 开平方向
	LeverRate      int    `json:"lever_rate"`       // 杠杆倍数
}

// StopOrderRequest 止盈止损订单请求
type StopOrderRequest struct {
	Symbol           string `json:"symbol"`              // 交易对
	ContractCode     string `json:"contract_code"`       // 合约代码
	ContractType     string `json:"contract_type"`       // 合约类型
	Direction        string `json:"direction"`           // 买卖方向
	Volume           string `json:"volume"`              // 数量
	TpTriggerPrice   string `json:"tp_trigger_price"`    // 止盈触发价格
	TpOrderPrice     string `json:"tp_order_price"`      // 止盈委托价格
	TpOrderPriceType string `json:"tp_order_price_type"` // 止盈订单类型
	SlTriggerPrice   string `json:"sl_trigger_price"`    // 止损触发价格
	SlOrderPrice     string `json:"sl_order_price"`      // 止损委托价格
	SlOrderPriceType string `json:"sl_order_price_type"` // 止损订单类型
}

// MatchResult 撮合结果
type MatchResult struct {
	Symbol    string `json:"symbol"`    // 交易对
	OrderID   string `json:"order_id"`  // 订单ID
	TradeID   string `json:"trade_id"`  // 成交ID
	Volume    string `json:"volume"`    // 成交数量
	Price     string `json:"price"`     // 成交价格
	Timestamp int64  `json:"timestamp"` // 成交时间
}
