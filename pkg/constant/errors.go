package constant

import (
	"errors"
	"fmt"
	"strings"
)

// ErrorCode represents a Futu API error code. Negative values indicate errors;
// zero indicates success.
type ErrorCode int32

const (
	ErrCodeSuccess       ErrorCode = 0
	ErrCodeInvalidParams ErrorCode = -1
	ErrCodeTimeout       ErrorCode = -100
	ErrCodeDisconnected  ErrorCode = -200
	ErrCodeUnknown       ErrorCode = -400

	ErrCodeNetworkError ErrorCode = -101
	ErrCodeProtocolErr  ErrorCode = -102
	ErrCodeServerBusy   ErrorCode = -103

	ErrCodeAccNotFound ErrorCode = -201
	ErrCodeAccDisabled ErrorCode = -202
	ErrCodeAccLocked   ErrorCode = -203
	ErrCodeAccAuthFail ErrorCode = -204

	ErrCodeInsufficientBalance ErrorCode = -301
	ErrCodeMarketClosed        ErrorCode = -302
	ErrCodeOrderRejected       ErrorCode = -303
	ErrCodePriceOutOfRange     ErrorCode = -304
	ErrCodeQtyTooLarge         ErrorCode = -305
	ErrCodeTradingDisabled     ErrorCode = -306
	ErrCodeInvalidSecurity     ErrorCode = -307
	ErrCodeNoPermission        ErrorCode = -308

	ErrCodeAlreadySubbed ErrorCode = -401
	ErrCodeNotSubbed     ErrorCode = -402
)

// ErrorCategory classifies an error into a broad category such as connection,
// timeout, API, account, trading, or subscribe.
type ErrorCategory string

const (
	CategoryConnection ErrorCategory = "connection"
	CategoryTimeout    ErrorCategory = "timeout"
	CategoryAPI        ErrorCategory = "api"
	CategoryAccount   ErrorCategory = "account"
	CategoryTrading   ErrorCategory = "trading"
	CategorySubscribe ErrorCategory = "subscribe"
	CategoryUnknown   ErrorCategory = "unknown"
)

var codeToCategory = map[ErrorCode]ErrorCategory{
	ErrCodeSuccess:       CategoryAPI,
	ErrCodeInvalidParams: CategoryAPI,
	ErrCodeTimeout:       CategoryTimeout,
	ErrCodeDisconnected:  CategoryConnection,
	ErrCodeUnknown:       CategoryUnknown,
	ErrCodeNetworkError:  CategoryConnection,
	ErrCodeProtocolErr:   CategoryConnection,
	ErrCodeServerBusy:    CategoryAPI,
	ErrCodeAccNotFound:   CategoryAccount,
	ErrCodeAccDisabled:   CategoryAccount,
	ErrCodeAccLocked:     CategoryAccount,
	ErrCodeAccAuthFail:   CategoryAccount,
	ErrCodeInsufficientBalance: CategoryTrading,
	ErrCodeMarketClosed:        CategoryTrading,
	ErrCodeOrderRejected:       CategoryTrading,
	ErrCodePriceOutOfRange:     CategoryTrading,
	ErrCodeQtyTooLarge:         CategoryTrading,
	ErrCodeTradingDisabled:     CategoryTrading,
	ErrCodeInvalidSecurity:     CategoryTrading,
	ErrCodeNoPermission:        CategoryTrading,
	ErrCodeAlreadySubbed:       CategorySubscribe,
	ErrCodeNotSubbed:           CategorySubscribe,
}

var codeToRecovery = map[ErrorCode]string{
	ErrCodeInvalidParams:       "Check function parameters for validity.",
	ErrCodeTimeout:             "Increase timeout or check network connectivity.",
	ErrCodeDisconnected:        "Reconnect to server. Call Connect() first.",
	ErrCodeNetworkError:        "Check network connection and retry.",
	ErrCodeProtocolErr:         "Protocol mismatch. Reconnect to server.",
	ErrCodeServerBusy:          "Server is busy. Retry after a delay.",
	ErrCodeAccNotFound:         "Verify account ID and trading category.",
	ErrCodeAccDisabled:         "Account is disabled. Contact broker.",
	ErrCodeAccLocked:           "Unlock trading with UnlockTrade() first.",
	ErrCodeAccAuthFail:         "Verify trading password.",
	ErrCodeInsufficientBalance: "Check available buying power.",
	ErrCodeMarketClosed:        "Wait for market to open.",
	ErrCodeOrderRejected:       "Check order parameters and market rules.",
	ErrCodePriceOutOfRange:     "Adjust price within allowed range.",
	ErrCodeQtyTooLarge:         "Reduce order quantity.",
	ErrCodeTradingDisabled:     "Trading is not enabled for this account.",
	ErrCodeInvalidSecurity:     "Verify stock code format.",
	ErrCodeNoPermission:        "Check API subscription level.",
	ErrCodeAlreadySubbed:       "Already subscribed. Unsubscribe first if needed.",
	ErrCodeNotSubbed:           "Subscribe to the data feed first.",
}

