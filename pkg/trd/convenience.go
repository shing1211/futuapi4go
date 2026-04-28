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

// CancelAllOrdersResult holds the result of cancelling all orders for an account.
type CancelAllOrdersResult struct {
	AccID     uint64
	Cancelled int32
}

// CancelAllOrders cancels all open orders for the given account, market, and
// environment.
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

// QuickBuy places a limit buy order in a single call.
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

// QuickSell places a limit sell order in a single call.
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

// QuickMarketBuy places a market buy order (no price) in a single call.
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

// QuickMarketSell places a market sell order (no price) in a single call.
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

// PositionDetail is a simplified view of a single position.
type PositionDetail struct {
	Code        string
	Name        string
	Qty         float64
	CostPrice   float64
	CostBalance float64
	Market      constant.TrdMarket
}

// GetPositions returns a simplified list of all positions for the given account.
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

// GetTodayFills returns today's order fills for the given account, market, and
// environment.
func GetTodayFills(ctx context.Context, c *futuapi.Client, accID uint64, market constant.TrdMarket, env constant.TrdEnv) (*GetOrderFillListResponse, error) {
	req := &GetOrderFillListRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    env,
	}
	return GetOrderFillList(ctx, c, req)
}

// GetTodayOrders returns today's orders for the given account, market, and
// environment.
func GetTodayOrders(ctx context.Context, c *futuapi.Client, accID uint64, market constant.TrdMarket, env constant.TrdEnv) (*GetOrderListResponse, error) {
	req := &GetOrderListRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    env,
	}
	return GetOrderList(ctx, c, req)
}

// GetAccountFunds returns the account funds (buying power, cash, etc.) for the
// given account, market, and environment.
func GetAccountFunds(ctx context.Context, c *futuapi.Client, accID uint64, market constant.TrdMarket, env constant.TrdEnv) (*GetFundsResponse, error) {
	req := &GetFundsRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    env,
	}
	return GetFunds(ctx, c, req)
}
