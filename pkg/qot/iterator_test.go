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

	if it.HasNext() != true {
		t.Error("HasNext() should return true when client is nil (not initialized)")
	}

	_, err := it.Next()
	if err == nil {
		t.Error("Next() should error when client is nil (not connected)")
	}
}