// errorCodeNames maps ErrorCode values to their symbolic names for display.
var errorCodeNames = map[ErrorCode]string{
	ErrCodeSuccess:             "ErrCodeSuccess",
	ErrCodeInvalidParams:       "ErrCodeInvalidParams",
	ErrCodeTimeout:             "ErrCodeTimeout",
	ErrCodeDisconnected:        "ErrCodeDisconnected",
	ErrCodeUnknown:             "ErrCodeUnknown",
	ErrCodeNetworkError:        "ErrCodeNetworkError",
	ErrCodeProtocolErr:         "ErrCodeProtocolErr",
	ErrCodeServerBusy:          "ErrCodeServerBusy",
	ErrCodeAccNotFound:         "ErrCodeAccNotFound",
	ErrCodeAccDisabled:         "ErrCodeAccDisabled",
	ErrCodeAccLocked:           "ErrCodeAccLocked",
	ErrCodeAccAuthFail:         "ErrCodeAccAuthFail",
	ErrCodeInsufficientBalance: "ErrCodeInsufficientBalance",
	ErrCodeMarketClosed:        "ErrCodeMarketClosed",
	ErrCodeOrderRejected:       "ErrCodeOrderRejected",
	ErrCodePriceOutOfRange:     "ErrCodePriceOutOfRange",
	ErrCodeQtyTooLarge:         "ErrCodeQtyTooLarge",
	ErrCodeTradingDisabled:     "ErrCodeTradingDisabled",
	ErrCodeInvalidSecurity:     "ErrCodeInvalidSecurity",
	ErrCodeNoPermission:        "ErrCodeNoPermission",
	ErrCodeAlreadySubbed:       "ErrCodeAlreadySubbed",
	ErrCodeNotSubbed:           "ErrCodeNotSubbed",
}

// FutuError is the structured error type for all Futu API errors. It carries
// an error code, human-readable message, originating function name, optional
// wrapped inner error, category, and recovery hint.
type FutuError struct {
	Code     ErrorCode
	Message  string
	Func     string
	Err      error
	Category ErrorCategory
	Recovery string
}

func (e *FutuError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (code=%d, inner=%v)", e.Func, e.Message, e.Code, e.Err)
	}
	return fmt.Sprintf("%s: %s (code=%d)", e.Func, e.Message, e.Code)
}

func (e *FutuError) Unwrap() error {
	return e.Err
}

// Is reports whether e matches target. Two FutuErrors match if they share the
// same ErrorCode. It also delegates to errors.Is on the wrapped inner error.
func (e *FutuError) Is(target error) bool {
	if t, ok := target.(*FutuError); ok {
		return t.Code == e.Code
	}
	if e.Err != nil {
		return errors.Is(e.Err, target)
	}
	return false
}

// FullMessage returns a multi-line human-readable description of the error,
// including the code name, category, recovery hint, and inner cause.
func (e *FutuError) FullMessage() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s] %s: %s (code=%d)\n", e.CodeString(), e.Func, e.Message, e.Code))
	if e.Category != "" {
		sb.WriteString(fmt.Sprintf("Category: %s\n", e.Category))
	}
	if e.Recovery != "" {
		sb.WriteString(fmt.Sprintf("Recovery: %s\n", e.Recovery))
	}
	if e.Err != nil {
		sb.WriteString(fmt.Sprintf("Caused by: %v\n", e.Err))
	}
	return sb.String()
}

// CodeString returns the symbolic name of the error code (e.g. "ErrCodeTimeout").
// Falls back to "ErrCode_<number>" for unknown codes.
func (e *FutuError) CodeString() string {
	if name, ok := errorCodeNames[e.Code]; ok {
		return name
	}
	return fmt.Sprintf("ErrCode_%d", e.Code)
}

func getFutuError(err error) (*FutuError, bool) {
	fe, ok := err.(*FutuError)
	return fe, ok
}

// IsTimeout reports whether err is a FutuError with code ErrCodeTimeout.
func IsTimeout(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeTimeout
	}
	return false
}

// IsDisconnected reports whether err is a FutuError with code ErrCodeDisconnected.
func IsDisconnected(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeDisconnected
	}
	return false
}

// IsInvalidParams reports whether err is a FutuError with code ErrCodeInvalidParams.
func IsInvalidParams(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeInvalidParams
	}
	return false
}

// IsSuccess reports whether err is a FutuError with code ErrCodeSuccess.
func IsSuccess(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeSuccess
	}
	return false
}

