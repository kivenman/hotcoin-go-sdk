package hotcoin

// AccountBalance 账户余额
type AccountBalance struct {
	Symbol            string `json:"symbol"`            // 币种符号
	MarginBalance     string `json:"marginBalance"`     // 账户权益
	MarginStatic      string `json:"marginStatic"`      // 静态权益
	MarginPosition    string `json:"marginPosition"`    // 持仓保证金
	MarginFrozen      string `json:"marginFrozen"`      // 冻结保证金
	MarginAvailable   string `json:"marginAvailable"`   // 可用保证金
	ProfitReal        string `json:"profitReal"`        // 已实现盈亏
	ProfitUnreal      string `json:"profitUnreal"`      // 未实现盈亏
	RiskRate          string `json:"riskRate"`          // 保证金率
	LiquidationPrice  string `json:"liquidationPrice"`  // 强平价格
	WithdrawAvailable string `json:"withdrawAvailable"` // 可提取数量
	LeverRate         int    `json:"leverRate"`         // 杠杆倍数
	AdjustFactor      string `json:"adjustFactor"`      // 调整系数
}

// AccountInfo 账户信息
type AccountInfo struct {
	Symbol              string           `json:"symbol"`                // 保证金币种
	MarginBalance       string           `json:"margin_balance"`        // 账户权益
	MarginStatic        string           `json:"margin_static"`         // 静态权益
	MarginPosition      string           `json:"margin_position"`       // 持仓保证金
	MarginFrozen        string           `json:"margin_frozen"`         // 冻结保证金
	MarginAvailable     string           `json:"margin_available"`      // 可用保证金
	ProfitReal          string           `json:"profit_real"`           // 已实现盈亏
	ProfitUnreal        string           `json:"profit_unreal"`         // 未实现盈亏
	WithdrawAvailable   string           `json:"withdraw_available"`    // 可提取数量
	RiskRate            string           `json:"risk_rate"`             // 保证金率
	LiquidationPrice    string           `json:"liquidation_price"`     // 预估强平价
	LeverRate           int              `json:"lever_rate"`            // 杠杆倍数
	AdjustFactor        string           `json:"adjust_factor"`         // 调整系数
	MarginMode          string           `json:"margin_mode"`           // 保证金模式
	PositionMode        string           `json:"position_mode"`         // 持仓模式
	TransferProfitRatio string           `json:"transfer_profit_ratio"` // 止盈转移比例
	TransferLossRatio   string           `json:"transfer_loss_ratio"`   // 止损转移比例
	MarginAccount       string           `json:"margin_account"`        // 保证金账户
	Positions           []PositionDetail `json:"positions"`             // 持仓详情
}

// PositionDetail 持仓详情
type PositionDetail struct {
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
	MarginPosition string `json:"margin_position"` // 持仓保证金
	PositionMargin string `json:"position_margin"` // 持仓保证金
	Direction      string `json:"direction"`       // 仓位方向 buy-买 sell-卖
	LastPrice      string `json:"last_price"`      // 最新价
	LeverRate      int    `json:"lever_rate"`      // 杠杆倍数
}

// TransferRecord 划转记录
type TransferRecord struct {
	OrderID    string `json:"order_id"`    // 订单ID
	Currency   string `json:"currency"`    // 币种
	Amount     string `json:"amount"`      // 划转数量
	Type       string `json:"type"`        // 类型 pro-to-futures futures-to-pro
	Status     string `json:"status"`      // 状态
	CreateTime int64  `json:"create_time"` // 创建时间
	UpdateTime int64  `json:"update_time"` // 更新时间
}

// FeeRate 费率信息
type FeeRate struct {
	Symbol     string `json:"symbol"`      // 交易对
	OpenMaker  string `json:"open_maker"`  // 开仓maker费率
	OpenTaker  string `json:"open_taker"`  // 开仓taker费率
	CloseMaker string `json:"close_maker"` // 平仓maker费率
	CloseTaker string `json:"close_taker"` // 平仓taker费率
}

// TransferLimit 划转限额
type TransferLimit struct {
	Currency    string `json:"currency"`     // 币种
	MinAmount   string `json:"min_amount"`   // 最小划转金额
	MaxAmount   string `json:"max_amount"`   // 最大划转金额
	TransferIn  bool   `json:"transfer_in"`  // 是否支持转入
	TransferOut bool   `json:"transfer_out"` // 是否支持转出
}

// FinancialRecord 财务记录
type FinancialRecord struct {
	ID           int64  `json:"id"`            // 记录ID
	Type         int    `json:"type"`          // 记录类型
	Amount       string `json:"amount"`        // 金额
	Ts           int64  `json:"ts"`            // 时间戳
	Symbol       string `json:"symbol"`        // 交易对
	ContractCode string `json:"contract_code"` // 合约代码
}

// LeverageInfo 杠杆信息
type LeverageInfo struct {
	Symbol       string `json:"symbol"`        // 交易对
	LeverageList []int  `json:"leverage_list"` // 可用杠杆倍数列表
	CurrentLever int    `json:"current_lever"` // 当前杠杆倍数
}

// PositionLimitInfo 持仓限制信息
type PositionLimitInfo struct {
	Symbol    string `json:"symbol"`     // 交易对
	BuyLimit  string `json:"buy_limit"`  // 多仓限制
	SellLimit string `json:"sell_limit"` // 空仓限制
}

// MarginModeInfo 保证金模式信息
type MarginModeInfo struct {
	Symbol     string `json:"symbol"`      // 交易对
	MarginMode string `json:"margin_mode"` // 保证金模式 isolated-逐仓 crossed-全仓
}

// AssetValuation 资产估值
type AssetValuation struct {
	ValuationAsset string `json:"valuation_asset"` // 估值资产
	Balance        string `json:"balance"`         // 余额
}
