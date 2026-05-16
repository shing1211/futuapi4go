package trd

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

func defaultPlaceOrderRequest() *PlaceOrderRequest {
	return &PlaceOrderRequest{
		AccID:     12345,
		Code:      "US.AAPL",
		TrdSide:   constant.TrdSide_Buy,
		OrderType: constant.OrderType_Normal,
		Price:     150.0,
		Qty:       100,
	}
}

func TestValidateOrderNilInput(t *testing.T) {
	warnings := ValidateOrder(nil)
	if len(warnings) != 0 {
		t.Errorf("expected 0 warnings for nil input, got %d", len(warnings))
	}

	warnings = ValidateOrder(&OrderValidationInput{Order: nil})
	if len(warnings) != 0 {
		t.Errorf("expected 0 warnings for nil order, got %d", len(warnings))
	}
}

func TestValidateOrderMarketClosed(t *testing.T) {
	req := defaultPlaceOrderRequest()
	input := &OrderValidationInput{
		Order:      req,
		MarketOpen: false,
	}
	warnings := ValidateOrder(input)
	if len(warnings) == 0 {
		t.Fatal("expected warnings for closed market")
	}
	if warnings[0].Severity != SeverityError {
		t.Errorf("expected error severity, got %s", warnings[0].Severity)
	}
}

func TestValidateOrderInsufficientBuyingPower(t *testing.T) {
	req := defaultPlaceOrderRequest()
	// Price=150, Qty=100 → cost=15000, buying power=10000
	input := &OrderValidationInput{
		Order:       req,
		MarketOpen:  true,
		BuyingPower: 10000,
	}
	warnings := ValidateOrder(input)

	found := false
	for _, w := range warnings {
		if w.Field == "buying_power" && w.Severity == SeverityError {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected buying_power error warning")
	}
}

func TestValidateOrderQtyExceedsMaxBuy(t *testing.T) {
	req := defaultPlaceOrderRequest()
	// Qty=100, MaxBuyQty=50
	input := &OrderValidationInput{
		Order:      req,
		MarketOpen: true,
		MaxBuyQty:  50,
	}
	warnings := ValidateOrder(input)

	found := false
	for _, w := range warnings {
		if w.Field == "qty" && w.Severity == SeverityError {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected qty error for exceeding max buy qty")
	}
}

func TestValidateOrderQtyExceedsMaxSell(t *testing.T) {
	req := defaultPlaceOrderRequest()
	req.TrdSide = constant.TrdSide_Sell
	input := &OrderValidationInput{
		Order:      req,
		MarketOpen: true,
		MaxSellQty: 50,
	}
	warnings := ValidateOrder(input)

	found := false
	for _, w := range warnings {
		if w.Field == "qty" && w.Severity == SeverityError {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected qty error for exceeding max sell qty")
	}
}

func TestValidateOrderZeroQty(t *testing.T) {
	req := defaultPlaceOrderRequest()
	req.Qty = 0
	input := &OrderValidationInput{
		Order:      req,
		MarketOpen: true,
	}
	warnings := ValidateOrder(input)

	found := false
	for _, w := range warnings {
		if w.Field == "qty" && w.Severity == SeverityError {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected qty error for zero quantity")
	}

	req.Qty = -1
	warnings = ValidateOrder(input)
	found = false
	for _, w := range warnings {
		if w.Field == "qty" && w.Severity == SeverityError {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected qty error for negative quantity")
	}
}

func TestValidateOrderNonMarketZeroPrice(t *testing.T) {
	req := defaultPlaceOrderRequest()
	req.Price = 0
	input := &OrderValidationInput{
		Order:      req,
		MarketOpen: true,
	}
	warnings := ValidateOrder(input)

	found := false
	for _, w := range warnings {
		if w.Field == "price" && w.Severity == SeverityWarning {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected price warning for zero price on non-market order")
	}

	hasError := false
	for _, w := range warnings {
		if w.Severity == SeverityError {
			hasError = true
			break
		}
	}
	if hasError {
		t.Error("zero price on non-market order should be warning, not error")
	}
}

func TestHasErrors(t *testing.T) {
	warnings := []ValidationWarning{
		{Field: "qty", Severity: SeverityWarning, Message: "test warning"},
		{Field: "price", Severity: SeverityError, Message: "test error"},
	}
	if !HasErrors(warnings) {
		t.Error("expected HasErrors to return true with mixed warnings")
	}
}

func TestHasErrorsNoErrors(t *testing.T) {
	warnings := []ValidationWarning{
		{Field: "price", Severity: SeverityWarning, Message: "test warning"},
	}
	if HasErrors(warnings) {
		t.Error("expected HasErrors to return false with only warnings")
	}
}
