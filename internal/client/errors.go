// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package futuapi

import (
	"errors"
	"fmt"
	"strings"
)

type ErrorCategory string

const (
	CategoryConnection ErrorCategory = "connection"
	CategoryProtocol   ErrorCategory = "protocol"
	CategoryEncryption ErrorCategory = "encryption"
	CategoryAPI        ErrorCategory = "api"
	CategoryTimeout    ErrorCategory = "timeout"
	CategoryPool       ErrorCategory = "pool"
	CategoryUnknown    ErrorCategory = "unknown"
)

type ErrorCode int32

const (
	CodeNotConnected ErrorCode = iota - 1000
	CodeConnIDNotSet
	CodeInvalidResponse
	CodeRequestTimeout
	CodeServerError
	CodeInvalidPacket
	CodeEncryptionFailed
	CodeDecryptionFailed
	CodeChecksumMismatch
	CodeInvalidHeader
	CodeInvalidMagic
	CodePacketTooBig
	CodeInvalidBodyLen
	CodePoolExhausted
	CodePoolClosed
	CodeReconnectFailed
	CodeMarshalFailed
	CodeUnmarshalFailed
	CodeTryAgain
	CodeTLSHandshakeFailed
	CodeRateLimited
	CodeRetryExhausted
	CodeConfigInvalid
)

var codeToCategory = map[ErrorCode]ErrorCategory{
	CodeNotConnected:     CategoryConnection,
	CodeConnIDNotSet:     CategoryConnection,
	CodeRequestTimeout:   CategoryTimeout,
	CodeReconnectFailed:  CategoryConnection,
	CodeInvalidHeader:    CategoryProtocol,
	CodeInvalidMagic:     CategoryProtocol,
	CodePacketTooBig:     CategoryProtocol,
	CodeInvalidBodyLen:   CategoryProtocol,
	CodeInvalidResponse:  CategoryProtocol,
	CodeInvalidPacket:    CategoryProtocol,
	CodeChecksumMismatch: CategoryProtocol,
	CodeServerError:      CategoryAPI,
	CodeEncryptionFailed: CategoryEncryption,
	CodeDecryptionFailed: CategoryEncryption,
	CodePoolExhausted:    CategoryPool,
	CodePoolClosed:       CategoryPool,
	CodeMarshalFailed:    CategoryProtocol,
	CodeUnmarshalFailed:  CategoryProtocol,
	CodeTryAgain:         CategoryPool,
	CodeTLSHandshakeFailed: CategoryConnection,
	CodeRateLimited:     CategoryAPI,
	CodeRetryExhausted:  CategoryAPI,
	CodeConfigInvalid:   CategoryAPI,
}

var codeToRecovery = map[ErrorCode]string{
	CodeNotConnected:     "Ensure client is connected before making API calls. Call Connect() first.",
	CodeConnIDNotSet:     "Connection not fully established. Wait for Connect() to complete successfully.",
	CodeRequestTimeout:   "Increase timeout with WithAPITimeout() option or check network connectivity.",
	CodeReconnectFailed:  "Check if Futu OpenD is running and accessible. Verify network/firewall settings.",
	CodeInvalidHeader:    "Protocol mismatch or corrupted data. Reconnect to server.",
	CodeInvalidMagic:     "Connected to non-Futu server. Verify server address.",
	CodePacketTooBig:     "Server returned unexpectedly large response. Check API parameters.",
	CodeInvalidBodyLen:   "Corrupted packet received. Try reconnecting.",
	CodeInvalidResponse:  "Server response malformed. Reconnect and retry.",
	CodeInvalidPacket:    "Invalid packet received. Connection may be corrupted, try reconnecting.",
	CodeChecksumMismatch: "Data integrity compromised. Reconnect and retry request.",
	CodeServerError:      "Server returned error. Check RetMsg in response for details.",
	CodeEncryptionFailed: "RSA encryption failed. Verify public key format is valid PEM.",
	CodeDecryptionFailed: "AES decryption failed. Connection key may be invalid.",
	CodePoolExhausted:    "Connection pool at capacity. Wait for connections or increase MaxSize.",
	CodePoolClosed:       "Pool closed. Create a new pool with NewClientPool().",
	CodeMarshalFailed:    "Failed to serialize request. Check proto message fields.",
	CodeUnmarshalFailed:  "Failed to deserialize response. Server may have returned unexpected format.",
	CodeTryAgain:         "Pool is busy. Retry the operation.",
	CodeTLSHandshakeFailed: "TLS handshake failed. Verify TLS certificate and server address.",
	CodeRateLimited:     "API rate limit exceeded. Reduce request frequency or increase rate limit.",
	CodeRetryExhausted:  "All retry attempts failed. Check server availability and network.",
	CodeConfigInvalid:   "Invalid configuration. Check client options for correctness.",
}

var (
	ErrNotConnected     = errors.New("not connected")
	ErrConnIDNotSet     = errors.New("connection ID not set")
	ErrInvalidResponse  = errors.New("invalid response")
	ErrRequestTimeout   = errors.New("request timeout")
	ErrServerError      = errors.New("server error")
	ErrInvalidPacket    = errors.New("invalid packet")
	ErrEncryptionFailed = errors.New("encryption failed")
	ErrDecryptionFailed = errors.New("decryption failed")
	ErrChecksumMismatch = errors.New("response checksum mismatch: body integrity compromised")
	ErrInvalidHeader    = errors.New("invalid packet header")
	ErrInvalidMagic     = errors.New("invalid magic bytes")
	ErrPacketTooBig     = errors.New("packet too large")
	ErrInvalidBodyLen   = errors.New("invalid body length")
	ErrPoolExhausted    = errors.New("pool exhausted: all connections in use")
	ErrPoolClosed       = errors.New("pool is closed")
	ErrReconnectFailed  = errors.New("reconnect failed: max retries exceeded")
	ErrMarshalFailed    = errors.New("marshal failed")
	ErrUnmarshalFailed  = errors.New("unmarshal failed")
)

