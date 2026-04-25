package trd

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"google.golang.org/protobuf/proto"
)

type CancelAllOrdersResult struct {
	AccID     uint64
	Cancelled int32
}

func CancelAllOrders(ctx context.Context, c *futuapi.Client, accID uint64, market constant.TrdMarket, env constant.TrdEnv) (*CancelAllOrdersResult, error) {
	trdMarket := int32(market)
	trdEnv := int32(env)
	header := &trdcommon.TrdHeader{
		AccID:     &accID,
		TrdMarket: &trdMarket,
		TrdEnv:    &trdEnv,
	}
	c2s := &trdmodifyorder.C2S{
		Header:        header,
		ModifyOrderOp: proto.Int32(int32(constant.ModifyOrderOp_Cancel)),
		ForAll:        proto.Bool(true),
	}
	pkt := &trdmodifyorder.Request{C2S: c2s}
	var rsp trdmodifyorder.Response

	if err := c.RequestContext(ctx, ProtoID_ModifyOrder, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("CancelAllOrders", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return &CancelAllOrdersResult{
		AccID: accID,
	}, nil
}

func QuickBuy(ctx context.Context, c *futuapi.Client, accID uint64, market constant.TrdMarket, env constant.TrdEnv, code string, qty float64, price float64) (*PlaceOrderResponse, error) {
	return PlaceOrder(ctx, c, &PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    env,
		Code:      code,
		TrdSide:   constant.TrdSide_Buy,
		OrderType: constant.OrderType_Normal,
		Price:     price,
		Qty:       qty,
	})
}

func QuickSell(ctx context.Context, c *futuapi.Client, accID uint64, market constant.TrdMarket, env constant.TrdEnv, code string, qty float64, price float64) (*PlaceOrderResponse, error) {
	return PlaceOrder(ctx, c, &PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    env,
		Code:      code,
		TrdSide:   constant.TrdSide_Sell,
		OrderType: constant.OrderType_Normal,
		Price:     price,
		Qty:       qty,
	})
}

func QuickMarketBuy(ctx context.Context, c *futuapi.Client, accID uint64, market constant.TrdMarket, env constant.TrdEnv, code string, qty float64) (*PlaceOrderResponse, error) {
	return PlaceOrder(ctx, c, &PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    env,
		Code:      code,
		TrdSide:   constant.TrdSide_Buy,
		OrderType: constant.OrderType_Market,
		Price:     0,
		Qty:       qty,
	})
}

func QuickMarketSell(ctx context.Context, c *futuapi.Client, accID uint64, market constant.TrdMarket, env constant.TrdEnv, code string, qty float64) (*PlaceOrderResponse, error) {
	return PlaceOrder(ctx, c, &PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    env,
		Code:      code,
		TrdSide:   constant.TrdSide_Sell,
		OrderType: constant.OrderType_Market,
		Price:     0,
		Qty:       qty,
	})
}

type PositionDetail struct {
	Code        string
	Name        string
	Qty         float64
	CostPrice   float64
	CostBalance float64
	Market      constant.TrdMarket
}

func GetPositions(ctx context.Context, c *futuapi.Client, accID uint64) ([]PositionDetail, error) {
	req := &GetPositionListRequest{
		AccID:     accID,
		TrdMarket: constant.TrdMarket_None,
	}
	rsp, err := GetPositionList(ctx, c, req)
	if err != nil {
		return nil, fmt.Errorf("GetPositions: %w", err)
	}

	positions := make([]PositionDetail, 0, len(rsp.PositionList))
	for _, p := range rsp.PositionList {
		if p == nil {
			continue
		}
		positions = append(positions, PositionDetail{
			Code:        p.Code,
			Name:        p.Name,
			Qty:         p.Qty,
			CostPrice:   p.CostPrice,
			CostBalance: p.CostPrice * p.Qty,
			Market:      constant.TrdMarket(p.TrdMarket),
		})
	}
	return positions, nil
}
