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
	"testing"
)

func TestErrorConstants(t *testing.T) {
	if ErrNotConnected == nil {
		t.Error("ErrNotConnected should be defined")
	}
	if ErrConnIDNotSet == nil {
		t.Error("ErrConnIDNotSet should be defined")
	}
	if ErrInvalidResponse == nil {
		t.Error("ErrInvalidResponse should be defined")
	}
	if ErrRequestTimeout == nil {
		t.Error("ErrRequestTimeout should be defined")
	}
	if ErrServerError == nil {
		t.Error("ErrServerError should be defined")
	}
	if ErrInvalidPacket == nil {
		t.Error("ErrInvalidPacket should be defined")
	}
	if ErrEncryptionFailed == nil {
		t.Error("ErrEncryptionFailed should be defined")
	}
	if ErrDecryptionFailed == nil {
		t.Error("ErrDecryptionFailed should be defined")
	}
}

func TestErrorMessages(t *testing.T) {
	tests := []struct {
		err      error
		expected string
	}{
		{ErrNotConnected, "not connected"},
		{ErrConnIDNotSet, "connection ID not set"},
		{ErrInvalidResponse, "invalid response"},
		{ErrRequestTimeout, "request timeout"},
		{ErrServerError, "server error"},
		{ErrInvalidPacket, "invalid packet"},
		{ErrEncryptionFailed, "encryption failed"},
		{ErrDecryptionFailed, "decryption failed"},
	}

	for _, tt := range tests {
		if tt.err.Error() != tt.expected {
			t.Errorf("%v.Error() = %q, want %q", tt.err, tt.err.Error(), tt.expected)
		}
	}
}

func TestNewError(t *testing.T) {
	err := NewError(CodeNotConnected, "test error")
	if err == nil {
		t.Fatal("NewError returned nil")
	}
	if err.Code != CodeNotConnected {
		t.Errorf("expected Code %v, got %v", CodeNotConnected, err.Code)
	}
	if err.Message != "test error" {
		t.Errorf("expected Message 'test error', got %q", err.Message)
	}
	if err.Category != CategoryConnection {
		t.Errorf("expected Category %v, got %v", CategoryConnection, err.Category)
	}
	if err.Recovery == "" {
		t.Error("Recovery should not be empty")
	}
}

func TestErrorImplementsErrorInterface(t *testing.T) {
	err := NewError(CodeNotConnected, "test")
	var e error = err
	if e.Error() != "test" {
		t.Errorf("Error interface not implemented correctly")
	}
}

func TestErrorsAreComparable(t *testing.T) {
	if !errors.Is(ErrNotConnected, ErrNotConnected) {
		t.Error("ErrNotConnected should be comparable with itself")
	}
}

func TestNewErrorWithWrap(t *testing.T) {
	original := errors.New("original error")
	err := NewErrorWithWrap(CodeServerError, "wrapped error", original)
	if err.Wrapped == nil {
		t.Error("Wrapped should not be nil")
	}
	if !errors.Is(err, original) {
		t.Error("errors.Is should find wrapped error")
	}
}

func TestErrorFullMessage(t *testing.T) {
	err := NewErrorWithSuggestions(CodeRequestTimeout, "request timeout", []string{"increase timeout", "check network"})
	msg := err.FullMessage()
	if msg == "" {
		t.Error("FullMessage should not be empty")
	}
}

func TestCategoryOf(t *testing.T) {
	tests := []struct {
		err      error
		expected ErrorCategory
	}{
		{ErrNotConnected, CategoryConnection},
		{ErrRequestTimeout, CategoryTimeout},
		{ErrEncryptionFailed, CategoryEncryption},
		{ErrInvalidPacket, CategoryProtocol},
		{ErrPoolExhausted, CategoryPool},
		{NewError(CodeNotConnected, "test"), CategoryConnection},
		{NewError(CodeServerError, "test"), CategoryAPI},
	}

	for _, tt := range tests {
		if got := CategoryOf(tt.err); got != tt.expected {
			t.Errorf("CategoryOf(%v) = %v, want %v", tt.err, got, tt.expected)
		}
	}
}

func TestRecoveryHint(t *testing.T) {
	hint := RecoveryHint(ErrNotConnected)
	if hint == "" {
		t.Error("RecoveryHint should not be empty")
	}

	err := NewError(CodeRequestTimeout, "timeout")
	hint = RecoveryHint(err)
	if hint == "" {
		t.Error("RecoveryHint should not be empty for Error type")
	}
}

func TestIsConnectionError(t *testing.T) {
	if !IsConnectionError(ErrNotConnected) {
		t.Error("ErrNotConnected should be a connection error")
	}
	if !IsConnectionError(NewError(CodeNotConnected, "test")) {
		t.Error("NewError(CodeNotConnected) should be a connection error")
	}
	if IsConnectionError(ErrRequestTimeout) {
		t.Error("ErrRequestTimeout should not be a connection error")
	}
}

func TestIsTimeoutError(t *testing.T) {
	if !IsTimeoutError(ErrRequestTimeout) {
		t.Error("ErrRequestTimeout should be a timeout error")
	}
	if !IsTimeoutError(NewError(CodeRequestTimeout, "test")) {
		t.Error("NewError(CodeRequestTimeout) should be a timeout error")
	}
	if IsTimeoutError(ErrNotConnected) {
		t.Error("ErrNotConnected should not be a timeout error")
	}
}

func TestIsProtocolError(t *testing.T) {
	if !IsProtocolError(ErrInvalidPacket) {
		t.Error("ErrInvalidPacket should be a protocol error")
	}
	if !IsProtocolError(NewError(CodeInvalidMagic, "test")) {
		t.Error("NewError(CodeInvalidMagic) should be a protocol error")
	}
	if IsProtocolError(ErrNotConnected) {
		t.Error("ErrNotConnected should not be a protocol error")
	}
}

func TestNewAPIError(t *testing.T) {
	err := NewAPIError(-1, "insufficient permission")
	if err == nil {
		t.Fatal("NewAPIError returned nil")
	}
	if err.Category != CategoryAPI {
		t.Errorf("expected Category API, got %v", err.Category)
	}
	if len(err.Suggestions) == 0 {
		t.Error("Suggestions should not be empty for retType=-1")
	}
}

func TestErrorIs(t *testing.T) {
	err := NewError(CodeNotConnected, "not connected")
	var target *Error
	if !errors.As(err, &target) {
		t.Error("Error.As should work for Error type")
	}
}

func TestErrorUnwrap(t *testing.T) {
	original := errors.New("original")
	err := NewErrorWithWrap(CodeServerError, "wrapped", original)
	if err.Unwrap() != original {
		t.Error("Unwrap should return wrapped error")
	}
}
