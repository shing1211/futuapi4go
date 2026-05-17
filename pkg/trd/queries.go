// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trd

import (
	"context"
	"fmt"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfee"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderlist"
)

// GetOrderListRequest is the request to retrieve order list.
type GetOrderListRequest struct {
	AccID            uint64
	TrdMarket        constant.TrdMarket
	TrdEnv           constant.TrdEnv
	FilterConditions *trdcommon.TrdFilterConditions
	FilterStatusList []int32
	RefreshCache     bool
}

// GetOrderListResponse is the response containing a list of orders.
type GetOrderListResponse struct {
	OrderList []*Order
}

// GetOrderList retrieves the current order list for the account.
// Returns the order list or an error if the request fails.
func GetOrderList(ctx context.Context, c *futuapi.Client, req *GetOrderListRequest) (*GetOrderListResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOrderList: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetorderlist.C2S{
		Header:           header,
		FilterConditions: req.FilterConditions,
		FilterStatusList: req.FilterStatusList,
	}
	if req.RefreshCache {
		c2s.RefreshCache = &req.RefreshCache
	}

	pkt := &trdgetorderlist.Request{C2S: c2s}
	var rsp trdgetorderlist.Response

	if err := c.RequestContext(ctx, ProtoID_GetOrderList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOrderList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOrderList: s2c is nil")
	}

	result := &GetOrderListResponse{
		OrderList: make([]*Order, 0, len(s2c.GetOrderList())),
	}

	for _, o := range s2c.GetOrderList() {
		if o == nil {
			continue
		}
		result.OrderList = append(result.OrderList, &Order{
			OrderID:         o.GetOrderID(),
			OrderIDEx:       o.GetOrderIDEx(),
			Code:            o.GetCode(),
			Name:            o.GetName(),
			TrdSide:         o.GetTrdSide(),
			OrderType:       o.GetOrderType(),
			OrderStatus:     o.GetOrderStatus(),
			Price:           o.GetPrice(),
			Qty:             o.GetQty(),
			FillQty:         o.GetFillQty(),
			FillAvgPrice:    o.GetFillAvgPrice(),
			CreateTime:      o.GetCreateTime(),
			UpdateTime:      o.GetUpdateTime(),
			LastErrMsg:      o.GetLastErrMsg(),
			SecMarket:       o.GetSecMarket(),
			CreateTimestamp: o.GetCreateTimestamp(),
			UpdateTimestamp: o.GetUpdateTimestamp(),
			Remark:          o.GetRemark(),
			TimeInForce:     o.GetTimeInForce(),
			FillOutsideRTH:  o.GetFillOutsideRTH(),
			AuxPrice:        o.GetAuxPrice(),
			TrailType:       o.GetTrailType(),
			TrailValue:      o.GetTrailValue(),
			TrailSpread:     o.GetTrailSpread(),
			Currency:        o.GetCurrency(),
			TrdMarket:       o.GetTrdMarket(),
			Session:         o.GetSession(),
			JpAccType:       o.GetJpAccType(),
		})
	}

	return result, nil
}

// OrderFill represents a filled (executed) order with its order ID, fill ID, code, side, price, and quantity.
type OrderFill struct {
	FillID            uint64
	FillIDEx          string
	OrderID           uint64
	OrderIDEx         string
	Code              string
	Name              string
	TrdSide           int32
	Price             float64
	Qty               float64
	CreateTime        string
	CounterBrokerID   int32
	CounterBrokerName string
	SecMarket         int32
	CreateTimestamp   float64
	UpdateTimestamp   float64
	Status            int32
	TrdMarket         int32
	JpAccType         int32
}

// GetOrderFillListRequest is the request to retrieve order fill list.
type GetOrderFillListRequest struct {
	AccID            uint64
	TrdMarket        constant.TrdMarket
	TrdEnv           constant.TrdEnv
	FilterConditions *trdcommon.TrdFilterConditions
}

// GetOrderFillListResponse is the response containing a list of order fills.
type GetOrderFillListResponse struct {
	OrderFillList []*OrderFill
}

