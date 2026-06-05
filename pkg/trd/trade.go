// Package trd provides trading APIs for the Futu OpenD SDK.
//
// This package covers account management, order placement and modification,
// position and funds queries, order history, and trading flow analysis.
// All trading functions require an unlocked trading account.
//
// For Python SDK migration, use the constant package for Python-style constants:
//
//	import (
//	    "github.com/shing1211/futuapi4go/pkg/constant"
//	    "github.com/shing1211/futuapi4go/pkg/trd"
//	)
//
//	// Trading environment: constant.TrdEnv_Real or constant.TrdEnv_Simulate
//	// Trade side: constant.TrdSide_Buy, constant.TrdSide_Sell
//	// Order type: constant.OrderType_Normal, constant.OrderType_Market
//	// TrdMarket: constant.TrdMarket_HK, constant.TrdMarket_US, etc.
//
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

// accs, err := trd.GetAccList(cli, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
// req := &trd.PlaceOrderRequest{
package trd

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/protobuf/proto"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetacclist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetcombomaxtrdqtys"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplacecomboorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplaceorder"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetAccList              = 2001
	ProtoID_UnlockTrade             = 2005
	ProtoID_GetFunds                = 2101
	ProtoID_GetOrderFee             = 2225
	ProtoID_GetMarginRatio          = 2223
	ProtoID_GetMaxTrdQtys           = 2111
	ProtoID_GetPositionList         = 2102
	ProtoID_GetOrderList            = 2201
	ProtoID_GetOrderFillList        = 2211
	ProtoID_GetHistoryOrderList     = 2221
	ProtoID_GetHistoryOrderFillList = 2222
	ProtoID_PlaceOrder              = 2202
	ProtoID_ModifyOrder             = 2205
	ProtoID_UpdateOrder             = 2208
	ProtoID_UpdateOrderFill         = 2218
	ProtoID_SubAccPush              = 2008
	ProtoID_ReconfirmOrder          = 2209
	ProtoID_GetFlowSummary          = 2226
	ProtoID_GetComboMaxTrdQtys     = 2112
	ProtoID_PlaceComboOrder        = 2227
)

// Acc represents a trading account with its environment, ID, type, and status.
type Acc struct {
	TrdEnv            int32
	AccID             uint64
	AccType           int32
	CardNum           string
	AccStatus         int32
	TrdMarketAuthList []int32
	SecurityFirm      int32
	SimAccType        int32
	UniCardNum        string
	AccRole           int32
	JpAccType         []int32
}

// GetAccListResponse is the response containing a list of trading accounts.
type GetAccListResponse struct {
	AccList []*Acc
}

var (
	trdHeaderPool = sync.Pool{
		New: func() interface{} { return &trdcommon.TrdHeader{} },
	}
	placeOrderC2SPool = sync.Pool{
		New: func() interface{} { return &trdplaceorder.C2S{} },
	}
)

func getTrdHeader() *trdcommon.TrdHeader {
	return trdHeaderPool.Get().(*trdcommon.TrdHeader)
}

func putTrdHeader(h *trdcommon.TrdHeader) {
	*h = trdcommon.TrdHeader{}
	trdHeaderPool.Put(h)
}

func getPlaceOrderC2S() *trdplaceorder.C2S {
	return placeOrderC2SPool.Get().(*trdplaceorder.C2S)
}

func putPlaceOrderC2S(c *trdplaceorder.C2S) {
	*c = trdplaceorder.C2S{}
	placeOrderC2SPool.Put(c)
}

// wrapError standardizes error messages for proto response failures
func wrapError(funcName string, retType int32, retMsg string) error {
	code := constant.ErrorCode(retType)
	switch retType {
	case 0:
		code = constant.ErrCodeSuccess
	case -1:
		code = constant.ErrCodeInvalidParams
	case -100:
		code = constant.ErrCodeTimeout
	case -101:
		code = constant.ErrCodeNetworkError
	case -102:
		code = constant.ErrCodeProtocolErr
	case -103:
		code = constant.ErrCodeServerBusy
	case -200:
		code = constant.ErrCodeDisconnected
	case -201:
		code = constant.ErrCodeAccNotFound
	case -202:
		code = constant.ErrCodeAccDisabled
	case -203:
		code = constant.ErrCodeAccLocked
	case -204:
		code = constant.ErrCodeAccAuthFail
	case -301:
		code = constant.ErrCodeInsufficientBalance
	case -302:
		code = constant.ErrCodeMarketClosed
	case -303:
		code = constant.ErrCodeOrderRejected
	case -304:
		code = constant.ErrCodePriceOutOfRange
	case -305:
		code = constant.ErrCodeQtyTooLarge
	case -306:
		code = constant.ErrCodeTradingDisabled
	case -307:
		code = constant.ErrCodeInvalidSecurity
	case -308:
		code = constant.ErrCodeNoPermission
	case -400:
		code = constant.ErrCodeUnknown
	case -401:
		code = constant.ErrCodeAlreadySubbed
	case -402:
		code = constant.ErrCodeNotSubbed
	default:
		if retType > 0 {
			code = constant.ErrCodeUnknown
		}
	}
	return constant.NewFutuError(code, funcName, retMsg)
}

