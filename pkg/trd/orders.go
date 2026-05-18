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

	"google.golang.org/protobuf/proto"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplaceorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdreconfirmorder"
)

// Order represents an order with its ID, code, side, type, status, price, quantity, and fill information.
type Order struct {
	OrderID         uint64
	OrderIDEx       string
	Code            string
	Name            string
	TrdSide         int32
	OrderType       int32
	OrderStatus     int32
	Price           float64
	Qty             float64
	FillQty         float64
	FillAvgPrice    float64
	CreateTime      string
	UpdateTime      string
	LastErrMsg      string
	SecMarket       int32
	CreateTimestamp float64
	UpdateTimestamp float64
	Remark          string
	TimeInForce     int32
	FillOutsideRTH  bool
	AuxPrice        float64
	TrailType       int32
	TrailValue      float64
	TrailSpread     float64
	Currency        int32
	TrdMarket       int32
	Session         int32
	JpAccType       int32
}

// PlaceOrderRequest is the request to place a new order.
type PlaceOrderRequest struct {
	AccID              uint64
	TrdMarket          constant.TrdMarket
	TrdEnv             constant.TrdEnv
	Code               string
	TrdSide            constant.TrdSide
	OrderType          constant.OrderType
	Price              float64
	Qty                float64
	AdjustPrice        bool
	AdjustSideAndLimit float64
	SecMarket          constant.TrdSecMarket
	Remark             string
	TimeInForce        int32
	FillOutsideRTH     bool
	AuxPrice           float64
	TrailType          constant.TrailType
	TrailValue         float64
	TrailSpread        float64
	Session            int32
	PositionID         uint64
}

// PlaceOrderResponse is the response containing the newly placed order ID.
type PlaceOrderResponse struct {
	Header    *trdcommon.TrdHeader
	OrderID   uint64
	OrderIDEx string
}