// GetOrderFillList retrieves the current order fill (execution) list for the account.
// Returns the order fill list or an error if the request fails.
func GetOrderFillList(ctx context.Context, c *futuapi.Client, req *GetOrderFillListRequest) (*GetOrderFillListResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOrderFillList: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetorderfilllist.C2S{
		Header:           header,
		FilterConditions: req.FilterConditions,
	}

	pkt := &trdgetorderfilllist.Request{C2S: c2s}
	var rsp trdgetorderfilllist.Response

	if err := c.RequestContext(ctx, ProtoID_GetOrderFillList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOrderFillList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOrderFillList: s2c is nil")
	}

	result := &GetOrderFillListResponse{
		OrderFillList: make([]*OrderFill, 0, len(s2c.GetOrderFillList())),
	}

	for _, f := range s2c.GetOrderFillList() {
		if f == nil {
			continue
		}
		result.OrderFillList = append(result.OrderFillList, &OrderFill{
			FillID:            f.GetFillID(),
			FillIDEx:          f.GetFillIDEx(),
			OrderID:           f.GetOrderID(),
			OrderIDEx:         f.GetOrderIDEx(),
			Code:              f.GetCode(),
			Name:              f.GetName(),
			TrdSide:           f.GetTrdSide(),
			Price:             f.GetPrice(),
			Qty:               f.GetQty(),
			CreateTime:        f.GetCreateTime(),
			CounterBrokerID:   f.GetCounterBrokerID(),
			CounterBrokerName: f.GetCounterBrokerName(),
			SecMarket:         f.GetSecMarket(),
			CreateTimestamp:   f.GetCreateTimestamp(),
			UpdateTimestamp:   f.GetUpdateTimestamp(),
			Status:            f.GetStatus(),
			TrdMarket:         f.GetTrdMarket(),
			JpAccType:         f.GetJpAccType(),
		})
	}

	return result, nil
}

// GetOrderFeeRequest is the request to retrieve order fee information.
type GetOrderFeeRequest struct {
	AccID         uint64
	TrdMarket     constant.TrdMarket
	TrdEnv        constant.TrdEnv
	OrderIDExList []string
}

// OrderFeeInfo represents the fee information for a single order.
type OrderFeeInfo struct {
	OrderIDEx string
	FeeAmount float64
	FeeList   []*OrderFeeItemInfo
}

// OrderFeeItemInfo represents a single fee item with its title and value.
type OrderFeeItemInfo struct {
	Title string
	Value float64
}

// GetOrderFeeResponse is the response containing order fee information.
type GetOrderFeeResponse struct {
	OrderFeeList []*OrderFeeInfo
}

// GetOrderFee retrieves the fee details for specified orders.
// Returns the order fee list or an error if the request fails.
func GetOrderFee(ctx context.Context, c *futuapi.Client, req *GetOrderFeeRequest) (*GetOrderFeeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOrderFee: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}
	if len(req.OrderIDExList) == 0 {
		return nil, fmt.Errorf("order ID list is empty")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetorderfee.C2S{
		Header:        header,
		OrderIdExList: req.OrderIDExList,
	}

	pkt := &trdgetorderfee.Request{C2S: c2s}
	var rsp trdgetorderfee.Response

	if err := c.RequestContext(ctx, ProtoID_GetOrderFee, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOrderFee", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOrderFee: s2c is nil")
	}

	result := &GetOrderFeeResponse{
		OrderFeeList: make([]*OrderFeeInfo, 0, len(s2c.GetOrderFeeList())),
	}

	for _, f := range s2c.GetOrderFeeList() {
		if f == nil {
			continue
		}
		feeInfo := &OrderFeeInfo{
			OrderIDEx: f.GetOrderIDEx(),
			FeeAmount: f.GetFeeAmount(),
			FeeList:   make([]*OrderFeeItemInfo, 0, len(f.GetFeeList())),
		}
		for _, item := range f.GetFeeList() {
			if item == nil {
				continue
			}
			feeInfo.FeeList = append(feeInfo.FeeList, &OrderFeeItemInfo{
				Title: item.GetTitle(),
				Value: item.GetValue(),
			})
		}
		result.OrderFeeList = append(result.OrderFeeList, feeInfo)
	}

	return result, nil
}

// GetHistoryOrderListRequest is the request to retrieve historical order list.
type GetHistoryOrderListRequest struct {
	AccID            uint64
	TrdMarket        constant.TrdMarket
	TrdEnv           constant.TrdEnv
	FilterConditions *trdcommon.TrdFilterConditions
	FilterStatusList []int32
}

// GetHistoryOrderListResponse is the response containing historical orders.
type GetHistoryOrderListResponse struct {
	OrderList []*Order
}

