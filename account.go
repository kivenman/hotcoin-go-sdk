package hotcoin

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// AccountService 账户/资产接口服务

// GetAccountInfo 获取账户信息
// symbol: 保证金币种，如 USDT
func (a *AccountService) GetAccountInfo(symbol string) (*AccountInfo, error) {
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}

	resp, err := a.client.get(context.Background(), "/api/v1/perpetual/account/info", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string      `json:"status"`
		Data   AccountInfo `json:"data"`
		Ts     int64       `json:"ts"`
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

// GetAccountBalance 获取账户余额
// symbol: 保证金币种，可选，如果不传则返回所有
func (a *AccountService) GetAccountBalance(symbol string) ([]AccountBalance, error) {
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}

	resp, err := a.client.get(context.Background(), "/api/v1/perpetual/account/balance", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string           `json:"status"`
		Data   []AccountBalance `json:"data"`
		Ts     int64            `json:"ts"`
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

// GetPositions 获取用户持仓信息
// symbol: 交易对，可选，如果不传则返回所有
func (a *AccountService) GetPositions(symbol string) ([]PositionDetail, error) {
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}

	resp, err := a.client.get(context.Background(), "/api/v1/perpetual/account/positions", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string           `json:"status"`
		Data   []PositionDetail `json:"data"`
		Ts     int64            `json:"ts"`
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

// SetLeverage 设置杠杆倍数
// symbol: 交易对
// leverRate: 杠杆倍数
func (a *AccountService) SetLeverage(symbol string, leverRate int) error {
	if symbol == "" {
		return fmt.Errorf("symbol is required")
	}
	if leverRate <= 0 {
		return fmt.Errorf("leverRate must be positive")
	}

	body := map[string]interface{}{
		"symbol":     symbol,
		"lever_rate": leverRate,
	}

	resp, err := a.client.post(context.Background(), "/api/v1/perpetual/account/leverage", body, true)
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

// GetLeverageInfo 获取可用杠杆倍数
// symbol: 交易对
func (a *AccountService) GetLeverageInfo(symbol string) (*LeverageInfo, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := a.client.get(context.Background(), "/api/v1/perpetual/account/leverage-info", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string       `json:"status"`
		Data   LeverageInfo `json:"data"`
		Ts     int64        `json:"ts"`
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

// SetMarginMode 设置保证金模式
// symbol: 交易对
// marginMode: 保证金模式 isolated-逐仓 crossed-全仓
func (a *AccountService) SetMarginMode(symbol, marginMode string) error {
	if symbol == "" {
		return fmt.Errorf("symbol is required")
	}
	if marginMode == "" {
		return fmt.Errorf("marginMode is required")
	}

	body := map[string]interface{}{
		"symbol":      symbol,
		"margin_mode": marginMode,
	}

	resp, err := a.client.post(context.Background(), "/api/v1/perpetual/account/position-mode", body, true)
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

// GetFeeRate 获取合约费率
// symbol: 交易对
func (a *AccountService) GetFeeRate(symbol string) (*FeeRate, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := a.client.get(context.Background(), "/api/v1/perpetual/account/fee-rate", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string  `json:"status"`
		Data   FeeRate `json:"data"`
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

// GetTransferLimit 获取划转限额
// currency: 币种，可选，如果不传则返回所有
func (a *AccountService) GetTransferLimit(currency string) ([]TransferLimit, error) {
	params := make(map[string]string)
	if currency != "" {
		params["currency"] = currency
	}

	resp, err := a.client.get(context.Background(), "/api/v1/perpetual/account/transfer-limit", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string          `json:"status"`
		Data   []TransferLimit `json:"data"`
		Ts     int64           `json:"ts"`
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

// GetPositionLimit 获取用户持仓量限制
// symbol: 交易对
func (a *AccountService) GetPositionLimit(symbol string) (*PositionLimitInfo, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := a.client.get(context.Background(), "/api/v1/perpetual/account/position-limit", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string            `json:"status"`
		Data   PositionLimitInfo `json:"data"`
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

	return &response.Data, nil
}

// GetFinancialRecord 获取财务记录
// symbol: 保证金币种
// recordType: 记录类型，可选
// startTime: 开始时间戳，可选
// endTime: 结束时间戳，可选
// size: 条数，可选，默认20，最大50
func (a *AccountService) GetFinancialRecord(symbol string, recordType int, startTime, endTime int64, size int) ([]FinancialRecord, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}
	if recordType > 0 {
		params["type"] = strconv.Itoa(recordType)
	}
	if startTime > 0 {
		params["start_time"] = strconv.FormatInt(startTime, 10)
	}
	if endTime > 0 {
		params["end_time"] = strconv.FormatInt(endTime, 10)
	}
	if size > 0 {
		params["size"] = strconv.Itoa(size)
	}

	resp, err := a.client.get(context.Background(), "/api/v1/perpetual/account/financial-record", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string            `json:"status"`
		Data   []FinancialRecord `json:"data"`
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

// GetAssetValuation 获取总资产估值
// valuationAsset: 估值币种，可选，如USDT、BTC等
func (a *AccountService) GetAssetValuation(valuationAsset string) (*AssetValuation, error) {
	params := make(map[string]string)
	if valuationAsset != "" {
		params["valuation_asset"] = valuationAsset
	}

	resp, err := a.client.get(context.Background(), "/api/v1/perpetual/account/balance-valuation", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string         `json:"status"`
		Data   AssetValuation `json:"data"`
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

	return &response.Data, nil
}
