package trd

import (
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/util"
)

// OrderBuilder provides a fluent API for constructing a PlaceOrderRequest.
type OrderBuilder struct {
	req           *PlaceOrderRequest
	marketAutoSet bool // true when AutoDetectMarket has been called
}

// NewOrder creates an OrderBuilder initialised with the given account, market, and
// environment. The order type defaults to OrderType_Normal (limit order).
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

// Buy sets the order side to buy with the given stock code and quantity.
func (b *OrderBuilder) Buy(code string, qty float64) *OrderBuilder {
	b.req.TrdSide = constant.TrdSide_Buy
	b.req.Code = code
	b.req.Qty = qty
	return b
}

// Sell sets the order side to sell with the given stock code and quantity.
func (b *OrderBuilder) Sell(code string, qty float64) *OrderBuilder {
	b.req.TrdSide = constant.TrdSide_Sell
	b.req.Code = code
	b.req.Qty = qty
	return b
}

// At sets the limit price for the order.
func (b *OrderBuilder) At(price float64) *OrderBuilder {
	b.req.Price = price
	return b
}

// Market changes the order type to market order and clears the price.
func (b *OrderBuilder) Market() *OrderBuilder {
	b.req.OrderType = constant.OrderType_Market
	b.req.Price = 0
	return b
}

// WithRemark attaches a user-defined remark to the order.
func (b *OrderBuilder) WithRemark(remark string) *OrderBuilder {
	b.req.Remark = remark
	return b
}

// AutoDetectMarket infers TrdMarket and SecMarket from the stock code prefix
// (e.g. "HK.00700" → TrdMarket_HK). Must be called after Buy or Sell so the
// code is already set. No-op if the code is empty.
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

// WithSecMarket explicitly sets the secondary market on the order request.
func (b *OrderBuilder) WithSecMarket(secMarket constant.TrdSecMarket) *OrderBuilder {
	b.req.SecMarket = secMarket
	return b
}

// Build validates the builder state and returns the assembled PlaceOrderRequest.
// Returns an error if the code is empty, quantity is not positive, or TrdMarket
// is unset and AutoDetectMarket was not called.
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

// WithTimeInForce sets the time-in-force policy for the order.
func (b *OrderBuilder) WithTimeInForce(tif constant.TimeInForce) *OrderBuilder {
	b.req.TimeInForce = int32(tif)
	return b
}

// WithFillOutsideRTH enables or disables filling outside regular trading hours.
func (b *OrderBuilder) WithFillOutsideRTH(outside bool) *OrderBuilder {
	b.req.FillOutsideRTH = outside
	return b
}

// WithAuxPrice sets the auxiliary price (e.g. stop price for stop-limit orders).
func (b *OrderBuilder) WithAuxPrice(auxPrice float64) *OrderBuilder {
	b.req.AuxPrice = auxPrice
	return b
}
