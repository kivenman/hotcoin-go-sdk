package hotcoin

// SystemStatus 系统状态
type SystemStatus struct {
	Symbol    string `json:"symbol"`    // 交易对
	Status    int    `json:"status"`    // 系统状态 1-正常 2-系统维护
	Heartbeat int64  `json:"heartbeat"` // 系统心跳时间戳
}

// ServerTime 服务器时间
type ServerTime struct {
	Timestamp int64 `json:"timestamp"` // 服务器时间戳
}

// APIInfo API信息
type APIInfo struct {
	APIKey      string   `json:"api_key"`      // API密钥
	Permissions []string `json:"permissions"`  // 权限列表
	IPWhitelist []string `json:"ip_whitelist"` // IP白名单
	Created     int64    `json:"created"`      // 创建时间
	Updated     int64    `json:"updated"`      // 更新时间
	Status      int      `json:"status"`       // 状态
}

// RiskLimit 风控限制
type RiskLimit struct {
	Symbol        string `json:"symbol"`          // 交易对
	RiskLimitType string `json:"risk_limit_type"` // 风控类型
	RiskLimit     string `json:"risk_limit"`      // 风控限制值
}

// SettlementRecord 结算记录
type SettlementRecord struct {
	Symbol         string `json:"symbol"`          // 交易对
	SettlementTime int64  `json:"settlement_time"` // 结算时间
	ClampPrice     string `json:"clamp_price"`     // 交割价格
	SettlementType string `json:"settlement_type"` // 结算类型
}

// InsuranceFund 保险基金
type InsuranceFund struct {
	Symbol string `json:"symbol"` // 交易对
	Tick   struct {
		Symbol    string `json:"symbol"`    // 交易对
		Amount    string `json:"amount"`    // 保险基金余额
		Timestamp int64  `json:"timestamp"` // 时间戳
	} `json:"tick"`
}

// HistoricalSettlement 历史结算记录
type HistoricalSettlement struct {
	Symbol         string `json:"symbol"`          // 交易对
	SettlementTime int64  `json:"settlement_time"` // 结算时间
	ClampPrice     string `json:"clamp_price"`     // 交割价格
	SettlementType string `json:"settlement_type"` // 结算类型
	PairValue      string `json:"pair_value"`      // 合约面值
}

// LiquidationOrder 强平订单
type LiquidationOrder struct {
	QueryID      int64  `json:"query_id"`      // 查询ID
	Symbol       string `json:"symbol"`        // 交易对
	ContractCode string `json:"contract_code"` // 合约代码
	Direction    string `json:"direction"`     // 强平方向
	Offset       string `json:"offset"`        // 开平方向
	Volume       string `json:"volume"`        // 强平数量
	Price        string `json:"price"`         // 强平价格
	CreatedAt    int64  `json:"created_at"`    // 强平时间
}

// ElitePositionRatio 精英持仓多空比
type ElitePositionRatio struct {
	Symbol     string `json:"symbol"`      // 交易对
	LongRatio  string `json:"long_ratio"`  // 多仓比例
	ShortRatio string `json:"short_ratio"` // 空仓比例
	Timestamp  int64  `json:"timestamp"`   // 时间戳
}

// EliteAccountRatio 精英账户多空比
type EliteAccountRatio struct {
	Symbol       string `json:"symbol"`        // 交易对
	LongAccount  string `json:"long_account"`  // 做多账户比例
	ShortAccount string `json:"short_account"` // 做空账户比例
	Timestamp    int64  `json:"timestamp"`     // 时间戳
}

// ContractElement 合约要素
type ContractElement struct {
	ContractCode   string `json:"contract_code"`   // 合约代码
	ContractType   string `json:"contract_type"`   // 合约类型
	ContractSize   string `json:"contract_size"`   // 合约面值
	PriceTick      string `json:"price_tick"`      // 最小变动价位
	DeliveryTime   string `json:"delivery_time"`   // 交割时间
	CreateDate     string `json:"create_date"`     // 上市日期
	ContractStatus int    `json:"contract_status"` // 合约状态
	SettlementTime string `json:"settlement_time"` // 结算时间
}

// MasterSubTransferRecord 母子账户划转记录
type MasterSubTransferRecord struct {
	ID           int64  `json:"id"`            // 划转ID
	Symbol       string `json:"symbol"`        // 币种
	Amount       string `json:"amount"`        // 划转数量
	TransferType int    `json:"transfer_type"` // 划转类型 34-转出到子账户 35-从子账户转入
	MasterUID    string `json:"master_uid"`    // 母账户UID
	SubUID       string `json:"sub_uid"`       // 子账户UID
	Ts           int64  `json:"ts"`            // 划转时间
}
