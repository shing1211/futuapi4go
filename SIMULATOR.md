# futuapi4go OpenD Simulator

Local mock server for testing SDK without a real Futu OpenD server.

## Project Status

| Module | Status | Description |
|------|------|------|
| Core Server | Complete | TCP listening, protocol parsing (46-byte header), handler registration |
| System APIs | Complete | InitConnect, KeepAlive, GetGlobalState, GetUserInfo |
| Qot Market Data APIs | Complete | 42 Handlers registered |
| Trd Trading APIs | Complete | 13 Handlers registered |
| Push Simulation | Complete | 11 Push Handlers registered |

## Architecture

```
simulator/
├── server.go         # TCP server core (46-byte protocol header, LittleEndian)
├── handlers.go       # System API handlers (4 handlers)
├── handlers_qot.go    # Market data API handlers (42 handlers)
├── handlers_trd.go    # Trading API handlers (13 handlers)
├── handlers_push.go   # Push handlers (11 handlers)
└── simulator_test.go  # Tests
```

## Run Simulator

```bash
cd examples/simulator
go run main.go
```

## Protocol Compatibility

Simulator uses exactly the same protocol format as real Futu OpenD:

| Field | Size | Byte Order | Description |
|------|------|--------|------|
| Magic | 2 | - | "FT" |
| ProtoID | 4 | Little | Protocol ID |
| ProtoFmt | 4 | Little | Protocol format (Protobuf) |
| ProtoVer | 2 | Little | Protocol version |
| SerialNo | 4 | Little | Serial number |
| BodyLen | 4 | Little | Body length |
| BodySHA1 | 20 | - | SHA1 (not used yet) |
| Reserved | 8 | - | Reserved field |

**Total Header Length: 46 bytes**

## Implemented Qot Handlers (2101-2405)

| API | ProtoID | Status |
|-----|---------|------|
| GetBasicQot | 2101 | Full implementation |
| GetKL | 2102 | Full implementation |
| GetOrderBook | 2106 | Full implementation |
| GetTicker | 2107 | Stub |
| GetRT | 2108 | Stub |
| GetSecuritySnapshot | 2110 | Stub |
| GetBroker | 2111 | Stub |
| GetStaticInfo | 2201 | Full implementation |
| GetPlateSet | 2202 | Stub |
| GetPlateSecurity | 2203 | Stub |
| GetOwnerPlate | 2204 | Stub |
| GetReference | 2205 | Stub |
| GetTradeDate | 2206 | Full implementation |
| RequestTradeDate | 2207 | Stub |
| GetMarketState | 2208 | Stub |
| GetSuspend | 2209 | Stub |
| GetCodeChange | 2210 | Stub |
| GetFutureInfo | 2211 | Stub |
| GetIpoList | 2212 | Stub |
| GetHoldingChangeList | 2213 | Stub |
| RequestRehab | 2214 | Stub |
| GetCapitalFlow | 2301 | Stub |
| GetCapitalDistribution | 2302 | Stub |
| StockFilter | 2303 | Stub |
| GetOptionChain | 2304 | Stub |
| GetOptionExpirationDate | 2305 | Stub |
| GetWarrant | 2306 | Stub |
| GetUserSecurity | 2401 | Stub |
| GetUserSecurityGroup | 2402 | Stub |
| ModifyUserSecurity | 2403 | Stub |
| GetPriceReminder | 2404 | Stub |
| SetPriceReminder | 2405 | Stub |
| Subscribe | 3001 | Full implementation |
| GetSubInfo | 3002 | Full implementation |
| RegQotPush | 3003 | Full implementation |

## Usage

### Start Simulator

```go
package main

import (
    "log"

    "gitee.com/shing1211/futuapi4go/simulator"
)

func main() {
    srv := simulator.New("127.0.0.1:11111")
    srv.RegisterDefaultHandlers()  // System handlers
    srv.RegisterQotHandlers()     // Market data handlers
    
    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }
    
    log.Println("Simulator started on 127.0.0.1:11111")
    <-make(chan struct{}) // Block until signal received
}
```

### Use Simulator to Test SDK

```go
package main

import (
    "log"

    futuapi "gitee.com/shing1211/futuapi4go/client"
    "gitee.com/shing1211/futuapi4go/qot"
    "gitee.com/shing1211/futuapi4go/pb/qotcommon"
)

func main() {
    cli := futuapi.New()
    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
    securities := []*qotcommon.Security{
        {Market: &market, Code: func() *string { s := "00700"; return &s }()},
    }

    result, err := qot.GetBasicQot(cli, securities)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Received: %+v", result)
}
```

## Next Steps

1. Implement Qot Market Data API Handler framework
2. Enhance stub handlers with realistic mock data
3. Implement Trd Trading API Handler
4. Add push simulation support
5. Add configurable mock data

## Advanced Simulator Planning

### 1. Price/Order Simulation Engine
- [ ] Real-time price movement simulation (tick-by-tick)
- [ ] Random walk/trend simulation
- [ ] Order book simulation (bid/ask depth)
- [ ] Order matching engine (simulate fills)

### 2. Error/Boundary Test Injection
- [ ] Network latency simulation (fixed/random)
- [ ] Network failure injection (disconnect, timeout)
- [ ] Return error responses (various RetTypes)
- [ ] Boundary value test data

### 3. Scenario Recording/Playback
- [ ] Record real market data
- [ ] Replay historical data
- [ ] Batch test scenarios

---

*Last updated: 2026-04-07*
