# futuapi4go OpenD Simulator

Local mock server for testing SDK without a real Futu OpenD server.

## Project Status

| Module | Status | Description |
|--------|--------|-------------|
| Core Server | Complete | TCP listening, protocol parsing (46-byte header), handler registration |
| System APIs | Complete | InitConnect, KeepAlive, GetGlobalState, GetUserInfo |
| Qot Market Data APIs | Complete | 42 Handlers registered |
| Trd Trading APIs | Complete | 13 Handlers registered |
| Push Simulation | Complete | 11 Push Handlers registered |

## Architecture

```
cmd/simulator/
├── server.go          # TCP server core (46-byte protocol header, LittleEndian)
├── handlers.go        # System API handlers (4 handlers)
├── handlers_qot.go    # Market data API handlers (42 handlers)
├── handlers_trd.go    # Trading API handlers (13 handlers)
└── handlers_push.go  # Push handlers (11 handlers)
```

## Run Simulator

```bash
go run ./cmd/simulator/main.go
```

## Protocol Compatibility

Simulator uses exactly the same protocol format as real Futu OpenD:

| Field | Size | Byte Order | Description |
|-------|------|------------|-------------|
| Magic | 2 | - | "FT" |
| ProtoID | 4 | Little | Protocol ID |
| ProtoFmt | 4 | Little | Protocol format (Protobuf) |
| ProtoVer | 2 | Little | Protocol version |
| SerialNo | 4 | Little | Serial number |
| BodyLen | 4 | Little | Body length |
| BodySHA1 | 20 | - | SHA1 (not used yet) |
| Reserved | 8 | - | Reserved field |

**Total Header Length: 46 bytes**

## Implemented Qot Handlers (3001-3224)

| API | ProtoID | Status |
|-----|---------|--------|
| GetBasicQot | 3004 | Full implementation |
| GetKL | 3006 | Full implementation |
| GetOrderBook | 3012 | Full implementation |
| GetTicker | 3010 | Stub |
| GetRT | 3008 | Stub |
| GetSecuritySnapshot | 3203 | Stub |
| GetBroker | 3014 | Stub |
| GetStaticInfo | 3202 | Full implementation |
| GetPlateSet | 3204 | Stub |
| GetPlateSecurity | 3205 | Stub |
| GetOwnerPlate | 3207 | Stub |
| GetReference | 3206 | Stub |
| GetTradeDate | 3201 | Full implementation |
| RequestTradeDate | 3219 | Stub |
| GetMarketState | 3223 | Stub |
| GetSuspend | 3220 | Stub |
| GetCodeChange | 3216 | Stub |
| GetFutureInfo | 3218 | Stub |
| GetIpoList | 3217 | Stub |
| GetHoldingChangeList | 3230 | Stub |
| RequestRehab | 3200 | Stub |
| GetCapitalFlow | 3211 | Stub |
| GetCapitalDistribution | 3212 | Stub |
| StockFilter | 3215 | Stub |
| GetOptionChain | 3209 | Stub |
| GetOptionExpirationDate | 3224 | Stub |
| GetWarrant | 3210 | Stub |
| GetUserSecurity | 3213 | Stub |
| GetUserSecurityGroup | 3222 | Stub |
| ModifyUserSecurity | 3214 | Stub |
| GetPriceReminder | 3221 | Stub |
| SetPriceReminder | 3220 | Stub |
| Subscribe | 3001 | Full implementation |
| GetSubInfo | 3002 | Full implementation |
| RegQotPush | 3003 | Full implementation |

## Usage

### Start Simulator

```go
package main

import (
    "log"
    "github.com/shing1211/futuapi4go/cmd/simulator"
)

func main() {
    srv := simulator.New("127.0.0.1:11111")
    srv.RegisterDefaultHandlers()
    srv.RegisterQotHandlers()

    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }

    log.Println("Simulator started on 127.0.0.1:11111")
    <-make(chan struct{})
}
```

### Use Simulator to Test SDK

```go
package main

import (
    "log"

    futuapi "github.com/shing1211/futuapi4go/internal/client"
    "github.com/shing1211/futuapi4go/pkg/qot"
    "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

func main() {
    cli := futuapi.New()
    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
    code := "00700"
    securities := []*qotcommon.Security{
        {Market: &market, Code: &code},
    }

    result, err := qot.GetBasicQot(cli, securities)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Received: %+v", result)
}
```

## Next Steps

1. Enhance stub handlers with realistic mock data
2. Add configurable mock data
3. Implement scenario recording/playback

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

*Last updated: 2026-04-10*
