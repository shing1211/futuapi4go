package trd

import (
	"github.com/shing1211/futuapi4go/pkg/constant"
)

type OrderBuilder struct {
	req *PlaceOrderRequest
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

func (b *OrderBuilder) Build() *PlaceOrderRequest {
	return b.req
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