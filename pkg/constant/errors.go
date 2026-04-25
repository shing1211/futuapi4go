package constant

import "fmt"

type ErrorCode int32

const (
	ErrCodeSuccess        ErrorCode = 0
	ErrCodeInvalidParams   ErrorCode = -1
	ErrCodeTimeout       ErrorCode = -100
	ErrCodeDisconnected  ErrorCode = -200
	ErrCodeUnknown       ErrorCode = -400

	// Connection errors
	ErrCodeNetworkError    ErrorCode = -101
	ErrCodeProtocolErr  ErrorCode = -102
	ErrCodeServerBusy    ErrorCode = -103

	// Account errors
	ErrCodeAccNotFound      ErrorCode = -201
	ErrCodeAccDisabled    ErrorCode = -202
	ErrCodeAccLocked      ErrorCode = -203
	ErrCodeAccAuthFail    ErrorCode = -204

	// Trading errors
	ErrCodeInsufficientBalance ErrorCode = -301
	ErrCodeMarketClosed      ErrorCode = -302
	ErrCodeOrderRejected    ErrorCode = -303
	ErrCodePriceOutOfRange ErrorCode = -304
	ErrCodeQtyTooLarge      ErrorCode = -305
	ErrCodeTradingDisabled ErrorCode = -306
	ErrCodeInvalidSecurity ErrorCode = -307
	ErrCodeNoPermission    ErrorCode = -308

	// Subscription errors
	ErrCodeAlreadySubbed  ErrorCode = -401
	ErrCodeNotSubbed      ErrorCode = -402
)

type FutuError struct {
	Code    ErrorCode
	Message string
	Func   string
	Err    error
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

func IsTimeout(err error) bool {
	if fe, ok := err.(*FutuError); ok {
		return fe.Code == ErrCodeTimeout
	}
	return false
}

func IsDisconnected(err error) bool {
	if fe, ok := err.(*FutuError); ok {
		return fe.Code == ErrCodeDisconnected
	}
	return false
}

func IsInvalidParams(err error) bool {
	if fe, ok := err.(*FutuError); ok {
		return fe.Code == ErrCodeInvalidParams
	}
	return false
}

func IsSuccess(err error) bool {
	if fe, ok := err.(*FutuError); ok {
		return fe.Code == ErrCodeSuccess
	}
	return false
}

func getFutuError(err error) (*FutuError, bool) {
	fe, ok := err.(*FutuError)
	return fe, ok
}

func IsNetworkError(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeNetworkError || fe.Code == ErrCodeProtocolErr
	}
	return false
}

func IsServerBusy(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeServerBusy
	}
	return false
}

func IsAccountError(err error) bool {
	if fe, ok := getFutuError(err); ok {
		switch fe.Code {
		case ErrCodeAccNotFound, ErrCodeAccDisabled, ErrCodeAccLocked, ErrCodeAccAuthFail:
			return true
		}
	}
	return false
}

func IsInsufficientBalance(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeInsufficientBalance
	}
	return false
}

func IsMarketClosed(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeMarketClosed
	}
	return false
}

func IsOrderRejected(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeOrderRejected
	}
	return false
}

func IsSubscriptionError(err error) bool {
	if fe, ok := getFutuError(err); ok {
		return fe.Code == ErrCodeAlreadySubbed || fe.Code == ErrCodeNotSubbed
	}
	return false
}

func AsFutuError(err error) (*FutuError, bool) {
	fe, ok := err.(*FutuError)
	return fe, ok
}