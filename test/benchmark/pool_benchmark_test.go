package benchmark_test

import (
	"testing"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
)

func BenchmarkClientPoolGet_Pooled(b *testing.B) {
	config := futuapi.DefaultPoolConfig("127.0.0.1:11111")
	config.MaxSize = 3
	config.MinIdle = 1
	pool := futuapi.NewClientPool(config)
	defer pool.Close()

	pool.StartHealthChecker()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		client, err := pool.Get(futuapi.PoolTypeMarketData)
		if err != nil {
			b.Fatalf("Pool Get failed: %v", err)
		}
		pool.Put(client)
	}
}

func BenchmarkClientPoolGet_Concurrent(b *testing.B) {
	config := futuapi.DefaultPoolConfig("127.0.0.1:11111")
	config.MaxSize = 5
	config.MinIdle = 1
	pool := futuapi.NewClientPool(config)
	defer pool.Close()

	pool.StartHealthChecker()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client, err := pool.Get(futuapi.PoolTypeMarketData)
			if err != nil {
				b.Fatalf("Pool Get failed: %v", err)
			}
			pool.Put(client)
		}
	})
}

func BenchmarkClientPoolReuse(b *testing.B) {
	config := futuapi.DefaultPoolConfig("127.0.0.1:11111")
	config.MaxSize = 3
	config.MinIdle = 1
	pool := futuapi.NewClientPool(config)
	defer pool.Close()

	client, err := pool.Get(futuapi.PoolTypeMarketData)
	if err != nil {
		b.Fatalf("Initial Get failed: %v", err)
	}
	pool.Put(client)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		c, err := pool.Get(futuapi.PoolTypeMarketData)
		if err != nil {
			b.Fatalf("Pool Get failed: %v", err)
		}
		pool.Put(c)
	}
}