// GetHistoryOrderList retrieves the historical order list based on filter conditions.
// Returns the historical order list or an error if the request fails.
func GetHistoryOrderList(ctx context.Context, c *futuapi.Client, req *GetHistoryOrderListRequest) (*GetHistoryOrderListResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHistoryOrderList: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgethistoryorderlist.C2S{
		Header:           header,
		FilterConditions: req.FilterConditions,
		FilterStatusList: req.FilterStatusList,
	}
	if c2s.FilterConditions == nil {
		c2s.FilterConditions = &trdcommon.TrdFilterConditions{}
	}

	pkt := &trdgethistoryorderlist.Request{C2S: c2s}
	var rsp trdgethistoryorderlist.Response

	if err := c.RequestContext(ctx, ProtoID_GetHistoryOrderList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHistoryOrderList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetHistoryOrderList: s2c is nil")
	}

	orderList := make([]*Order, 0, len(s2c.GetOrderList()))
	for _, o := range s2c.GetOrderList() {
		if o == nil {
			continue
		}
		orderList = append(orderList, &Order{
			OrderID:         o.GetOrderID(),
			OrderIDEx:       o.GetOrderIDEx(),
			Code:            o.GetCode(),
			Name:            o.GetName(),
			TrdSide:         o.GetTrdSide(),
			OrderType:       o.GetOrderType(),
			OrderStatus:     o.GetOrderStatus(),
			Price:           o.GetPrice(),
			Qty:             o.GetQty(),
			FillQty:         o.GetFillQty(),
			FillAvgPrice:    o.GetFillAvgPrice(),
			CreateTime:      o.GetCreateTime(),
			UpdateTime:      o.GetUpdateTime(),
			LastErrMsg:      o.GetLastErrMsg(),
			SecMarket:       o.GetSecMarket(),
			CreateTimestamp: o.GetCreateTimestamp(),
			UpdateTimestamp: o.GetUpdateTimestamp(),
			Remark:          o.GetRemark(),
			TimeInForce:     o.GetTimeInForce(),
			FillOutsideRTH:  o.GetFillOutsideRTH(),
			AuxPrice:        o.GetAuxPrice(),
			TrailType:       o.GetTrailType(),
			TrailValue:      o.GetTrailValue(),
			TrailSpread:     o.GetTrailSpread(),
			Currency:        o.GetCurrency(),
			TrdMarket:       o.GetTrdMarket(),
			Session:         o.GetSession(),
			JpAccType:       o.GetJpAccType(),
		})
	}
	return &GetHistoryOrderListResponse{
		OrderList: orderList,
	}, nil
}

// GetHistoryOrderFillListRequest is the request to retrieve historical order fill list.
type GetHistoryOrderFillListRequest struct {
	AccID            uint64
	TrdMarket        constant.TrdMarket
	TrdEnv           constant.TrdEnv
	FilterConditions *trdcommon.TrdFilterConditions
}

// GetHistoryOrderFillListResponse is the response containing historical order fills.
type GetHistoryOrderFillListResponse struct {
	OrderFillList []*OrderFill
}

// GetHistoryOrderFillList retrieves the historical order fill (execution) list based on filter conditions.
// Returns the historical order fill list or an error if the request fails.
func GetHistoryOrderFillList(ctx context.Context, c *futuapi.Client, req *GetHistoryOrderFillListRequest) (*GetHistoryOrderFillListResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHistoryOrderFillList: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgethistoryorderfilllist.C2S{
		Header:           header,
		FilterConditions: req.FilterConditions,
	}
	if c2s.FilterConditions == nil {
		c2s.FilterConditions = &trdcommon.TrdFilterConditions{}
	}
	if c2s.FilterConditions.GetBeginTime() == "" {
		begin := time.Now().AddDate(0, 0, -30).Format("2006-01-02 15:04:05")
		c2s.FilterConditions.BeginTime = &begin
	}
	if c2s.FilterConditions.GetEndTime() == "" {
		end := time.Now().Format("2006-01-02 15:04:05")
		c2s.FilterConditions.EndTime = &end
	}

	pkt := &trdgethistoryorderfilllist.Request{C2S: c2s}
	var rsp trdgethistoryorderfilllist.Response

	if err := c.RequestContext(ctx, ProtoID_GetHistoryOrderFillList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHistoryOrderFillList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetHistoryOrderFillList: s2c is nil")
	}

	list := make([]*OrderFill, 0, len(s2c.GetOrderFillList()))
	for _, f := range s2c.GetOrderFillList() {
		if f == nil {
			continue
		}
		list = append(list, &OrderFill{
			FillID:            f.GetFillID(),
			FillIDEx:          f.GetFillIDEx(),
			OrderID:           f.GetOrderID(),
			OrderIDEx:         f.GetOrderIDEx(),
			Code:              f.GetCode(),
			Name:              f.GetName(),
			TrdSide:           f.GetTrdSide(),
			Price:             f.GetPrice(),
			Qty:               f.GetQty(),
			CreateTime:        f.GetCreateTime(),
			CounterBrokerID:   f.GetCounterBrokerID(),
			CounterBrokerName: f.GetCounterBrokerName(),
			SecMarket:         f.GetSecMarket(),
			CreateTimestamp:   f.GetCreateTimestamp(),
			UpdateTimestamp:   f.GetUpdateTimestamp(),
			Status:            f.GetStatus(),
			TrdMarket:         f.GetTrdMarket(),
			JpAccType:         f.GetJpAccType(),
		})
	}

	return &GetHistoryOrderFillListResponse{
		OrderFillList: list,
	}, nil
}
