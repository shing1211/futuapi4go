package cache

import (
	"sync"
	"testing"
	"time"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

func TestKLCacheSetGet(t *testing.T) {
	c := NewKLCache()
	sec := &qotcommon.Security{
		Market: protoInt32(1),
		Code:   protoString("US.AAPL"),
	}
	klType := int32(1)
	rehabType := int32(0)
	klines := []*qot.KLine{
		{Time: "2024-01-01", OpenPrice: 150.0},
	}

	c.Set(sec, klType, rehabType, klines)
	got, ok := c.Get(sec, klType, rehabType)
	if !ok {
		t.Fatal("expected cache hit")
	}
	if len(got) != 1 {
		t.Fatalf("expected 1 kline, got %d", len(got))
	}
	if got[0].Time != "2024-01-01" {
		t.Errorf("expected time 2024-01-01, got %s", got[0].Time)
	}
}

func TestKLCacheMiss(t *testing.T) {
	c := NewKLCache()
	sec := &qotcommon.Security{
		Market: protoInt32(1),
		Code:   protoString("US.AAPL"),
	}

	_, ok := c.Get(sec, 1, 0)
	if ok {
		t.Error("expected cache miss for non-existent key")
	}
}

func TestKLCacheTTL(t *testing.T) {
	c := NewKLCache(WithTTL(50 * time.Millisecond))
	sec := &qotcommon.Security{
		Market: protoInt32(1),
		Code:   protoString("US.AAPL"),
	}
	klines := []*qot.KLine{{Time: "2024-01-01"}}

	c.Set(sec, 1, 0, klines)

	_, ok := c.Get(sec, 1, 0)
	if !ok {
		t.Fatal("expected cache hit before TTL expiry")
	}

	time.Sleep(100 * time.Millisecond)

	_, ok = c.Get(sec, 1, 0)
	if ok {
		t.Error("expected cache miss after TTL expiry")
	}
}

func TestKLCacheLRUEviction(t *testing.T) {
	c := NewKLCache(WithMaxEntries(2))

	sec1 := &qotcommon.Security{Market: protoInt32(1), Code: protoString("A")}
	sec2 := &qotcommon.Security{Market: protoInt32(1), Code: protoString("B")}
	sec3 := &qotcommon.Security{Market: protoInt32(1), Code: protoString("C")}

	c.Set(sec1, 1, 0, []*qot.KLine{{Time: "v1"}})
	c.Set(sec2, 1, 0, []*qot.KLine{{Time: "v2"}})
	c.Set(sec3, 1, 0, []*qot.KLine{{Time: "v3"}})

	if c.Size() != 2 {
		t.Errorf("expected size 2 after eviction, got %d", c.Size())
	}

	_, ok := c.Get(sec1, 1, 0)
	if ok {
		t.Error("expected sec1 to be evicted (LRU)")
	}

	_, ok = c.Get(sec2, 1, 0)
	if !ok {
		t.Error("expected sec2 to still be in cache")
	}

	_, ok = c.Get(sec3, 1, 0)
	if !ok {
		t.Error("expected sec3 to still be in cache")
	}
}

func TestKLCacheClear(t *testing.T) {
	c := NewKLCache()
	sec := &qotcommon.Security{Market: protoInt32(1), Code: protoString("US.AAPL")}
	c.Set(sec, 1, 0, []*qot.KLine{{Time: "v1"}})

	if c.Size() != 1 {
		t.Fatalf("expected size 1 after set, got %d", c.Size())
	}

	c.Clear()

	if c.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", c.Size())
	}

	_, ok := c.Get(sec, 1, 0)
	if ok {
		t.Error("expected cache miss after clear")
	}
}

func TestKLCacheConcurrent(t *testing.T) {
	c := NewKLCache(WithMaxEntries(100))
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			code := string(rune('A' + n))
			sec := &qotcommon.Security{Market: protoInt32(1), Code: &code}
			klines := []*qot.KLine{{Time: "v"}}

			for j := 0; j < 20; j++ {
				c.Set(sec, int32(j), 0, klines)
			}

			for j := 0; j < 20; j++ {
				_, _ = c.Get(sec, int32(j), 0)
			}
		}(i)
	}
	wg.Wait()

	if c.Size() == 0 {
		t.Error("expected non-zero cache size after concurrent operations")
	}
}

func protoInt32(v int32) *int32    { return &v }
func protoString(v string) *string { return &v }
