# Testing Guide for futuapi4go

## Overview

This document provides a comprehensive guide to testing the futuapi4go SDK with HSI (Hang Seng Index) as the primary test symbol.

## Test Structure

```
test/
├── fixtures/
│   └── hsi_fixtures.go          # HSI test data and helpers
├── util/
│   └── mock_server.go           # Mock OpenD server for unit tests
├── qot_api/
│   └── qot_test.go              # Qot API unit tests (12+ tests)
├── trd_api/
│   └── trd_test.go              # Trading API unit tests (10+ tests)
├── integration/
│   ├── integration_test.go      # Original integration tests
│   └── integration_hsi_test.go  # HSI-focused integration tests (15+ tests)
└── benchmark/
    └── benchmark_test.go        # Performance benchmarks (10 tests)
```

## Quick Start

### Run All Tests

```bash
# Unit tests (mock server, no OpenD required)
go test -v ./test/qot_api
go test -v ./test/trd_api

# Integration tests (requires running OpenD)
FUTU_INTEGRATION_TESTS=1 go test -v ./test/integration

# Benchmarks
go test -bench=. -benchmem ./test/benchmark
```

### Run Specific HSI Tests

```bash
# HSI quote tests
go test -v ./test/qot_api -run HSI

# HSI trading tests
go test -v ./test/trd_api -run HSI

# HSI integration tests
FUTU_INTEGRATION_TESTS=1 go test -v ./test/integration -run HSI
```

## Test Coverage

### Qot API Tests (12 tests)

| Test | Description | Data |
|------|-------------|------|
| `TestGetBasicQot_HSI` | Real-time quote for HSI | Price: 18523.45, Vol: 2.3B |
| `TestGetKL_HSI_Day` | Daily K-line data | 10 days realistic OHLCV |
| `TestGetKL_HSI_Min1` | 1-minute K-line | 5 bars realistic data |
| `TestGetOrderBook_HSI` | Order book (10 levels) | Realistic spread: 5.0 |
| `TestGetTicker_HSI` | Tick-by-tick trades | 20 ticks with timestamps |
| `TestGetRT_HSI` | Intraday time-share | 240 minutes (full day) |
| `TestGetBroker_HSI` | Broker queue | 10 brokers each side |
| `TestGetStaticInfo_HSI` | Security static info | Index metadata |
| `TestGetTradeDate_HK` | HK trading dates | 5 consecutive dates |
| `TestSubscribe_HSI` | Subscribe to pushes | Basic + KL subtypes |
| `TestGetCapitalFlow_HSI` | Capital flow data | Realistic flow values |
| `TestGetCapitalDistribution_HSI` | Capital distribution | Big/Mid/Small breakdown |

### Trading API Tests (10 tests)

| Test | Description | Data |
|------|-------------|------|
| `TestGetAccList` | Account list | Simulated account |
| `TestUnlockTrade` | Unlock trading | MD5 password hash |
| `TestGetFunds_HSI` | Account funds | Total: 500K, Power: 425K |
| `TestGetPositionList_HSI` | HSI futures position | 2 contracts, P/L: 86.90 |
| `TestPlaceOrder_HSI` | Place buy order | HSI futures @ 18520 |
| `TestPlaceOrder_HSI_Sell` | Place sell order | HSI futures @ 18530 |
| `TestGetOrderList` | Order list | 2 sample orders |
| `TestModifyOrder_HSI` | Modify order price | New price: 18525 |
| `TestModifyOrder_Cancel` | Cancel order | Cancel operation |
| `TestGetOrderFillList_HSI` | Order fills | 2 fill records |
| `TestTradingWorkflow_Complete` | End-to-end workflow | 5-step workflow |

### Integration Tests (15 tests)

| Test | Description | Requires OpenD |
|------|-------------|----------------|
| `TestIntegration_HSI_ConnectAndGlobalState` | Connection + state | Yes |
| `TestIntegration_HSI_BasicQot` | Real-time quote | Yes |
| `TestIntegration_HSI_KLine` | K-line data | Yes |
| `TestIntegration_HSI_OrderBook` | Order book | Yes |
| `TestIntegration_HSI_Ticker` | Tick data | Yes |
| `TestIntegration_HSI_RT` | Intraday RT | Yes |
| `TestIntegration_HSI_Subscribe_Push` | Push notifications | Yes |
| `TestIntegration_HSI_StationaryInfo` | Static info | Yes |
| `TestIntegration_HSI_TradeDate` | Trade dates | Yes |
| `TestIntegration_HSI_CapitalFlow` | Capital flow | Yes |
| `TestIntegration_Trading_Workflow` | Full trading | Yes |
| `TestIntegration_ContextCancellation` | Context handling | Yes |
| `TestIntegration_HSI_ComprehensiveMarketData` | All Qot APIs | Yes |

### Benchmark Tests (10 tests)

| Benchmark | Description | Metrics |
|-----------|-------------|---------|
| `BenchmarkGetBasicQot_Mock` | Quote API perf | Latency, allocs |
| `BenchmarkGetKL_Mock` | K-line API perf | 100 bars |
| `BenchmarkGetOrderBook_Mock` | Order book perf | 10 levels |
| `BenchmarkProtobufMarshal_HSIQuote` | Marshal perf | Serialization |
| `BenchmarkProtobufUnmarshal_HSIQuote` | Unmarshal perf | Deserialization |
| `BenchmarkMultipleSecurities_Mock` | Multi-security | 10 securities |
| `BenchmarkConcurrentRequests_Mock` | Concurrent API | Parallel goroutines |
| `BenchmarkHSIFixtures_Quote` | Fixture creation | Quote creation |
| `BenchmarkHSIFixtures_KLine` | K-line fixture | 100 bars creation |
| `BenchmarkHSIFixtures_OrderBook` | Order book fixture | 10 levels |

