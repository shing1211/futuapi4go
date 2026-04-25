package constant

import "fmt"

var (
	ErrNilRequest   = &FutuError{Code: ErrCodeInvalidParams, Message: "request is nil"}
	ErrInvalidAccID = &FutuError{Code: ErrCodeInvalidParams, Message: "account ID is required"}
	ErrInvalidPrice = &FutuError{Code: ErrCodeInvalidParams, Message: "price must be positive"}
	ErrInvalidQty   = &FutuError{Code: ErrCodeInvalidParams, Message: "quantity must be positive"}
	ErrInvalidCode  = &FutuError{Code: ErrCodeInvalidParams, Message: "stock code is required"}
)

type PlaceOrderRequest interface {
	GetAccID() uint64
	GetCode() string
	GetPrice() float64
	GetQty() float64
}

func validatePlaceOrder(req PlaceOrderRequest) error {
	if req == nil {
		return ErrNilRequest
	}
	if req.GetAccID() == 0 {
		return ErrInvalidAccID
	}
	if req.GetCode() == "" {
		return ErrInvalidCode
	}
	if req.GetPrice() < 0 {
		return ErrInvalidPrice
	}
	if req.GetQty() <= 0 {
		return ErrInvalidQty
	}
	return nil
}

type AccIDRequest interface {
	GetAccID() uint64
}

func validateAccIDRequest(req AccIDRequest) error {
	if req == nil {
		return ErrNilRequest
	}
	if req.GetAccID() == 0 {
		return ErrInvalidAccID
	}
	return nil
}

type OrderIDRequest interface {
	GetAccID() uint64
	GetOrderID() uint64
}

func validateOrderIDRequest(req OrderIDRequest) error {
	if req == nil {
		return ErrNilRequest
	}
	if req.GetAccID() == 0 {
		return ErrInvalidAccID
	}
	if req.GetOrderID() == 0 {
		return &FutuError{Code: ErrCodeInvalidParams, Message: "order ID is required"}
	}
	return nil
}

func wrapValidationError(prefix string, vErr error) error {
	if fe, ok := vErr.(*FutuError); ok {
		return &FutuError{
			Code:    fe.Code,
			Message: fe.Message,
			Func:    prefix,
		}
	}
	return fmt.Errorf("%s: %w", prefix, vErr)
}

var lotSizeMap = map[TrdMarket]float64{
	TrdMarket_HK:   100,
	TrdMarket_US:   1,
	TrdMarket_CN:   100,
	TrdMarket_HKCC: 100,
}

func LotSize(market TrdMarket) (float64, bool) {
	lot, ok := lotSizeMap[market]
	return lot, ok
}

func PriceTick(market TrdMarket) float64 {
	switch market {
	case TrdMarket_HK, TrdMarket_HKCC:
		return 0.01
	case TrdMarket_US:
		return 0.01
	case TrdMarket_CN:
		return 0.01
	default:
		return 0.01
	}
}
