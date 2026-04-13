package futuapi

import "errors"

var (
	ErrNotConnected       = errors.New("not connected")
	ErrConnIDNotSet       = errors.New("connection ID not set")
	ErrInvalidResponse    = errors.New("invalid response")
	ErrRequestTimeout     = errors.New("request timeout")
	ErrServerError        = errors.New("server error")
	ErrInvalidPacket      = errors.New("invalid packet")
	ErrEncryptionFailed   = errors.New("encryption failed")
	ErrDecryptionFailed   = errors.New("decryption failed")
	ErrChecksumMismatch   = errors.New("response checksum mismatch: body integrity compromised")
)

type Error struct {
	Code    int32
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int32, msg string) *Error {
	return &Error{Code: code, Message: msg}
}
