package futuapi

import (
	"testing"
)

func BenchmarkNextSerialNo(b *testing.B) {
	c := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.nextSerialNo()
		}
	})
}

func BenchmarkRecordRequest(b *testing.B) {
	c := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.recordRequest(1001, 1, nil)
	}
}

func BenchmarkRecordPush(b *testing.B) {
	c := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.recordPush()
	}
}

func BenchmarkPoolGetPut(b *testing.B) {
	config := DefaultPoolConfig("127.0.0.1:11111")
	pool := NewClientPool(config)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client, err := pool.Get(b.Context(), PoolTypeGeneral)
		if err != nil {
			b.Fatal(err)
		}
		pool.Put(client)
	}
}
