package hotcoin

// Contract 合约信息
type Contract struct {
	Code                     string  `json:"code"`                     // 合约代码
	Base                     string  `json:"base"`                     // 基础货币
	BaseDisplayName          string  `json:"baseDisplayName"`          // 基础货币显示名称
	Quote                    string  `json:"quote"`                    // 计价货币
	QuoteDisplayName         string  `json:"quoteDisplayName"`         // 计价货币显示名称
	IndexBase                string  `json:"indexBase"`                // 指数基础货币
	IndexBaseDisplayName     string  `json:"indexBaseDisplayName"`     // 指数基础货币显示名称
	IndexBaseAppLogo         string  `json:"indexBaseAppLogo"`         // APP Logo
	IndexBaseWebLogo         string  `json:"indexBaseWebLogo"`         // Web Logo
	Direction                int     `json:"direction"`                // 方向 0:正向合约,1:反向合约
	Price                    string  `json:"price"`                    // 最新价
	MarkPrice                string  `json:"markPrice"`                // 标记价格
	IndexPrice               string  `json:"indexPrice"`               // 指数价格
	High                     string  `json:"high"`                     // 最高价
	Low                      string  `json:"low"`                      // 最低价
	Amount24                 string  `json:"amount24"`                 // 24小时成交张数
	Size24                   string  `json:"size24"`                   // 24小时成交价值
	Fluctuation              string  `json:"fluctuation"`              // 涨跌幅
	TotalPosition            string  `json:"totalPosition"`            // 持仓量
	Fund                     string  `json:"fund"`                     // 资金费率
	UnitAmount               float64 `json:"unitAmount"`               // 一张合约对应的quote面值
	MinTradeUnit             float64 `json:"minTradeUnit"`             // 最小交易单位
	MinTradeDigit            int     `json:"minTradeDigit"`            // 基础货币最小交易小数位
	MinQuoteDigit            int     `json:"minQuoteDigit"`            // 计价货币最小交易小数位
	MarketPriceDigit         int     `json:"marketPriceDigit"`         // 市价小数位
	MaxLever                 int     `json:"maxLever"`                 // 最大杠杆
	Env                      int     `json:"env"`                      // 是否测试盘 0:线上盘,1:测试盘
	TradeStatus              int     `json:"tradeStatus"`              // 交易状态
	GuaranteedStopLossRate   string  `json:"guaranteedStopLossRate"`   // 保证止损费率
	GuaranteedStopLossStatus int     `json:"guaranteedStopLossStatus"` // 保证止损状态
	LiquidationTime          int64   `json:"liquidationTime"`          // 清算时间
	NextLiquidationInterval  int     `json:"nextLiquidationInterval"`  // 下次清算间隔
	OpenTradeTime            string  `json:"openTradeTime"`            // 开放交易时间
	PreDeliveryPrice         string  `json:"preDeliveryPrice"`         // 预交割价格
}

// KlineData K线数据
type KlineData struct {
	Timestamp int64  `json:"timestamp"` // 时间戳
	Open      string `json:"open"`      // 开盘价
	High      string `json:"high"`      // 最高价
	Low       string `json:"low"`       // 最低价
	Close     string `json:"close"`     // 收盘价
	Volume    string `json:"volume"`    // 成交量
}

// DepthData 深度数据
type DepthData struct {
	Bids [][]string `json:"bids"` // 买盘 [价格, 数量]
	Asks [][]string `json:"asks"` // 卖盘 [价格, 数量]
}

// TradeData 交易数据
type TradeData struct {
	ID        int64  `json:"id"`        // 交易ID
	Price     string `json:"price"`     // 成交价格
	Amount    string `json:"amount"`    // 成交数量
	Side      string `json:"side"`      // 交易方向
	Timestamp int64  `json:"timestamp"` // 成交时间
}

// IndexPriceComponent 指数价格成分
type IndexPriceComponent struct {
	Symbol    string `json:"symbol"`    // 交易对
	Price     string `json:"price"`     // 价格
	Weight    string `json:"weight"`    // 权重
	Exchange  string `json:"exchange"`  // 交易所
	Timestamp int64  `json:"timestamp"` // 时间戳
}

// FundingRate 资金费率
type FundingRate struct {
	Symbol      string `json:"symbol"`      // 交易对
	FundingRate string `json:"fundingRate"` // 资金费率
	FundingTime int64  `json:"fundingTime"` // 资金费用时间
	MarkPrice   string `json:"markPrice"`   // 标记价格
	IndexPrice  string `json:"indexPrice"`  // 指数价格
}

// TickerData 行情ticker数据（基于HOTCOIN API实际响应格式）
type TickerData struct {
	TickerID       string  `json:"code"`          // ticker标识符
	BaseCurrency   string  `json:"base"`          // 基础货币
	TargetCurrency string  `json:"quote"`         // 目标货币
	LastPrice      string  `json:"price"`         // 最新成交价格
	Amount24       string  `json:"amount24"`      // 24小时成交张数
	Size24         string  `json:"size24"`        // 24小时成交价值
	Bid            string  `json:"bid"`           // 当前最高买入价
	Ask            string  `json:"ask"`           // 当前最低卖出价
	High           string  `json:"high"`          // 过去24小时最高成交价格
	Low            string  `json:"low"`           // 过去24小时最低成交价格
	Fluctuation    string  `json:"fluctuation"`   // 涨跌幅
	TotalPosition  string  `json:"totalPosition"` // 持仓量
	FundingRate    string  `json:"fund"`          // 资金费率
	MarkPrice      string  `json:"markPrice"`     // 标记价格
	IndexPrice     string  `json:"indexPrice"`    // 指数价格
	Direction      int     `json:"direction"`     // 方向 0:正向合约,1:反向合约
	Env            int     `json:"env"`           // 是否测试盘 0:线上盘,1:测试盘
	TradeStatus    int     `json:"tradeStatus"`   // 交易状态
	UnitAmount     float64 `json:"unitAmount"`    // 一张合约对应的quote面值
	MinTradeUnit   float64 `json:"minTradeUnit"`  // 最小交易单位
	MaxLever       int     `json:"maxLever"`      // 最大杠杆
}

// GeckoContract Gecko格式的合约信息
type GeckoContract struct {
	Ticker         string  `json:"ticker"`          // 交易对标识
	BaseCurrency   string  `json:"base_currency"`   // 基础货币
	QuoteCurrency  string  `json:"quote_currency"`  // 计价货币
	LastPrice      float64 `json:"last_price"`      // 最新价格
	BaseVolume     float64 `json:"base_volume"`     // 基础货币成交量
	QuoteVolume    float64 `json:"quote_volume"`    // 计价货币成交量
	Bid            float64 `json:"bid"`             // 买一价
	Ask            float64 `json:"ask"`             // 卖一价
	High           float64 `json:"high"`            // 最高价
	Low            float64 `json:"low"`             // 最低价
	ProductType    string  `json:"product_type"`    // 产品类型
	OpenInterest   float64 `json:"open_interest"`   // 持仓量
	IndexPrice     float64 `json:"index_price"`     // 指数价格
	IndexCurrency  string  `json:"index_currency"`  // 指数货币
	StartTimestamp int64   `json:"start_timestamp"` // 开始时间戳
	EndTimestamp   int64   `json:"end_timestamp"`   // 结束时间戳
	FundingRate    float64 `json:"funding_rate"`    // 资金费率
	ContractType   string  `json:"contract_type"`   // 合约类型
}