## HSI Test Data

### Hang Seng Index (800100.HK)

**Realistic Market Data:**
- Current Price: 18,523.45
- Open: 18,480.00
- High: 18,590.12
- Low: 18,420.50
- Last Close: 18,498.23
- Volume: 2,345,678,900
- Turnover: 98,765,432,100.50
- Turnover Rate: 2.34%
- Amplitude: 0.92%

**Order Book:**
- Spread: 5.0 points
- 10 levels bid/ask
- Realistic volume progression

**K-Line Data:**
- Daily bars with realistic OHLCV
- 1-minute, 5-minute, 15-minute intervals
- Proper price relationships

### HSI Futures (HSImain)

**Trading Data:**
- Contract: HSI Futures
- Lot Size: 1 contract
- Multiplier: HKD 50 per point
- Current Position: 2 contracts
- Cost Price: 18,480.00
- Market Price: 18,523.45
- P/L: HKD 86.90

**Account:**
- Total Assets: HKD 500,000
- Cash: HKD 250,000
- Buying Power: HKD 425,906
- Margin Used: HKD 37,047

## Mock Server

The mock server (`testutil.MockServer`) provides:

- **Full TCP protocol** simulation (44-byte header + protobuf body)
- **Handler registry** for all ProtoIDs
- **Request logging** for assertions
- **Realistic responses** for all APIs
- **Concurrent safe** for parallel tests

### Usage Example

```go
func TestMyAPI(t *testing.T) {
    server := testutil.NewMockServer(t)
    
    // Register custom handler
    server.RegisterHandler(3004, func(req []byte) ([]byte, error) {
        // Build response
        return proto.Marshal(resp)
    })
    
    server.Start()
    defer server.Stop()
    
    cli, cleanup := testutil.NewTestClient(t, server)
    defer cleanup()
    
    // Test your API
    result, err := qot.GetBasicQot(cli, securities)
    
    // Assert
    server.AssertProtoID(t, 3004)
}
```

## Running with Real OpenD

### Prerequisites

1. **Futu OpenD running** on `127.0.0.1:11111`
2. **Logged in** with valid credentials
3. **Market hours** for live data (9:30-16:00 HKT)

### Environment Variables

```bash
# Enable integration tests
export FUTU_INTEGRATION_TESTS=1

# Custom OpenD address (default: 127.0.0.1:11111)
export FUTU_OPEND_ADDR=127.0.0.1:11111
```

### Run Integration Tests

```bash
# All integration tests
FUTU_INTEGRATION_TESTS=1 go test -v ./test/integration

# Specific test
FUTU_INTEGRATION_TESTS=1 go test -v ./test/integration -run HSI_BasicQot

# With timeout
FUTU_INTEGRATION_TESTS=1 go test -v -timeout 5m ./test/integration
```

## Test Best Practices

### Writing New Tests

1. **Use fixtures** from `test/fixtures/hsi_fixtures.go`
2. **Mock server** for unit tests (no OpenD dependency)
3. **Integration tests** for end-to-end validation
4. **Benchmarks** for performance-critical paths

### Test Data

- Use `testutil.HSIQuote()` for realistic HSI quotes
- Use `testutil.HSIKLineData(n, type)` for K-lines
- Use `testutil.HSIOrderBookLevels(n)` for order books
- Use `testutil.HSIFunds()` for account funds

### Assertions

```go
// ProtoID assertion
server.AssertProtoID(t, 3004)

// Request count
server.AssertRequestCount(t, 5)

// Custom assertions
if quote.CurPrice <= 0 {
    t.Error("Price should be positive")
}
```

## Troubleshooting

### Integration Tests Skipped

```
=== SKIP: TestIntegration_HSI_BasicQot
    integration_hsi_test.go:14: Skipping integration test
```

**Solution**: Set `FUTU_INTEGRATION_TESTS=1`

### Connection Refused

```
Connect failed: dial tcp 127.0.0.1:11111: connect: connection refused
```

**Solution**: Start Futu OpenD or use mock server

### Market Closed

```
No push received within timeout (market may be closed)
```

**Solution**: This is expected outside market hours. Tests still validate subscription logic.

## Continuous Integration

### GitHub Actions Example

```yaml
name: Test
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: 1.21
      
      - name: Unit Tests
        run: go test -v ./test/qot_api ./test/trd_api
      
      - name: Benchmarks
        run: go test -bench=. ./test/benchmark
```

## Performance Benchmarks

Run on your machine:

```bash
go test -bench=. -benchmem -cpu=4 ./test/benchmark
```

**Expected Results (mock server):**
- GetBasicQot: ~0.5ms latency, 5 allocations
- GetKL (100 bars): ~1.2ms latency, 150 allocations
- OrderBook (10 levels): ~0.8ms latency, 25 allocations
- Protobuf marshal: ~50μs, 2 allocations
- Protobuf unmarshal: ~30μs, 1 allocation

## Future Enhancements

- [ ] WebSocket transport tests
- [ ] Push notification stress tests
- [ ] Network failure simulation
- [ ] Race condition detection
- [ ] Coverage reporting integration
- [ ] Property-based testing
- [ ] Fuzz testing for protobuf parsing

## Support

For issues or questions:
- Check existing tests for examples
- Review `test/fixtures/hsi_fixtures.go` for test data
- See `test/util/mock_server.go` for mock server API
- Open an issue on GitHub
