package trd

import (
	"fmt"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

type ValidationSeverity string

const (
	SeverityError   ValidationSeverity = "error"
	SeverityWarning ValidationSeverity = "warning"
)

type ValidationWarning struct {
	Field    string
	Severity ValidationSeverity
	Message  string
}

func (v ValidationWarning) Error() string {
	return fmt.Sprintf("[%s] %s: %s", v.Severity, v.Field, v.Message)
}

type OrderValidationInput struct {
	Order       *PlaceOrderRequest
	MarketOpen  bool
	BuyingPower float64
	MaxBuyQty   float64
	MaxSellQty  float64
}

func ValidateOrder(input *OrderValidationInput) []ValidationWarning {
	if input == nil || input.Order == nil {
		return nil
	}
	req := input.Order
	var warnings []ValidationWarning

	if !input.MarketOpen {
		warnings = append(warnings, ValidationWarning{
			Field:    "market",
			Severity: SeverityError,
			Message:  "market is not open for trading",
		})
	}

	if req.TrdSide == constant.TrdSide_Buy && input.BuyingPower > 0 {
		estimatedCost := req.Price * req.Qty
		if estimatedCost > input.BuyingPower {
			warnings = append(warnings, ValidationWarning{
				Field:    "buying_power",
				Severity: SeverityError,
				Message:  fmt.Sprintf("estimated cost %.2f exceeds buying power %.2f", estimatedCost, input.BuyingPower),
			})
		}
	}

	if input.MaxBuyQty > 0 && req.TrdSide == constant.TrdSide_Buy && req.Qty > input.MaxBuyQty {
		warnings = append(warnings, ValidationWarning{
			Field:    "qty",
			Severity: SeverityError,
			Message:  fmt.Sprintf("buy qty %.0f exceeds max allowed %.0f", req.Qty, input.MaxBuyQty),
		})
	}

	if input.MaxSellQty > 0 && req.TrdSide == constant.TrdSide_Sell && req.Qty > input.MaxSellQty {
		warnings = append(warnings, ValidationWarning{
			Field:    "qty",
			Severity: SeverityError,
			Message:  fmt.Sprintf("sell qty %.0f exceeds max allowed %.0f", req.Qty, input.MaxSellQty),
		})
	}

	if req.Qty <= 0 {
		warnings = append(warnings, ValidationWarning{
			Field:    "qty",
			Severity: SeverityError,
			Message:  "quantity must be positive",
		})
	}

	if req.Price <= 0 && req.OrderType != constant.OrderType_Market && req.OrderType != constant.OrderType_AbsoluteLimit {
		warnings = append(warnings, ValidationWarning{
			Field:    "price",
			Severity: SeverityWarning,
			Message:  "price is zero or negative for a non-market order",
		})
	}

	return warnings
}

func HasErrors(warnings []ValidationWarning) bool {
	for _, w := range warnings {
		if w.Severity == SeverityError {
			return true
		}
	}
	return false
}
