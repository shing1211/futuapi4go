package trd

import (
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/util"
)

type OrderBuilder struct {
	req           *PlaceOrderRequest
	marketAutoSet bool
}

func NewOrder(accID uint64, market constant.TrdMarket, env constant.TrdEnv) *OrderBuilder {
	return &OrderBuilder{
		req: &PlaceOrderRequest{
			AccID:     accID,
			TrdMarket: market,
			TrdEnv:    env,
			OrderType: constant.OrderType_Normal,
		},
	}
}

func (b *OrderBuilder) Buy(code string, qty float64) *OrderBuilder {
	b.req.TrdSide = constant.TrdSide_Buy
	b.req.Code = code
	b.req.Qty = qty
	return b
}

func (b *OrderBuilder) Sell(code string, qty float64) *OrderBuilder {
	b.req.TrdSide = constant.TrdSide_Sell
	b.req.Code = code
	b.req.Qty = qty
	return b
}

func (b *OrderBuilder) At(price float64) *OrderBuilder {
	b.req.Price = price
	return b
}

func (b *OrderBuilder) Market() *OrderBuilder {
	b.req.OrderType = constant.OrderType_Market
	b.req.Price = 0
	return b
}

func (b *OrderBuilder) WithRemark(remark string) *OrderBuilder {
	b.req.Remark = remark
	return b
}

func (b *OrderBuilder) AutoDetectMarket() *OrderBuilder {
	if b.req.Code == "" {
		return b
	}
	trdMarket, secMarket := util.DetectTradingMarkets(b.req.Code)
	b.req.TrdMarket = trdMarket
	b.req.SecMarket = secMarket
	b.marketAutoSet = true
	return b
}

func (b *OrderBuilder) WithSecMarket(secMarket constant.TrdSecMarket) *OrderBuilder {
	b.req.SecMarket = secMarket
	return b
}

func (b *OrderBuilder) Build() (*PlaceOrderRequest, error) {
	if b.req.Code == "" {
		return nil, constant.NewFutuError(constant.ErrCodeInvalidParams, "OrderBuilder.Build", "stock code is required")
	}
	if b.req.Qty <= 0 {
		return nil, constant.NewFutuError(constant.ErrCodeInvalidParams, "OrderBuilder.Build", "quantity must be positive")
	}
	if !b.marketAutoSet && b.req.TrdMarket == 0 {
		return nil, constant.NewFutuError(constant.ErrCodeInvalidParams, "OrderBuilder.Build", "TrdMarket not set; call AutoDetectMarket() or set TrdMarket explicitly")
	}
	return b.req, nil
}

func (b *OrderBuilder) WithTimeInForce(tif constant.TimeInForce) *OrderBuilder {
	b.req.TimeInForce = int32(tif)
	return b
}

func (b *OrderBuilder) WithFillOutsideRTH(outside bool) *OrderBuilder {
	b.req.FillOutsideRTH = outside
	return b
}

func (b *OrderBuilder) WithAuxPrice(auxPrice float64) *OrderBuilder {
	b.req.AuxPrice = auxPrice
	return b
}
