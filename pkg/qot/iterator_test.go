package qot

import (
	"context"
	"testing"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
)

func TestHistoryKLineIteratorNilClient(t *testing.T) {
	client := futuapi.New()
	ctx := context.Background()

	req := &RequestHistoryKLRequest{
		Security: nil,
	}

	it := NewHistoryKLineIterator(ctx, client, req)

	if it.HasNext() {
		t.Error("HasNext() should return false when client is nil (no connection)")
	}

	_, err := it.Next()
	if err == nil {
		t.Error("Next() should error when not connected")
	}
}