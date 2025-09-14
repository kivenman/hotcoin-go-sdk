package hotcoin

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 行情接口相关方法

// GetContracts 获取合约列表
// symbol: 交易对符号，可选，如果不传则返回所有合约
func (m *MarketService) GetContracts(symbol string) ([]Contract, error) {
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}

	var result []Contract
	resp, err := m.client.get(context.Background(), "/api/v1/perpetual/public", params, false)
	if err != nil {
		return nil, err
	}

	// 直接解析到目标结构体
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetKline 获取K线数据
// symbol: 交易对符号
// period: K线周期，支持: 1min, 5min, 15min, 30min, 1hour, 4hour, 1day, 1week, 1mon
// size: 获取数量，默认150，最大2000
func (m *MarketService) GetKline(symbol, period string, size int) ([]KlineData, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if period == "" {
		return nil, fmt.Errorf("period is required")
	}

	// 构建正确的API路径，contractCode从symbol中提取（去掉-转换为小写）
	contractCode := strings.ToLower(strings.Replace(symbol, "-", "", -1))
	path := fmt.Sprintf("/api/v1/perpetual/public/%s/candles", contractCode)

	// 设置查询参数
	params := map[string]string{
		"kline": period, // 新API使用kline参数
	}
	if size > 0 {
		params["size"] = strconv.Itoa(size)
	}

	resp, err := m.client.get(context.Background(), path, params, false)
	if err != nil {
		return nil, err
	}

	// 检查响应数据结构
	if resp.Data == nil {
		return nil, fmt.Errorf("empty response data")
	}

	// 尝试解析为二维数组格式（根据文档示例）
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}

	var klineArray [][]interface{}
	err = json.Unmarshal(dataBytes, &klineArray)
	if err != nil {
		return nil, fmt.Errorf("unmarshal kline data: %w", err)
	}

	// 转换为KlineData结构
	var klines []KlineData
	for _, item := range klineArray {
		if len(item) >= 6 {
			// 解析各个字段
			timestamp, _ := item[0].(float64) // 时间戳
			low, _ := item[1].(string)        // 最低价
			high, _ := item[2].(string)       // 最高价
			open, _ := item[3].(string)       // 开盘价
			close, _ := item[4].(string)      // 收盘价
			volume, _ := item[5].(string)     // 成交量

			kline := KlineData{
				Timestamp: int64(timestamp),
				Low:       low,
				High:      high,
				Open:      open,
				Close:     close,
				Volume:    volume,
			}
			klines = append(klines, kline)
		}
	}

	return klines, nil
}

// GetDepth 获取深度信息
// symbol: 交易对符号
// depthType: 深度类型，支持: step0, step1, step2, step3, step4, step5
func (m *MarketService) GetDepth(symbol, depthType string) (*DepthData, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	// 构建正确的API路径
	contractCode := strings.ToLower(strings.Replace(symbol, "-", "", -1))
	path := fmt.Sprintf("/api/v1/perpetual/public/products/%s/orderbook", contractCode)

	// 设置查询参数
	params := map[string]string{}
	if depthType != "" {
		params["size"] = "20" // 深度数量，根据需要调整
	}

	resp, err := m.client.get(context.Background(), path, params, false)
	if err != nil {
		return nil, err
	}

	// 检查响应数据
	if resp.Data == nil {
		return nil, fmt.Errorf("empty response data")
	}

	// 直接解析到目标结构体（新API直接返回深度数据）
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}

	var response DepthData
	err = json.Unmarshal(dataBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal depth data: %w", err)
	}

	return &response, nil
}

// GetTrades 获取交易记录
// symbol: 交易对符号
// size: 获取数量，默认1，最大2000
func (m *MarketService) GetTrades(symbol string, size int) ([]TradeData, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}
	if size > 0 {
		params["size"] = strconv.Itoa(size)
	}

	var response struct {
		Status string    `json:"status"`
		Ch     string    `json:"ch"`
		Ts     int64     `json:"ts"`
		Tick   TradeData `json:"tick"`
	}

	// 构建正确的API路径
	contractCode := strings.ToLower(strings.Replace(symbol, "-", "", -1))
	path := fmt.Sprintf("/api/v1/perpetual/public/%s/fills", contractCode)

	resp, err := m.client.get(context.Background(), path, params, false)
	if err != nil {
		return nil, err
	}

	// 直接解析到目标结构体
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API error: status=%s", response.Status)
	}

	return []TradeData{response.Tick}, nil
}

// GetIndexPrice 获取指数价格
// symbol: 交易对符号，可选，如果不传则返回所有
func (m *MarketService) GetIndexPrice(symbol string) ([]IndexPriceComponent, error) {
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}

	var response struct {
		Status string                `json:"status"`
		Data   []IndexPriceComponent `json:"data"`
		Ts     int64                 `json:"ts"`
	}

	resp, err := m.client.get(context.Background(), "/api/v1/perpetual/public/index-price", params, false)
	if err != nil {
		return nil, err
	}

	// 直接解析到目标结构体
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API error: status=%s", response.Status)
	}

	return response.Data, nil
}

