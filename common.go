package hotcoin

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// CommonService 通用接口服务

// GetServerTime 获取服务器时间
func (c *CommonService) GetServerTime() (*ServerTime, error) {
	resp, err := c.client.get(context.Background(), "/api/v1/timestamp", nil, false)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string     `json:"status"`
		Data   ServerTime `json:"data"`
		Ts     int64      `json:"ts"`
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

// GetSystemStatus 获取系统状态
// symbol: 交易对，可选
func (c *CommonService) GetSystemStatus(symbol string) ([]SystemStatus, error) {
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}

	resp, err := c.client.get(context.Background(), "/api/v1/perpetual/public/api-state", params, false)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string         `json:"status"`
		Data   []SystemStatus `json:"data"`
		Ts     int64          `json:"ts"`
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

// GetContractElements 获取合约要素
// symbol: 交易对，可选
func (c *CommonService) GetContractElements(symbol string) ([]ContractElement, error) {
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}

	resp, err := c.client.get(context.Background(), "/api/v1/perpetual/public/query-elements", params, false)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string            `json:"status"`
		Data   []ContractElement `json:"data"`
		Ts     int64             `json:"ts"`
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

// GetInsuranceFund 获取保险基金历史数据
// symbol: 交易对
func (c *CommonService) GetInsuranceFund(symbol string) (*InsuranceFund, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := c.client.get(context.Background(), "/api/v1/perpetual/public/insurance-fund", params, false)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string        `json:"status"`
		Data   InsuranceFund `json:"data"`
		Ts     int64         `json:"ts"`
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

// GetLiquidationOrders 获取强平订单
// symbol: 交易对
// tradeType: 交易类型 0-全部 5-强平多 6-强平空
// pageIndex: 页码，从0开始
// pageSize: 每页大小，最大50
func (c *CommonService) GetLiquidationOrders(symbol string, tradeType, pageIndex, pageSize int) ([]LiquidationOrder, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}
	if tradeType >= 0 {
		params["trade_type"] = strconv.Itoa(tradeType)
	}
	if pageIndex >= 0 {
		params["page_index"] = strconv.Itoa(pageIndex)
	}
	if pageSize > 0 {
		params["page_size"] = strconv.Itoa(pageSize)
	}

	resp, err := c.client.get(context.Background(), "/api/v1/perpetual/public/liquidation-orders", params, false)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string             `json:"status"`
		Data   []LiquidationOrder `json:"data"`
		Ts     int64              `json:"ts"`
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

// GetHistoricalSettlement 获取平台历史结算记录
// symbol: 交易对
// pageIndex: 页码，从1开始
// pageSize: 每页大小，最大50
func (c *CommonService) GetHistoricalSettlement(symbol string, pageIndex, pageSize int) ([]HistoricalSettlement, error) {
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

	resp, err := c.client.get(context.Background(), "/api/v1/perpetual/public/settlement-records", params, false)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string                 `json:"status"`
		Data   []HistoricalSettlement `json:"data"`
		Ts     int64                  `json:"ts"`
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

// GetElitePositionRatio 获取精英账户多空持仓对比-持仓量
// symbol: 交易对
// period: 周期 5min, 15min, 30min, 1hour, 4hour, 1day
func (c *CommonService) GetElitePositionRatio(symbol, period string) ([]ElitePositionRatio, error) {
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

	resp, err := c.client.get(context.Background(), "/api/v1/perpetual/public/elite-position-ratio", params, false)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string               `json:"status"`
		Data   []ElitePositionRatio `json:"data"`
		Ts     int64                `json:"ts"`
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

// GetEliteAccountRatio 获取精英账户多空持仓对比-账户数
// symbol: 交易对
// period: 周期 5min, 15min, 30min, 1hour, 4hour, 1day
func (c *CommonService) GetEliteAccountRatio(symbol, period string) ([]EliteAccountRatio, error) {
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

	resp, err := c.client.get(context.Background(), "/api/v1/perpetual/public/elite-account-ratio", params, false)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string              `json:"status"`
		Data   []EliteAccountRatio `json:"data"`
		Ts     int64               `json:"ts"`
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

// GetAPIInfo 获取用户API指标禁用信息
func (c *CommonService) GetAPIInfo() (*APIInfo, error) {
	resp, err := c.client.get(context.Background(), "/api/v1/perpetual/public/api-trading-status", nil, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string  `json:"status"`
		Data   APIInfo `json:"data"`
		Ts     int64   `json:"ts"`
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

// MasterSubTransfer 母子账户划转
// subUID: 子账户UID
// symbol: 币种
// amount: 划转数量
// transferType: 划转类型 "master_to_sub"-母账户向子账户划转 "sub_to_master"-子账户向母账户划转
func (c *CommonService) MasterSubTransfer(subUID int64, symbol, amount, transferType string) error {
	if subUID <= 0 {
		return fmt.Errorf("subUID is required")
	}
	if symbol == "" {
		return fmt.Errorf("symbol is required")
	}
	if amount == "" {
		return fmt.Errorf("amount is required")
	}
	if transferType == "" {
		return fmt.Errorf("transferType is required")
	}

	body := map[string]interface{}{
		"sub_uid": subUID,
		"symbol":  symbol,
		"amount":  amount,
		"type":    transferType,
	}

	resp, err := c.client.post(context.Background(), "/api/v1/perpetual/account/master-sub-transfer", body, true)
	if err != nil {
		return err
	}

	var response struct {
		Status string `json:"status"`
		Data   string `json:"data"`
		Ts     int64  `json:"ts"`
	}

	// 直接解析到目标结构体
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return fmt.Errorf("marshal response data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &response)
	if err != nil {
		return err
	}

	if response.Status != "ok" {
		return fmt.Errorf("API error: status=%s", response.Status)
	}

	return nil
}

// GetMasterSubTransferRecord 获取母子账户划转记录
// symbol: 币种
// transferType: 划转类型，可选
// startTime: 开始时间，可选
// endTime: 结束时间，可选
// from: 查询起始ID，可选
// size: 查询条数，可选，默认20，最大50
func (c *CommonService) GetMasterSubTransferRecord(symbol, transferType string, startTime, endTime int64, from, size int) ([]MasterSubTransferRecord, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}
	if transferType != "" {
		params["type"] = transferType
	}
	if startTime > 0 {
		params["start_time"] = strconv.FormatInt(startTime, 10)
	}
	if endTime > 0 {
		params["end_time"] = strconv.FormatInt(endTime, 10)
	}
	if from > 0 {
		params["from"] = strconv.Itoa(from)
	}
	if size > 0 {
		params["size"] = strconv.Itoa(size)
	}

	resp, err := c.client.get(context.Background(), "/api/v1/perpetual/account/master-sub-transfer-record", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string                    `json:"status"`
		Data   []MasterSubTransferRecord `json:"data"`
		Ts     int64                     `json:"ts"`
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
