package hotcoin

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// PositionService 持仓接口服务

// GetPositions 获取用户持仓信息
// symbol: 交易对，可选，如果不传则返回所有
func (p *PositionService) GetPositions(symbol string) ([]PositionDetail, error) {
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}

	resp, err := p.client.get(context.Background(), "/api/v1/perpetual/positions", params, true)
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

// GetSubPositions 获取所有子账户资产信息
// 仅对母账户有效
func (p *PositionService) GetSubPositions() ([]PositionDetail, error) {
	resp, err := p.client.get(context.Background(), "/api/v1/perpetual/account/sub-accounts", nil, true)
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

// GetSubPositionInfo 获取单个子账户资产信息
// subUID: 子账户UID
// symbol: 交易对，可选
func (p *PositionService) GetSubPositionInfo(subUID int64, symbol string) ([]PositionDetail, error) {
	if subUID <= 0 {
		return nil, fmt.Errorf("subUID is required")
	}

	params := map[string]string{
		"sub_uid": strconv.FormatInt(subUID, 10),
	}
	if symbol != "" {
		params["symbol"] = symbol
	}

	resp, err := p.client.get(context.Background(), "/api/v1/perpetual/account/sub-account-info", params, true)
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

// GetSubAccountPositions 获取单个子账户持仓信息
// subUID: 子账户UID
// symbol: 交易对，可选
func (p *PositionService) GetSubAccountPositions(subUID int64, symbol string) ([]PositionDetail, error) {
	if subUID <= 0 {
		return nil, fmt.Errorf("subUID is required")
	}

	params := map[string]string{
		"sub_uid": strconv.FormatInt(subUID, 10),
	}
	if symbol != "" {
		params["symbol"] = symbol
	}

	resp, err := p.client.get(context.Background(), "/api/v1/perpetual/positions/sub-account", params, true)
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

// ClosePosition 一键平仓
// symbol: 交易对
// direction: 持仓方向 "buy"-多仓 "sell"-空仓，为空则平所有仓位
func (p *PositionService) ClosePosition(symbol, direction string) error {
	if symbol == "" {
		return fmt.Errorf("symbol is required")
	}

	body := map[string]interface{}{
		"symbol": symbol,
	}
	if direction != "" {
		body["direction"] = direction
	}

	resp, err := p.client.post(context.Background(), "/api/v1/perpetual/positions/close", body, true)
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