// GetFundingRate 获取资金费率
// symbol: 交易对符号
func (m *MarketService) GetFundingRate(symbol string) (*FundingRate, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}

	var response struct {
		Status string      `json:"status"`
		Data   FundingRate `json:"data"`
		Ts     int64       `json:"ts"`
	}

	// 构建正确的API路径
	contractCode := strings.ToLower(strings.Replace(symbol, "-", "", -1))
	path := fmt.Sprintf("/api/v1/perpetual/public/products/%s/funding-rate", contractCode)

	resp, err := m.client.get(context.Background(), path, params, false)
	if err != nil {
		return nil, err
	}

	// 直接解析到目标结构体
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API error: status=%s", response.Status)
	}

	return &response.Data, nil
}

// GetHistoricalFundingRate 获取历史资金费率
// symbol: 交易对符号
// pageIndex: 页码，从1开始
// pageSize: 每页数量，默认20，最大50
func (m *MarketService) GetHistoricalFundingRate(symbol string, pageIndex, pageSize int) ([]FundingRate, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}
	if pageIndex > 0 {
		params["page_index"] = strconv.Itoa(pageIndex)
	}
	if pageSize > 0 {
		params["page_size"] = strconv.Itoa(pageSize)
	}

	var response struct {
		Status string        `json:"status"`
		Data   []FundingRate `json:"data"`
		Ts     int64         `json:"ts"`
	}

	// 构建正确的API路径
	contractCode := strings.ToLower(strings.Replace(symbol, "-", "", -1))
	path := fmt.Sprintf("/api/v1/perpetual/public/products/%s/funding-rate/history", contractCode)

	resp, err := m.client.get(context.Background(), path, params, false)
	if err != nil {
		return nil, err
	}

	// 直接解析到目标结构体
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API error: status=%s", response.Status)
	}

	return response.Data, nil
}

// GetTicker 获取24小时行情统计
// symbol: 交易对符号，可选，如果不传则返回所有
func (m *MarketService) GetTicker(symbol string) ([]TickerData, error) {
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}

	// 24小时行情统计从产品列表接口获取
	resp, err := m.client.get(context.Background(), "/api/v1/perpetual/public", params, false)
	if err != nil {
		return nil, err
	}

	// 检查响应数据
	if resp.Data == nil {
		return nil, fmt.Errorf("empty response data")
	}

	// 新API直接返回数组，不包装在status结构中
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}

	var tickers []TickerData
	err = json.Unmarshal(dataBytes, &tickers)
	if err != nil {
		return nil, fmt.Errorf("unmarshal ticker data: %w", err)
	}

	return tickers, nil
}

// GetHistoricalKline 获取历史K线数据（支持更大范围查询）
// symbol: 交易对符号
// period: K线周期
// from: 开始时间戳（秒）
// to: 结束时间戳（秒）
func (m *MarketService) GetHistoricalKline(symbol, period string, from, to int64) ([]KlineData, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if period == "" {
		return nil, fmt.Errorf("period is required")
	}

	params := map[string]string{
		"symbol": symbol,
		"period": period,
	}
	if from > 0 {
		params["from"] = strconv.FormatInt(from, 10)
	}
	if to > 0 {
		params["to"] = strconv.FormatInt(to, 10)
	}

	var response struct {
		Status string      `json:"status"`
		Data   []KlineData `json:"data"`
		Ts     int64       `json:"ts"`
	}

	// 构建正确的API路径
	contractCode := strings.ToLower(strings.Replace(symbol, "-", "", -1))
	path := fmt.Sprintf("/api/v1/perpetual/public/%s/candles/history", contractCode)

	resp, err := m.client.get(context.Background(), path, params, false)
	if err != nil {
		return nil, err
	}

	// 直接解析到目标结构体
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API error: status=%s", response.Status)
	}

	return response.Data, nil
}

// GetGeckoContracts 获取Gecko格式的合约信息（用于CoinGecko等第三方平台）
func (m *MarketService) GetGeckoContracts() ([]GeckoContract, error) {
	var result []GeckoContract
	resp, err := m.client.get(context.Background(), "/api/v1/perpetual/public/contracts", nil, false)
	if err != nil {
		return nil, err
	}

	// 直接解析到目标结构体
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetBatchTicker 批量获取行情ticker数据
// symbols: 交易对符号列表，多个用逗号分隔
func (m *MarketService) GetBatchTicker(symbols []string) ([]TickerData, error) {
	if len(symbols) == 0 {
		return m.GetTicker("") // 获取所有
	}

	params := map[string]string{
		"symbol": strings.Join(symbols, ","),
	}

	var response struct {
		Status string       `json:"status"`
		Data   []TickerData `json:"data"`
		Ts     int64        `json:"ts"`
	}

	// 批量获取行情数据也使用产品列表接口
	resp, err := m.client.get(context.Background(), "/api/v1/perpetual/public", params, false)
	if err != nil {
		return nil, err
	}

	// 直接解析到目标结构体
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API error: status=%s", response.Status)
	}

	return response.Data, nil
}