// PlaceOrder places a new order and returns the order ID.
// Returns the order ID or an error if the placement fails.
func PlaceOrder(ctx context.Context, c *futuapi.Client, req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("PlaceOrder: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}
	if req.Code == "" {
		return nil, fmt.Errorf("stock code is required")
	}
	if req.Qty <= 0 {
		return nil, fmt.Errorf("invalid quantity: must be positive")
	}
	if req.OrderType <= 0 {
		return nil, fmt.Errorf("invalid order type: must be valid order type constant")
	}
	if req.TrdSide <= 0 {
		return nil, fmt.Errorf("invalid trade side: must be buy/sell/other valid type")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)
	trdSide := int32(req.TrdSide)
	orderType := int32(req.OrderType)
	secMarket := int32(req.SecMarket)
	trailType := int32(req.TrailType)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdplaceorder.C2S{
		Header:    header,
		TrdSide:   &trdSide,
		OrderType: &orderType,
		Code:      &req.Code,
		Qty:       &req.Qty,
		PacketID: &common.PacketID{
			ConnID: proto.Uint64(c.GetConnID()),
		},
	}
	if req.Price != 0 {
		c2s.Price = &req.Price
	}
	if req.AdjustPrice {
		c2s.AdjustPrice = &req.AdjustPrice
	}
	if req.AdjustSideAndLimit != 0 {
		c2s.AdjustSideAndLimit = &req.AdjustSideAndLimit
	}
	if req.SecMarket != 0 {
		c2s.SecMarket = &secMarket
	}
	if req.Remark != "" {
		c2s.Remark = &req.Remark
	}
	if req.TimeInForce != 0 {
		c2s.TimeInForce = &req.TimeInForce
	}
	if req.FillOutsideRTH {
		c2s.FillOutsideRTH = &req.FillOutsideRTH
	}
	if req.AuxPrice != 0 {
		c2s.AuxPrice = &req.AuxPrice
	}
	if req.TrailType != 0 {
		c2s.TrailType = &trailType
	}
	if req.TrailValue != 0 {
		c2s.TrailValue = &req.TrailValue
	}
	if req.TrailSpread != 0 {
		c2s.TrailSpread = &req.TrailSpread
	}
	if req.Session != 0 {
		c2s.Session = &req.Session
	}
	if req.PositionID != 0 {
		c2s.PositionID = &req.PositionID
	}

	serialNo := c.NextSerialNo()
	c2s.PacketID.SerialNo = &serialNo

	pkt := &trdplaceorder.Request{C2S: c2s}
	var rsp trdplaceorder.Response

	if err := c.RequestContext(ctx, ProtoID_PlaceOrder, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("PlaceOrder", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("PlaceOrder", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &PlaceOrderResponse{
		Header:    s2c.GetHeader(),
		OrderID:   s2c.GetOrderID(),
		OrderIDEx: s2c.GetOrderIDEx(),
	}, nil
}

// ModifyOrderRequest is the request to modify an existing order (cancel, update price, or update quantity).
type ModifyOrderRequest struct {
	AccID              uint64
	TrdMarket          constant.TrdMarket
	TrdEnv             constant.TrdEnv
	OrderID            uint64
	ModifyOrderOp      constant.ModifyOrderOp
	Price              float64
	Qty                float64
	ForAll             bool
	TrdMarket2         constant.TrdMarket
	AdjustPrice        bool
	AdjustSideAndLimit float64
	AuxPrice           float64
	TrailType          constant.TrailType
	TrailValue         float64
	TrailSpread        float64
	OrderIDEx          string
}

// ModifyOrderResponse is the response returned after modifying an order.
type ModifyOrderResponse struct {
	AccID     uint64
	TrdEnv    int32
	TrdMarket int32
	OrderID   uint64
	OrderIDEx string
}

// ModifyOrder modifies or cancels an existing order.
// Returns the modification response or an error if it fails.
func ModifyOrder(ctx context.Context, c *futuapi.Client, req *ModifyOrderRequest) (*ModifyOrderResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("ModifyOrder: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}
	if req.OrderID == 0 && req.OrderIDEx == "" {
		return nil, fmt.Errorf("order ID or OrderIDEx must be provided")
	}
	if req.ModifyOrderOp <= 0 {
		return nil, fmt.Errorf("invalid modify operation: must be valid order operation type")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)
	trdMarket2 := int32(req.TrdMarket2)
	modifyOrderOp := int32(req.ModifyOrderOp)
	trailType := int32(req.TrailType)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	orderID := req.OrderID
	c2s := &trdmodifyorder.C2S{
		Header:        header,
		OrderID:       &orderID,
		ModifyOrderOp: &modifyOrderOp,
		PacketID: &common.PacketID{
			ConnID: proto.Uint64(c.GetConnID()),
		},
	}
	if req.Qty != 0 {
		c2s.Qty = &req.Qty
	}
	if req.Price != 0 {
		c2s.Price = &req.Price
	}
	if req.ForAll {
		c2s.ForAll = &req.ForAll
	}
	if req.TrdMarket2 != 0 {
		c2s.TrdMarket = &trdMarket2
	}
	if req.AdjustPrice {
		c2s.AdjustPrice = &req.AdjustPrice
	}
	if req.AdjustSideAndLimit != 0 {
		c2s.AdjustSideAndLimit = &req.AdjustSideAndLimit
	}
	if req.AuxPrice != 0 {
		c2s.AuxPrice = &req.AuxPrice
	}
	if req.TrailType != 0 {
		c2s.TrailType = &trailType
	}
	if req.TrailValue != 0 {
		c2s.TrailValue = &req.TrailValue
	}
	if req.TrailSpread != 0 {
		c2s.TrailSpread = &req.TrailSpread
	}
	if req.OrderIDEx != "" {
		c2s.OrderIDEx = &req.OrderIDEx
	}

	serialNo := c.NextSerialNo()
	c2s.PacketID.SerialNo = &serialNo

	pkt := &trdmodifyorder.Request{C2S: c2s}
	var rsp trdmodifyorder.Response

	if err := c.RequestContext(ctx, ProtoID_ModifyOrder, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("ModifyOrder", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("ModifyOrder", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &ModifyOrderResponse{
		AccID:     s2c.GetHeader().GetAccID(),
		TrdEnv:    s2c.GetHeader().GetTrdEnv(),
		TrdMarket: s2c.GetHeader().GetTrdMarket(),
		OrderID:   s2c.GetOrderID(),
		OrderIDEx: s2c.GetOrderIDEx(),
	}, nil
}

// ReconfirmOrderRequest is the request to reconfirm an order with a specified reason.
type ReconfirmOrderRequest struct {
	PacketID        *common.PacketID
	AccID           uint64
	TrdMarket       constant.TrdMarket
	TrdEnv          constant.TrdEnv
	OrderID         uint64
	ReconfirmReason int32
}

// ReconfirmOrderResponse is the response containing the reconfirmed order details.
type ReconfirmOrderResponse struct {
	AccID     uint64
	TrdEnv    int32
	TrdMarket int32
	JpAccType int32
	OrderID   uint64
}

// ReconfirmOrder reconfirms an order that requires additional verification.
// Returns the reconfirmed order details or an error if the request fails.
func ReconfirmOrder(ctx context.Context, c *futuapi.Client, req *ReconfirmOrderRequest) (*ReconfirmOrderResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("ReconfirmOrder: request is nil")
	}
	if req.OrderID == 0 {
		return nil, fmt.Errorf("invalid order ID: must be non-zero")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
		TrdEnv:    &trdEnv,
	}

	c2s := &trdreconfirmorder.C2S{
		PacketID:        req.PacketID,
		Header:          header,
		OrderID:         &req.OrderID,
		ReconfirmReason: &req.ReconfirmReason,
	}

	pkt := &trdreconfirmorder.Request{C2S: c2s}
	var rsp trdreconfirmorder.Response

	if err := c.RequestContext(ctx, ProtoID_ReconfirmOrder, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("ReconfirmOrder", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("ReconfirmOrder", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &ReconfirmOrderResponse{
		AccID:     s2c.GetHeader().GetAccID(),
		TrdEnv:    s2c.GetHeader().GetTrdEnv(),
		TrdMarket: s2c.GetHeader().GetTrdMarket(),
		JpAccType: s2c.GetHeader().GetJpAccType(),
		OrderID:   s2c.GetOrderID(),
	}, nil
}
