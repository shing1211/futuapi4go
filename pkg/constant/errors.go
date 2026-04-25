package constant

import "fmt"

type ErrorCode int32

const (
	ErrCodeSuccess       ErrorCode = 0
	ErrCodeInvalidParams ErrorCode = -1
	ErrCodeTimeout       ErrorCode = -100
	ErrCodeDisconnected   ErrorCode = -200
	ErrCodeUnknown      ErrorCode = -400
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

func AsFutuError(err error) (*FutuError, bool) {
	fe, ok := err.(*FutuError)
	return fe, ok
}