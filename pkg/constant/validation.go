package constant

import "fmt"

var (
	ErrNilRequest    = &FutuError{Code: ErrCodeInvalidParams, Message: "request is nil", Category: CategoryAPI}
	ErrInvalidAccID  = &FutuError{Code: ErrCodeInvalidParams, Message: "account ID is required", Category: CategoryAPI}
	ErrInvalidPrice  = &FutuError{Code: ErrCodeInvalidParams, Message: "price must be positive", Category: CategoryAPI}
	ErrInvalidQty    = &FutuError{Code: ErrCodeInvalidParams, Message: "quantity must be positive", Category: CategoryAPI}
	ErrInvalidCode   = &FutuError{Code: ErrCodeInvalidParams, Message: "stock code is required", Category: CategoryAPI}
	ErrCodeTooLong   = &FutuError{Code: ErrCodeInvalidParams, Message: "stock code exceeds 32 characters", Category: CategoryAPI}
	ErrQtyTooLarge   = &FutuError{Code: ErrCodeQtyTooLarge, Message: "quantity exceeds maximum (10,000,000)", Category: CategoryTrading}
	ErrPriceTooLarge = &FutuError{Code: ErrCodePriceOutOfRange, Message: "price exceeds maximum (1,000,000)", Category: CategoryTrading}
	ErrRemarkTooLong = &FutuError{Code: ErrCodeInvalidParams, Message: "remark exceeds 256 characters", Category: CategoryAPI}
	ErrInvalidMarket = &FutuError{Code: ErrCodeInvalidParams, Message: "invalid market", Category: CategoryAPI}
)

// MaxCodeLen is the maximum allowed length for a stock code string.
const (
	MaxCodeLen   = 32
	MaxRemarkLen = 256
	MaxQty       = 10_000_000.0
	MaxPrice     = 1_000_000.0
	MinQty       = 0.001
)

// ValidateAccID returns ErrInvalidAccID if accID is zero.
func ValidateAccID(accID uint64) error {
	if accID == 0 {
		return ErrInvalidAccID
	}
	return nil
}

// ValidateCode returns ErrInvalidCode if code is empty or ErrCodeTooLong if it
// exceeds MaxCodeLen characters.
func ValidateCode(code string) error {
	if code == "" {
		return ErrInvalidCode
	}
	if len(code) > MaxCodeLen {
		return ErrCodeTooLong
	}
	return nil
}

// ValidateQty returns ErrInvalidQty if qty is not positive or ErrQtyTooLarge if
// it exceeds MaxQty.
func ValidateQty(qty float64) error {
	if qty <= 0 {
		return ErrInvalidQty
	}
	if qty > MaxQty {
		return ErrQtyTooLarge
	}
	return nil
}

// ValidatePrice returns ErrInvalidPrice if price is negative or ErrPriceTooLarge
// if it exceeds MaxPrice. A price of zero is allowed (for market orders).
func ValidatePrice(price float64) error {
	if price < 0 {
		return ErrInvalidPrice
	}
	if price > MaxPrice {
		return ErrPriceTooLarge
	}
	return nil
}

// ValidateRemark returns ErrRemarkTooLong if remark exceeds MaxRemarkLen characters.
func ValidateRemark(remark string) error {
	if len(remark) > MaxRemarkLen {
		return ErrRemarkTooLong
	}
	return nil
}

// PlaceOrderRequest is the validation interface for place-order requests.
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

// AccIDRequest is the validation interface for requests that only need an account ID.
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

// OrderIDRequest is the validation interface for requests that need an account ID
// and an order ID.
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

// LotSize returns the lot size for the given trading market. Returns (0, false) if
// the market is not in the known map.
func LotSize(market TrdMarket) (float64, bool) {
	lot, ok := lotSizeMap[market]
	return lot, ok
}

// PriceTick returns the minimum price increment (tick size) for the given trading market.
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
