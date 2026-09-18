package constant

import (
	"errors"
	"fmt"
	"testing"
)

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

func TestFullMessage(t *testing.T) {
	err := NewFutuError(ErrCodeTimeout, "GetKL", "request timed out")
	msg := err.FullMessage()
	if msg == "" {
		t.Error("FullMessage should not be empty")
	}
	if !contains(msg, "ErrCodeTimeout") {
		t.Error("FullMessage should contain code name")
	}
	if !contains(msg, "GetKL") {
		t.Error("FullMessage should contain function name")
	}
	if !contains(msg, "Category: timeout") {
		t.Error("FullMessage should contain category")
	}
	if !contains(msg, "Recovery:") {
		t.Error("FullMessage should contain recovery hint")
	}
}

func TestFullMessageWithWrap(t *testing.T) {
	inner := errors.New("connection refused")
	err := NewFutuErrorWithWrap(ErrCodeNetworkError, "Connect", "dial failed", inner)
	msg := err.FullMessage()
	if !contains(msg, "Caused by:") {
		t.Error("FullMessage should contain inner error")
	}
	if !contains(msg, "connection refused") {
		t.Error("FullMessage should contain inner error message")
	}
}

func TestIsMethod(t *testing.T) {
	timeoutErr := NewFutuError(ErrCodeTimeout, "GetKL", "timed out")
	otherTimeout := NewFutuError(ErrCodeTimeout, "GetRT", "timed out")
	networkErr := NewFutuError(ErrCodeNetworkError, "Connect", "network error")

	if !errors.Is(timeoutErr, otherTimeout) {
		t.Error("errors.Is should match FutuErrors with same code")
	}
	if errors.Is(timeoutErr, networkErr) {
		t.Error("errors.Is should not match FutuErrors with different codes")
	}
}

func TestCodeString(t *testing.T) {
	tests := []struct {
		code ErrorCode
		want string
	}{
		{ErrCodeSuccess, "ErrCodeSuccess"},
		{ErrCodeTimeout, "ErrCodeTimeout"},
		{ErrCodeNetworkError, "ErrCodeNetworkError"},
		{ErrCodeAccNotFound, "ErrCodeAccNotFound"},
		{ErrCodeInsufficientBalance, "ErrCodeInsufficientBalance"},
		{ErrCodeAlreadySubbed, "ErrCodeAlreadySubbed"},
		{ErrorCode(-9999), "ErrCode_-9999"},
	}
	for _, tt := range tests {
		err := &FutuError{Code: tt.code}
		got := err.CodeString()
		if got != tt.want {
			t.Errorf("CodeString() = %q, want %q", got, tt.want)
		}
	}
}

func TestIsServerError(t *testing.T) {
	err := NewFutuError(ErrCodeServerBusy, "GetKL", "busy")
	if !IsServerError(err) {
		t.Error("expected IsServerError to return true for ServerBusy")
	}
	err2 := NewFutuError(ErrCodeUnknown, "GetKL", "unknown")
	if !IsServerError(err2) {
		t.Error("expected IsServerError to return true for Unknown")
	}
	err3 := NewFutuError(ErrCodeTimeout, "GetKL", "timeout")
	if IsServerError(err3) {
		t.Error("expected IsServerError to return false for Timeout")
	}
}

func TestIsAPIError(t *testing.T) {
	err := NewFutuError(ErrCodeInvalidParams, "GetKL", "bad params")
	if !IsAPIError(err) {
		t.Error("expected IsAPIError to return true for InvalidParams")
	}
	err2 := NewFutuError(ErrCodeTimeout, "GetKL", "timeout")
	if IsAPIError(err2) {
		t.Error("expected IsAPIError to return false for Timeout")
	}
}

func TestCategoryOf(t *testing.T) {
	tests := []struct {
		code     ErrorCode
		category ErrorCategory
	}{
		{ErrCodeTimeout, CategoryTimeout},
		{ErrCodeDisconnected, CategoryConnection},
		{ErrCodeNetworkError, CategoryConnection},
		{ErrCodeAccNotFound, CategoryAccount},
		{ErrCodeInsufficientBalance, CategoryTrading},
		{ErrCodeAlreadySubbed, CategorySubscribe},
		{ErrCodeInvalidParams, CategoryAPI},
	}
	for _, tt := range tests {
		err := NewFutuError(tt.code, "TestFunc", "msg")
		got := CategoryOf(err)
		if got != tt.category {
			t.Errorf("CategoryOf(code=%d) = %q, want %q", tt.code, got, tt.category)
		}
	}
}

func TestRecoveryHint(t *testing.T) {
	err := NewFutuError(ErrCodeTimeout, "GetKL", "timed out")
	hint := RecoveryHint(err)
	if hint == "" {
		t.Error("RecoveryHint should not be empty")
	}
	if hint != codeToRecovery[ErrCodeTimeout] {
		t.Errorf("RecoveryHint = %q, want %q", hint, codeToRecovery[ErrCodeTimeout])
	}
}

func TestIsConnectionError(t *testing.T) {
	err := NewFutuError(ErrCodeDisconnected, "Connect", "lost")
	if !IsConnectionError(err) {
		t.Error("expected IsConnectionError to return true for Disconnected")
	}
	err2 := NewFutuError(ErrCodeTimeout, "GetKL", "timeout")
	if IsConnectionError(err2) {
		t.Error("expected IsConnectionError to return false for Timeout")
	}
}

func TestIsTradingError(t *testing.T) {
	err := NewFutuError(ErrCodeMarketClosed, "PlaceOrder", "closed")
	if !IsTradingError(err) {
		t.Error("expected IsTradingError to return true for MarketClosed")
	}
	err2 := NewFutuError(ErrCodeTimeout, "GetKL", "timeout")
	if IsTradingError(err2) {
		t.Error("expected IsTradingError to return false for Timeout")
	}
}