// IsServerError reports whether err is a FutuError with code ErrCodeServerBusy or ErrCodeUnknown.
func IsServerError(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeServerBusy || fe.Code == ErrCodeUnknown
	}
	return false
}

// IsAPIError reports whether err belongs to the CategoryAPI category.
func IsAPIError(err error) bool {
	return CategoryOf(err) == CategoryAPI
}

// IsNetworkError reports whether err is a FutuError with code ErrCodeNetworkError or ErrCodeProtocolErr.
func IsNetworkError(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeNetworkError || fe.Code == ErrCodeProtocolErr
	}
	return false
}

// IsServerBusy reports whether err is a FutuError with code ErrCodeServerBusy.
func IsServerBusy(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeServerBusy
	}
	return false
}

// IsAccountError reports whether err is a FutuError with an account-related code
// (not found, disabled, locked, or auth failure).
func IsAccountError(err error) bool {
	if fe, ok := getFutuError(err); ok {
		switch fe.Code {
		case ErrCodeAccNotFound, ErrCodeAccDisabled, ErrCodeAccLocked, ErrCodeAccAuthFail:
			return true
		}
	}
	return false
}

// IsInsufficientBalance reports whether err is a FutuError with code ErrCodeInsufficientBalance.
func IsInsufficientBalance(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeInsufficientBalance
	}
	return false
}

// IsMarketClosed reports whether err is a FutuError with code ErrCodeMarketClosed.
func IsMarketClosed(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeMarketClosed
	}
	return false
}

// IsOrderRejected reports whether err is a FutuError with code ErrCodeOrderRejected.
func IsOrderRejected(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeOrderRejected
	}
	return false
}

// IsSubscriptionError reports whether err is a FutuError with code ErrCodeAlreadySubbed or ErrCodeNotSubbed.
func IsSubscriptionError(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeAlreadySubbed || fe.Code == ErrCodeNotSubbed
	}
	return false
}

// IsConnectionError reports whether err belongs to the CategoryConnection category.
func IsConnectionError(err error) bool {
	return CategoryOf(err) == CategoryConnection
}

// IsTimeoutError reports whether err belongs to the CategoryTimeout category.
func IsTimeoutError(err error) bool {
	return CategoryOf(err) == CategoryTimeout
}

// IsTradingError reports whether err belongs to the CategoryTrading category.
func IsTradingError(err error) bool {
	return CategoryOf(err) == CategoryTrading
}

// CategoryOf returns the ErrorCategory for err. It first checks the FutuError's
// own Category field, then falls back to the codeToCategory map, and finally
// recurses into the wrapped inner error. Returns CategoryUnknown if no match.
func CategoryOf(err error) ErrorCategory {
	if fe, ok := getFutuError(err); ok {
		if fe.Category != "" {
			return fe.Category
		}
		if cat, ok := codeToCategory[fe.Code]; ok {
			return cat
		}
	}
	if innerErr := errors.Unwrap(err); innerErr != nil {
		return CategoryOf(innerErr)
	}
	return CategoryUnknown
}

// RecoveryHint returns a human-readable suggestion for how to recover from err.
// It checks the FutuError's Recovery field, then the codeToRecovery map, and
// finally recurses into the wrapped inner error.
func RecoveryHint(err error) string {
	if fe, ok := getFutuError(err); ok {
		if fe.Recovery != "" {
			return fe.Recovery
		}
		if hint, ok := codeToRecovery[fe.Code]; ok {
			return hint
		}
	}
	if innerErr := errors.Unwrap(err); innerErr != nil {
		return RecoveryHint(innerErr)
	}
	return "Check error details for more information."
}

// NewFutuError creates a FutuError with the given code, function name, and message.
// Category and Recovery are derived automatically from the error code.
func NewFutuError(code ErrorCode, funcName, msg string) *FutuError {
	return &FutuError{
		Code:     code,
		Message:  msg,
		Func:     funcName,
		Category: codeToCategory[code],
		Recovery: codeToRecovery[code],
	}
}

// NewFutuErrorWithWrap creates a FutuError that wraps an inner error. Category
// and Recovery are derived automatically from the error code.
func NewFutuErrorWithWrap(code ErrorCode, funcName, msg string, inner error) *FutuError {
	return &FutuError{
		Code:     code,
		Message:  msg,
		Func:     funcName,
		Err:      inner,
		Category: codeToCategory[code],
		Recovery: codeToRecovery[code],
	}
}

// AsFutuError attempts to cast err to *FutuError. Returns the FutuError and true
// on success, or nil and false otherwise.
func AsFutuError(err error) (*FutuError, bool) {
	fe, ok := err.(*FutuError)
	return fe, ok
}