// GetAccList retrieves the list of trading accounts, optionally including general security account info.
// Returns the account list or an error if the request fails.
func GetAccList(ctx context.Context, c *futuapi.Client, trdCategory constant.TrdCategory, needGeneralSecAccount bool) (*GetAccListResponse, error) {
	trdCategoryInt := int32(trdCategory)
	c2s := &trdgetacclist.C2S{
		UserID:                proto.Uint64(0), // Deprecated but required by protocol, set to 0
		TrdCategory:           &trdCategoryInt,
		NeedGeneralSecAccount: &needGeneralSecAccount,
	}

	pkt := &trdgetacclist.Request{C2S: c2s}
	var rsp trdgetacclist.Response

	if err := c.RequestContext(ctx, ProtoID_GetAccList, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetAccList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetAccList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetAccListResponse{
		AccList: make([]*Acc, 0, len(s2c.AccList)),
	}

	for _, acc := range s2c.AccList {
		if acc == nil {
			continue
		}
		result.AccList = append(result.AccList, &Acc{
			TrdEnv:            util.ProtoInt32(acc.TrdEnv),
			AccID:             util.ProtoUint64(acc.AccID),
			AccType:           util.ProtoInt32(acc.AccType),
			CardNum:           util.ProtoStr(acc.CardNum),
			AccStatus:         util.ProtoInt32(acc.AccStatus),
			TrdMarketAuthList: acc.TrdMarketAuthList,
			SecurityFirm:      util.ProtoInt32(acc.SecurityFirm),
			SimAccType:        util.ProtoInt32(acc.SimAccType),
			UniCardNum:        util.ProtoStr(acc.UniCardNum),
			AccRole:           util.ProtoInt32(acc.AccRole),
			JpAccType:         acc.JpAccType,
		})
	}

	return result, nil
}

// GetComboMaxTrdQtysRequest defines parameters for GetComboMaxTrdQtys.
type GetComboMaxTrdQtysRequest struct {
	Header    *trdcommon.TrdHeader
	ComboLegs []*qotcommon.ComboLeg
	Qty       float64
	Price     float64
	OrderType int32
	OrderIDEx string
}

// GetComboMaxTrdQtysResponse is the response type for GetComboMaxTrdQtys.
type GetComboMaxTrdQtysResponse struct {
	Header     *trdcommon.TrdHeader
	MaxTrdQtys *trdcommon.ComboMaxTrdQtys
}

// GetComboMaxTrdQtys returns maximum tradable quantities for combo orders.
func GetComboMaxTrdQtys(ctx context.Context, c *futuapi.Client, req *GetComboMaxTrdQtysRequest) (*GetComboMaxTrdQtysResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetComboMaxTrdQtys: request is nil")
	}
	if req.Header == nil {
		return nil, fmt.Errorf("GetComboMaxTrdQtys: Header is nil")
	}
	if len(req.ComboLegs) == 0 {
		return nil, fmt.Errorf("GetComboMaxTrdQtys: ComboLegs is empty")
	}
	if req.Qty <= 0 {
		return nil, fmt.Errorf("GetComboMaxTrdQtys: Qty must be positive")
	}
	c2s := &trdgetcombomaxtrdqtys.C2S{
		Header:    req.Header,
		ComboLegs: req.ComboLegs,
		Qty:       &req.Qty,
		OrderType: &req.OrderType,
	}
	if req.Price != 0 {
		c2s.Price = &req.Price
	}
	if req.OrderIDEx != "" {
		c2s.OrderIDEx = &req.OrderIDEx
	}
	pkt := &trdgetcombomaxtrdqtys.Request{C2S: c2s}
	var rsp trdgetcombomaxtrdqtys.Response

	if err := c.RequestContext(ctx, ProtoID_GetComboMaxTrdQtys, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetComboMaxTrdQtys", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetComboMaxTrdQtys", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetComboMaxTrdQtysResponse{
		Header:     s2c.Header,
		MaxTrdQtys: s2c.MaxTrdQtys,
	}, nil
}

// PlaceComboOrderRequest defines parameters for PlaceComboOrder.
type PlaceComboOrderRequest struct {
	PacketID    *common.PacketID
	Header      *trdcommon.TrdHeader
	ComboLegs   []*qotcommon.ComboLeg
	Qty         float64
	Price       float64
	OrderType   int32
	TimeInForce int32
	ExpireTime  string
	Remark      string
}

// PlaceComboOrderResponse is the response type for PlaceComboOrder.
type PlaceComboOrderResponse struct {
	Header    *trdcommon.TrdHeader
	OrderIDEx string
}

// PlaceComboOrder places a combo order for option strategies.
func PlaceComboOrder(ctx context.Context, c *futuapi.Client, req *PlaceComboOrderRequest) (*PlaceComboOrderResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("PlaceComboOrder: request is nil")
	}
	if req.Header == nil {
		return nil, fmt.Errorf("PlaceComboOrder: Header is nil")
	}
	if len(req.ComboLegs) == 0 {
		return nil, fmt.Errorf("PlaceComboOrder: ComboLegs is empty")
	}
	if req.Qty <= 0 {
		return nil, fmt.Errorf("PlaceComboOrder: Qty must be positive")
	}
	c2s := &trdplacecomboorder.C2S{
		PacketID:  req.PacketID,
		Header:    req.Header,
		ComboLegs: req.ComboLegs,
		Qty:       &req.Qty,
		OrderType: &req.OrderType,
	}
	if req.Price != 0 {
		c2s.Price = &req.Price
	}
	if req.TimeInForce != 0 {
		c2s.TimeInForce = &req.TimeInForce
	}
	if req.ExpireTime != "" {
		c2s.ExpireTime = &req.ExpireTime
	}
	if req.Remark != "" {
		c2s.Remark = &req.Remark
	}
	pkt := &trdplacecomboorder.Request{C2S: c2s}
	var rsp trdplacecomboorder.Response

	if err := c.RequestContext(ctx, ProtoID_PlaceComboOrder, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("PlaceComboOrder", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("PlaceComboOrder", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &PlaceComboOrderResponse{
		Header:    s2c.Header,
		OrderIDEx: util.ProtoStr(s2c.OrderIDEx),
	}, nil
}
