package chanpkg

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/push"
)

func TestWithBufferSize(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"zero uses default", 0, DefaultChanBufferSize},
		{"negative uses default", -1, DefaultChanBufferSize},
		{"valid size", 50, 50},
		{"max capped", MaxChanBufferSize + 1000, MaxChanBufferSize},
		{"at max", MaxChanBufferSize, MaxChanBufferSize},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WithBufferSize(tt.input)
			if got != tt.expected {
				t.Errorf("WithBufferSize(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestNewQuoteChannel(t *testing.T) {
	ch := NewQuoteChannel(10)
	if cap(ch) != 10 {
		t.Errorf("cap(ch) = %d, want 10", cap(ch))
	}
}

func TestChannelTypes(t *testing.T) {
	klCh := NewKLChannel(5)
	if cap(klCh) != 5 {
		t.Errorf("cap(klCh) = %d, want 5", cap(klCh))
	}

	var _ chan<- *push.UpdateKL = klCh
}

func TestMaxBufferSize(t *testing.T) {
	if MaxChanBufferSize != 10000 {
		t.Errorf("MaxChanBufferSize = %d, want 10000", MaxChanBufferSize)
	}
	if DefaultChanBufferSize != 100 {
		t.Errorf("DefaultChanBufferSize = %d, want 100", DefaultChanBufferSize)
	}
}
