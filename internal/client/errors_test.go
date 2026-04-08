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
	err := NewError(-1, "test error")
	if err == nil {
		t.Fatal("NewError returned nil")
	}
	if err.Code != -1 {
		t.Errorf("expected Code -1, got %d", err.Code)
	}
	if err.Message != "test error" {
		t.Errorf("expected Message 'test error', got %q", err.Message)
	}
}

func TestErrorImplementsErrorInterface(t *testing.T) {
	err := NewError(-1, "test")
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
