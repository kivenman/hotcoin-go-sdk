package hotcoin

// Service 基础服务接口
type Service struct {
	client *Client
}

// MarketService 行情服务
type MarketService struct {
	client *Client
}

// AccountService 账户服务
type AccountService struct {
	client *Client
}

// TradingService 交易服务
type TradingService struct {
	client *Client
}

// PositionService 持仓服务
type PositionService struct {
	client *Client
}

// CommonService 通用服务
type CommonService struct {
	client *Client
}
