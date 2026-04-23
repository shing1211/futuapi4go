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

package benchmark_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetorderbook"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/test/fixtures"
	testutil "github.com/shing1211/futuapi4go/test/util"
	"google.golang.org/protobuf/proto"
)

// ============================================================================
// Benchmark Tests for Critical Paths
// Run with: go test -bench=. -benchmem ./test/benchmark
// ============================================================================

// BenchmarkGetBasicQot_Mock benchmarks GetBasicQot with mock server
func BenchmarkGetBasicQot_Mock(b *testing.B) {
	server := testutil.NewMockServer(&testing.T{})

	server.RegisterHandler(3004, func(req []byte) ([]byte, error) {
		hsiQuote := fixtures.HSIQuote()
		s2c := &qotgetbasicqot.S2C{
			BasicQotList: []*qotcommon.BasicQot{hsiQuote},
		}
		return proto.Marshal(&qotgetbasicqot.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		b.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(&testing.T{}, server)
	defer cleanup()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := qot.GetBasicQot(context.Background(), cli, []*qotcommon.Security{fixtures.HSISecurity()})
		if err != nil {
			b.Fatalf("GetBasicQot failed: %v", err)
		}
	}
}

// BenchmarkGetKL_Mock benchmarks GetKL with mock server
func BenchmarkGetKL_Mock(b *testing.B) {
	server := testutil.NewMockServer(&testing.T{})

	server.RegisterHandler(3006, func(req []byte) ([]byte, error) {
		klList := fixtures.HSIKLineData(100, int32(qotcommon.KLType_KLType_Day))
		s2c := &qotgetkl.S2C{
			Security: fixtures.HSISecurity(),
			Name:     ptrStr(fixtures.HSIName),
			KlList:   klList,
		}
		return proto.Marshal(&qotgetkl.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		b.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(&testing.T{}, server)
	defer cleanup()

	req := &qot.GetKLRequest{
		Security:  fixtures.HSISecurity(),
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    100,
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := qot.GetKL(cli, req)
		if err != nil {
			b.Fatalf("GetKL failed: %v", err)
		}
	}
}

// BenchmarkGetOrderBook_Mock benchmarks GetOrderBook with mock server
func BenchmarkGetOrderBook_Mock(b *testing.B) {
	server := testutil.NewMockServer(&testing.T{})

	server.RegisterHandler(3012, func(req []byte) ([]byte, error) {
		asks, bids := fixtures.HSIOrderBookLevels(10)
		s2c := &qotgetorderbook.S2C{
			Security:         fixtures.HSISecurity(),
			Name:             ptrStr(fixtures.HSIName),
			OrderBookAskList: asks,
			OrderBookBidList: bids,
		}
		return proto.Marshal(&qotgetorderbook.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		b.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(&testing.T{}, server)
	defer cleanup()

	req := &qot.GetOrderBookRequest{
		Security: fixtures.HSISecurity(),
		Num:      10,
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := qot.GetOrderBook(cli, req)
		if err != nil {
			b.Fatalf("GetOrderBook failed: %v", err)
		}
	}
}

// BenchmarkProtobufMarshal_HSIQuote benchmarks protobuf marshaling
func BenchmarkProtobufMarshal_HSIQuote(b *testing.B) {
	hsiQuote := fixtures.HSIQuote()
	resp := &qotgetbasicqot.Response{
		S2C: &qotgetbasicqot.S2C{
			BasicQotList: []*qotcommon.BasicQot{hsiQuote},
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(resp)
		if err != nil {
			b.Fatalf("Marshal failed: %v", err)
		}
	}
}

// BenchmarkProtobufUnmarshal_HSIQuote benchmarks protobuf unmarshaling
func BenchmarkProtobufUnmarshal_HSIQuote(b *testing.B) {
	hsiQuote := fixtures.HSIQuote()
	resp := &qotgetbasicqot.Response{
		S2C: &qotgetbasicqot.S2C{
			BasicQotList: []*qotcommon.BasicQot{hsiQuote},
		},
	}

	data, err := proto.Marshal(resp)
	if err != nil {
		b.Fatalf("Marshal failed: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var unmarshaled qotgetbasicqot.Response
		err := proto.Unmarshal(data, &unmarshaled)
		if err != nil {
			b.Fatalf("Unmarshal failed: %v", err)
		}
	}
}

// BenchmarkMultipleSecurities benchmarks fetching quotes for multiple securities
func BenchmarkMultipleSecurities_Mock(b *testing.B) {
	server := testutil.NewMockServer(&testing.T{})

	server.RegisterHandler(3004, func(req []byte) ([]byte, error) {
		// Return multiple quotes
		quotes := make([]*qotcommon.BasicQot, 10)
		for i := range quotes {
			quote := fixtures.HSIQuote()
			code := fmt.Sprintf("%06d", 700+i)
			quote.Security.Code = &code
			quotes[i] = quote
		}

		s2c := &qotgetbasicqot.S2C{
			BasicQotList: quotes,
		}
		return proto.Marshal(&qotgetbasicqot.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		b.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(&testing.T{}, server)
	defer cleanup()

	// Create 10 securities
	securities := make([]*qotcommon.Security, 10)
	for i := 0; i < 10; i++ {
		code := fmt.Sprintf("%06d", 700+i)
		securities[i] = fixtures.SecurityPtr(fixtures.HSIMarket, code)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := qot.GetBasicQot(context.Background(), cli, securities)
		if err != nil {
			b.Fatalf("GetBasicQot failed: %v", err)
		}
	}
}

// BenchmarkConcurrentRequests benchmarks concurrent API calls
func BenchmarkConcurrentRequests_Mock(b *testing.B) {
	server := testutil.NewMockServer(&testing.T{})

	server.RegisterHandler(3004, func(req []byte) ([]byte, error) {
		hsiQuote := fixtures.HSIQuote()
		s2c := &qotgetbasicqot.S2C{
			BasicQotList: []*qotcommon.BasicQot{hsiQuote},
		}
		return proto.Marshal(&qotgetbasicqot.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		b.Fatal(err)
	}
	defer server.Stop()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		cli, cleanup := testutil.NewTestClient(&testing.T{}, server)
		defer cleanup()

		for pb.Next() {
			_, err := qot.GetBasicQot(context.Background(), cli, []*qotcommon.Security{fixtures.HSISecurity()})
			if err != nil {
				b.Fatalf("GetBasicQot failed: %v", err)
			}
		}
	})
}

// BenchmarkHSIFixtures benchmarks fixture creation
func BenchmarkHSIFixtures_Quote(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = fixtures.HSIQuote()
	}
}

func BenchmarkHSIFixtures_KLine(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = fixtures.HSIKLineData(100, int32(qotcommon.KLType_KLType_Day))
	}
}

func BenchmarkHSIFixtures_OrderBook(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = fixtures.HSIOrderBookLevels(10)
	}
}

// Helper function
func ptrStr(s string) *string { return &s }
