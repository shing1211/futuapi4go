package constant

import "testing"

func TestErrorCodes(t *testing.T) {
	tests := []struct {
		name string
		code ErrorCode
	}{
		{"Success", ErrCodeSuccess},
		{"InvalidParams", ErrCodeInvalidParams},
		{"Timeout", ErrCodeTimeout},
		{"NetworkError", ErrCodeNetworkError},
		{"ServerBusy", ErrCodeServerBusy},
		{"Disconnected", ErrCodeDisconnected},
		{"AccNotFound", ErrCodeAccNotFound},
		{"InsufficientBalance", ErrCodeInsufficientBalance},
		{"MarketClosed", ErrCodeMarketClosed},
		{"OrderRejected", ErrCodeOrderRejected},
		{"AlreadySubbed", ErrCodeAlreadySubbed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.code == 0 && tt.name != "Success" {
				t.Errorf("error code %s should not be zero", tt.name)
			}
		})
	}
}

func TestIsInsufficientBalance(t *testing.T) {
	err := &FutuError{Code: ErrCodeInsufficientBalance, Message: "insufficient balance"}
	if !IsInsufficientBalance(err) {
		t.Error("expected IsInsufficientBalance to return true")
	}
}

func TestIsMarketClosed(t *testing.T) {
	err := &FutuError{Code: ErrCodeMarketClosed, Message: "market closed"}
	if !IsMarketClosed(err) {
		t.Error("expected IsMarketClosed to return true")
	}
}

func TestIsOrderRejected(t *testing.T) {
	err := &FutuError{Code: ErrCodeOrderRejected, Message: "order rejected"}
	if !IsOrderRejected(err) {
		t.Error("expected IsOrderRejected to return true")
	}
}

func TestIsNetworkError(t *testing.T) {
	tests := []ErrorCode{ErrCodeNetworkError, ErrCodeProtocolErr}
	for _, code := range tests {
		err := &FutuError{Code: code, Message: "network error"}
		if !IsNetworkError(err) {
			t.Errorf("expected IsNetworkError to return true for code %d", code)
		}
	}
}

func TestIsServerBusy(t *testing.T) {
	err := &FutuError{Code: ErrCodeServerBusy, Message: "server busy"}
	if !IsServerBusy(err) {
		t.Error("expected IsServerBusy to return true")
	}
}

func TestIsAccountError(t *testing.T) {
	tests := []ErrorCode{ErrCodeAccNotFound, ErrCodeAccDisabled, ErrCodeAccLocked, ErrCodeAccAuthFail}
	for _, code := range tests {
		err := &FutuError{Code: code, Message: "account error"}
		if !IsAccountError(err) {
			t.Errorf("expected IsAccountError to return true for code %d", code)
		}
	}
}

func TestIsSubscriptionError(t *testing.T) {
	err := &FutuError{Code: ErrCodeAlreadySubbed, Message: "already subscribed"}
	if !IsSubscriptionError(err) {
		t.Error("expected IsSubscriptionError to return true")
	}
}