func TestNewFutuError(t *testing.T) {
	err := NewFutuError(ErrCodeTimeout, "GetKL", "timed out")
	if err.Code != ErrCodeTimeout {
		t.Errorf("Code = %d, want %d", err.Code, ErrCodeTimeout)
	}
	if err.Func != "GetKL" {
		t.Errorf("Func = %q, want %q", err.Func, "GetKL")
	}
	if err.Category != CategoryTimeout {
		t.Errorf("Category = %q, want %q", err.Category, CategoryTimeout)
	}
	if err.Recovery == "" {
		t.Error("Recovery should not be empty")
	}
}

func TestNewFutuErrorWithWrap(t *testing.T) {
	inner := errors.New("dial: connection refused")
	err := NewFutuErrorWithWrap(ErrCodeNetworkError, "Connect", "dial failed", inner)
	if err.Err != inner {
		t.Error("inner error should be preserved")
	}
	if err.Category != CategoryConnection {
		t.Errorf("Category = %q, want %q", err.Category, CategoryConnection)
	}
	unwrapped := errors.Unwrap(err)
	if unwrapped != inner {
		t.Error("Unwrap should return inner error")
	}
}

func TestFutuErrorFormatting(t *testing.T) {
	err := &FutuError{Code: ErrCodeTimeout, Message: "timed out", Func: "GetKL"}
	s := fmt.Sprintf("%v", err)
	if !contains(s, "GetKL") {
		t.Error("Error() output should contain function name")
	}
	if !contains(s, "timed out") {
		t.Error("Error() output should contain message")
	}
}

func TestFutuErrorRecoverable(t *testing.T) {
	tests := []struct {
		code        ErrorCode
		recoverable bool
	}{
		{ErrCodeTimeout, true},
		{ErrCodeDisconnected, true},
		{ErrCodeNetworkError, true},
		{ErrCodeProtocolErr, true},
		{ErrCodeInvalidParams, false},
		{ErrCodeServerBusy, false},
		{ErrCodeAccNotFound, false},
		{ErrCodeAccDisabled, false},
		{ErrCodeInsufficientBalance, false},
		{ErrCodeMarketClosed, false},
		{ErrCodeOrderRejected, false},
		{ErrCodeTradingDisabled, false},
		{ErrCodeAlreadySubbed, false},
		{ErrCodeNotSubbed, false},
	}

	for _, tt := range tests {
		err := &FutuError{Code: tt.code, Message: "test"}
		got := err.Recoverable()
		if got != tt.recoverable {
			t.Errorf("Recoverable() for code=%d = %v, want %v", tt.code, got, tt.recoverable)
		}
	}
}

func TestIsRecoverable(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		recoverable bool
	}{
		{"nil error", nil, false},
		{"timeout error", &FutuError{Code: ErrCodeTimeout, Message: "test"}, true},
		{"disconnected error", &FutuError{Code: ErrCodeDisconnected, Message: "test"}, true},
		{"network error", &FutuError{Code: ErrCodeNetworkError, Message: "test"}, true},
		{"invalid params", &FutuError{Code: ErrCodeInvalidParams, Message: "test"}, false},
		{"trading error", &FutuError{Code: ErrCodeOrderRejected, Message: "test"}, false},
		{"account error", &FutuError{Code: ErrCodeAccNotFound, Message: "test"}, false},
		{"non-FutuError", errors.New("plain error"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsRecoverable(tt.err)
			if got != tt.recoverable {
				t.Errorf("IsRecoverable() = %v, want %v", got, tt.recoverable)
			}
		})
	}
}

func TestWrapError(t *testing.T) {
	tests := []struct {
		name        string
		retType     int32
		wantCode    ErrorCode
	}{
		{"success", 0, ErrCodeSuccess},
		{"invalid params", -1, ErrCodeInvalidParams},
		{"timeout", -100, ErrCodeTimeout},
		{"network error", -101, ErrCodeNetworkError},
		{"protocol error", -102, ErrCodeProtocolErr},
		{"server busy", -103, ErrCodeServerBusy},
		{"disconnected", -200, ErrCodeDisconnected},
		{"acc not found", -201, ErrCodeAccNotFound},
		{"acc disabled", -202, ErrCodeAccDisabled},
		{"acc locked", -203, ErrCodeAccLocked},
		{"acc auth fail", -204, ErrCodeAccAuthFail},
		{"insufficient balance", -301, ErrCodeInsufficientBalance},
		{"market closed", -302, ErrCodeMarketClosed},
		{"order rejected", -303, ErrCodeOrderRejected},
		{"price out of range", -304, ErrCodePriceOutOfRange},
		{"qty too large", -305, ErrCodeQtyTooLarge},
		{"trading disabled", -306, ErrCodeTradingDisabled},
		{"invalid security", -307, ErrCodeInvalidSecurity},
		{"no permission", -308, ErrCodeNoPermission},
		{"unknown", -400, ErrCodeUnknown},
		{"already subbed", -401, ErrCodeAlreadySubbed},
		{"not subbed", -402, ErrCodeNotSubbed},
		{"positive retType", 1, ErrCodeUnknown},
		{"unknown negative", -9999, ErrorCode(-9999)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WrapError("TestFunc", tt.retType, "test message")
			if err.Code != tt.wantCode {
				t.Errorf("WrapError code = %d, want %d", err.Code, tt.wantCode)
			}
			if err.Func != "TestFunc" {
				t.Errorf("WrapError Func = %q, want %q", err.Func, "TestFunc")
			}
			if err.Message != "test message" {
				t.Errorf("WrapError Message = %q, want %q", err.Message, "test message")
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsStr(s, substr))
}

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
