package hotcoin

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// TradingService 交易接口服务

// PlaceOrder 下单
func (t *TradingService) PlaceOrder(req *OrderPlaceRequest) (*OrderPlaceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if req.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if req.Direction == "" {
		return nil, fmt.Errorf("direction is required")
	}
	if req.Volume == "" {
		return nil, fmt.Errorf("volume is required")
	}

	resp, err := t.client.post(context.Background(), "/api/v1/perpetual/orders", req, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string             `json:"status"`
		Data   OrderPlaceResponse `json:"data"`
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

	return &response.Data, nil
}

// PlaceBatchOrders 批量下单
func (t *TradingService) PlaceBatchOrders(req *BatchOrderRequest) (*BatchOrderResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if len(req.Orders) == 0 {
		return nil, fmt.Errorf("orders is required")
	}
	if len(req.Orders) > 10 {
		return nil, fmt.Errorf("orders count cannot exceed 10")
	}

	resp, err := t.client.post(context.Background(), "/api/v1/perpetual/orders/batch", req, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string             `json:"status"`
		Data   BatchOrderResponse `json:"data"`
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

	return &response.Data, nil
}

// CancelOrder 撤销订单
func (t *TradingService) CancelOrder(req *OrderCancelRequest) (*OrderCancelResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if req.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if req.OrderID == "" && req.ClientOrderID == "" {
		return nil, fmt.Errorf("order_id or client_order_id is required")
	}

	resp, err := t.client.post(context.Background(), "/api/v1/perpetual/orders/cancel", req, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string              `json:"status"`
		Data   OrderCancelResponse `json:"data"`
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

	return &response.Data, nil
}

// CancelAllOrders 撤销所有订单
func (t *TradingService) CancelAllOrders(symbol, contractCode, contractType string) (*OrderCancelResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	body := map[string]interface{}{
		"symbol": symbol,
	}
	if contractCode != "" {
		body["contract_code"] = contractCode
	}
	if contractType != "" {
		body["contract_type"] = contractType
	}

	resp, err := t.client.post(context.Background(), "/api/v1/perpetual/orders/cancel-all", body, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string              `json:"status"`
		Data   OrderCancelResponse `json:"data"`
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

	return &response.Data, nil
}

// GetOrderInfo 获取订单信息
func (t *TradingService) GetOrderInfo(symbol, orderID, clientOrderID string) ([]Order, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if orderID == "" && clientOrderID == "" {
		return nil, fmt.Errorf("order_id or client_order_id is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}
	if orderID != "" {
		params["order_id"] = orderID
	}
	if clientOrderID != "" {
		params["client_order_id"] = clientOrderID
	}

	resp, err := t.client.get(context.Background(), "/api/v1/perpetual/orders/info", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string  `json:"status"`
		Data   []Order `json:"data"`
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

	return response.Data, nil
}

// GetOrderDetail 获取订单明细信息
func (t *TradingService) GetOrderDetail(symbol, orderID string) (*OrderDetail, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if orderID == "" {
		return nil, fmt.Errorf("order_id is required")
	}

	params := map[string]string{
		"symbol":   symbol,
		"order_id": orderID,
	}

	resp, err := t.client.get(context.Background(), "/api/v1/perpetual/orders/detail", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string      `json:"status"`
		Data   OrderDetail `json:"data"`
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

// GetOpenOrders 获取当前委托
func (t *TradingService) GetOpenOrders(symbol string, pageIndex, pageSize int) ([]Order, error) {
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

	resp, err := t.client.get(context.Background(), "/api/v1/perpetual/orders/open", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string  `json:"status"`
		Data   []Order `json:"data"`
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

	return response.Data, nil
}

// GetOrderHistory 获取历史委托
func (t *TradingService) GetOrderHistory(req *OrderQueryRequest) ([]Order, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if req.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": req.Symbol,
	}
	if req.OrderType > 0 {
		params["order_type"] = strconv.Itoa(req.OrderType)
	}
	if req.Status > 0 {
		params["status"] = strconv.Itoa(req.Status)
	}
	if req.CreateDate > 0 {
		params["create_date"] = strconv.Itoa(req.CreateDate)
	}
	if req.PageIndex > 0 {
		params["page_index"] = strconv.Itoa(req.PageIndex)
	}
	if req.PageSize > 0 {
		params["page_size"] = strconv.Itoa(req.PageSize)
	}

	resp, err := t.client.get(context.Background(), "/api/v1/perpetual/orders/history", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string  `json:"status"`
		Data   []Order `json:"data"`
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

	return response.Data, nil
}

// GetMatchResults 获取用户的成交记录
func (t *TradingService) GetMatchResults(symbol string, tradeType int, createDate int, pageIndex, pageSize int) ([]MatchResult, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}
	if tradeType > 0 {
		params["trade_type"] = strconv.Itoa(tradeType)
	}
	if createDate > 0 {
		params["create_date"] = strconv.Itoa(createDate)
	}
	if pageIndex > 0 {
		params["page_index"] = strconv.Itoa(pageIndex)
	}
	if pageSize > 0 {
		params["page_size"] = strconv.Itoa(pageSize)
	}

	resp, err := t.client.get(context.Background(), "/api/v1/perpetual/orders/match-results", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string        `json:"status"`
		Data   []MatchResult `json:"data"`
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

	return response.Data, nil
}

// PlacePlanOrder 计划委托下单
func (t *TradingService) PlacePlanOrder(req *PlanOrderRequest) (*OrderPlaceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if req.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if req.TriggerPrice == "" {
		return nil, fmt.Errorf("trigger_price is required")
	}
	if req.Volume == "" {
		return nil, fmt.Errorf("volume is required")
	}

	resp, err := t.client.post(context.Background(), "/api/v1/perpetual/orders/trigger", req, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string             `json:"status"`
		Data   OrderPlaceResponse `json:"data"`
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

	return &response.Data, nil
}

// CancelPlanOrder 撤销计划委托订单
func (t *TradingService) CancelPlanOrder(symbol, orderID string) (*OrderCancelResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if orderID == "" {
		return nil, fmt.Errorf("order_id is required")
	}

	body := map[string]interface{}{
		"symbol":   symbol,
		"order_id": orderID,
	}

	resp, err := t.client.post(context.Background(), "/api/v1/perpetual/orders/trigger/cancel", body, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string              `json:"status"`
		Data   OrderCancelResponse `json:"data"`
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

	return &response.Data, nil
}

// CancelAllPlanOrders 撤销所有计划委托订单
func (t *TradingService) CancelAllPlanOrders(symbol, contractCode, contractType string) (*OrderCancelResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	body := map[string]interface{}{
		"symbol": symbol,
	}
	if contractCode != "" {
		body["contract_code"] = contractCode
	}
	if contractType != "" {
		body["contract_type"] = contractType
	}

	resp, err := t.client.post(context.Background(), "/api/v1/perpetual/orders/trigger/cancel-all", body, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string              `json:"status"`
		Data   OrderCancelResponse `json:"data"`
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

	return &response.Data, nil
}

// GetPlanOrders 获取计划委托当前委托
func (t *TradingService) GetPlanOrders(symbol string, pageIndex, pageSize int) ([]PlanOrder, error) {
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

	resp, err := t.client.get(context.Background(), "/api/v1/perpetual/orders/trigger/open", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string      `json:"status"`
		Data   []PlanOrder `json:"data"`
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

	return response.Data, nil
}

// GetPlanOrderHistory 获取计划委托历史委托
func (t *TradingService) GetPlanOrderHistory(symbol string, status int, createDate int, pageIndex, pageSize int) ([]PlanOrder, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}
	if status > 0 {
		params["status"] = strconv.Itoa(status)
	}
	if createDate > 0 {
		params["create_date"] = strconv.Itoa(createDate)
	}
	if pageIndex > 0 {
		params["page_index"] = strconv.Itoa(pageIndex)
	}
	if pageSize > 0 {
		params["page_size"] = strconv.Itoa(pageSize)
	}

	resp, err := t.client.get(context.Background(), "/api/v1/perpetual/orders/trigger/history", params, true)
	if err != nil {
		return nil, err
	}

	var response struct {
		Status string      `json:"status"`
		Data   []PlanOrder `json:"data"`
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

	return response.Data, nil
}