type Error struct {
	Code        ErrorCode
	Category    ErrorCategory
	Message     string
	Recovery    string
	Wrapped     error
	Suggestions []string
}

func (e *Error) Error() string {
	if e.Wrapped != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Wrapped)
	}
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Wrapped
}

func (e *Error) Is(target error) bool {
	if t, ok := target.(*Error); ok {
		return t.Code == e.Code
	}
	if e.Wrapped != nil {
		return errors.Is(e.Wrapped, target)
	}
	return false
}

func (e *Error) CodeString() string {
	return fmt.Sprintf("ERR_%d", e.Code)
}

func (e *Error) CategoryString() string {
	return string(e.Category)
}

func (e *Error) FullMessage() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s] %s\n", e.CodeString(), e.Message))
	if e.Recovery != "" {
		sb.WriteString(fmt.Sprintf("Recovery: %s\n", e.Recovery))
	}
	if len(e.Suggestions) > 0 {
		sb.WriteString("Suggestions:\n")
		for _, s := range e.Suggestions {
			sb.WriteString(fmt.Sprintf("  - %s\n", s))
		}
	}
	if e.Wrapped != nil {
		sb.WriteString(fmt.Sprintf("Caused by: %v\n", e.Wrapped))
	}
	return sb.String()
}

func NewError(code ErrorCode, msg string) *Error {
	return &Error{
		Code:     code,
		Category: codeToCategory[code],
		Message:  msg,
		Recovery: codeToRecovery[code],
	}
}

func NewErrorWithWrap(code ErrorCode, msg string, wrapped error) *Error {
	return &Error{
		Code:     code,
		Category: codeToCategory[code],
		Message:  msg,
		Recovery: codeToRecovery[code],
		Wrapped:  wrapped,
	}
}

func NewErrorWithSuggestions(code ErrorCode, msg string, suggestions []string) *Error {
	return &Error{
		Code:        code,
		Category:    codeToCategory[code],
		Message:     msg,
		Recovery:    codeToRecovery[code],
		Suggestions: suggestions,
	}
}

func NewAPIError(retType int32, retMsg string) *Error {
	code, msg := apiErrorCode(retType, retMsg)
	return &Error{
		Code:        code,
		Category:    CategoryAPI,
		Message:     msg,
		Recovery:    "Check server response message for details. Verify request parameters.",
		Suggestions: apiErrorSuggestions(retType, retMsg),
	}
}

func apiErrorCode(retType int32, retMsg string) (ErrorCode, string) {
	switch retType {
	case 0:
		return ErrorCode(0), "success"
	case -1:
		return CodeServerError, fmt.Sprintf("request failed: %s", retMsg)
	case -100:
		return CodeRequestTimeout, fmt.Sprintf("request timeout: %s", retMsg)
	case -200:
		return CodeNotConnected, fmt.Sprintf("connection lost: %s", retMsg)
	case -400:
		return CodeServerError, fmt.Sprintf("unknown error: %s", retMsg)
	case -500:
		return CodeInvalidPacket, fmt.Sprintf("invalid request: %s", retMsg)
	default:
		return CodeServerError, fmt.Sprintf("server error (code %d): %s", retType, retMsg)
	}
}

func apiErrorSuggestions(retType int32, retMsg string) []string {
	var suggestions []string
	switch retType {
	case -1:
		suggestions = append(suggestions, "Check if request parameters are valid", "Verify market data permissions")
	case -100:
		suggestions = append(suggestions, "Increase API timeout", "Check network latency", "Retry request")
	case -200:
		suggestions = append(suggestions, "Reconnect to server", "Enable auto-reconnect")
	case -500:
		suggestions = append(suggestions, "Validate request fields", "Check API version compatibility")
	}
	return suggestions
}

func CategoryOf(err error) ErrorCategory {
	if e, ok := err.(*Error); ok {
		return e.Category
	}
	if errors.Is(err, ErrNotConnected) || errors.Is(err, ErrConnIDNotSet) {
		return CategoryConnection
	}
	if errors.Is(err, ErrRequestTimeout) {
		return CategoryTimeout
	}
	if errors.Is(err, ErrEncryptionFailed) || errors.Is(err, ErrDecryptionFailed) {
		return CategoryEncryption
	}
	if errors.Is(err, ErrInvalidPacket) || errors.Is(err, ErrInvalidResponse) {
		return CategoryProtocol
	}
	if errors.Is(err, ErrPoolExhausted) || errors.Is(err, ErrPoolClosed) {
		return CategoryPool
	}
	return CategoryUnknown
}

func RecoveryHint(err error) string {
	if e, ok := err.(*Error); ok {
		return e.Recovery
	}
	if errors.Is(err, ErrNotConnected) {
		return codeToRecovery[CodeNotConnected]
	}
	if errors.Is(err, ErrRequestTimeout) {
		return codeToRecovery[CodeRequestTimeout]
	}
	return "Check error details for more information."
}

func IsConnectionError(err error) bool {
	return CategoryOf(err) == CategoryConnection
}

func IsTimeoutError(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == CodeRequestTimeout
	}
	return errors.Is(err, ErrRequestTimeout)
}

func IsProtocolError(err error) bool {
	return CategoryOf(err) == CategoryProtocol
}
