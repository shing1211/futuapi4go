# futuapi4go Developer Guide

This guide is for SDK developers, covering project architecture, code standards, and contribution process.

## Table of Contents

1. [Project Architecture](#project-architecture)
2. [Protocol Layer Implementation](#protocol-layer-implementation)
3. [API Implementation Pattern](#api-implementation-pattern)
4. [Code Standards](#code-standards)
5. [Protobuf Definitions](#protobuf-definitions)
6. [Testing](#testing)
7. [Debugging](#debugging)

---

## Project Architecture

```
futuapi4go/
├── internal/client/     # Core client (Conn, Client, reconnect, KeepAlive)
├── pkg/qot/              # Market Data API (quote.go)
├── pkg/trd/              # Trading API (trade.go)
├── pkg/sys/              # System API (system.go)
├── pkg/push/             # Push notification handling (qot_push.go, trd_push.go)
├── pkg/pb/               # Protobuf generated Go code (local module, flat structure)
├── pkg/pb/common/        # Common proto types
├── pkg/pb/qotcommon/     # Qot common types
├── pkg/pb/qotgetkl/      # Qot_GetKL proto
│   └── ... (other pb packages, all under pkg/pb/)
├── api/proto/            # Original Protobuf definition files
├── cmd/simulator/        # OpenD simulator
├── cmd/examples/         # Usage examples
└── scripts/              # Build scripts
```

### Core Components

#### 1. Connection Layer (`internal/client/conn.go`)

TCP connection and custom binary protocol encapsulation:

- **Protocol header** (46 bytes): Magic "FT" + ProtoID + SerialNo + BodyLen
- **Heartbeat mechanism**: Auto-send KeepAlive for keep-alive
- **Packet read/write**: WritePacket / ReadPacket

#### 2. Client (`internal/client/client.go`)

High-level API, encapsulates connection and serialization:

```go
type Client struct {
    conn     *Conn
    serialNo int32
    // ...
}

// Create new client
func New() *Client

// Connect to OpenD
func (c *Client) Connect(addr string) error

// Close connection
func (c.Close() error

// Get next serial number
func (c *Client) NextSerialNo() int32
```

---

## Protocol Layer Implementation

### Futu OpenD Protocol Format

Each request/response packet contains:

| Field | Length | Description |
|-------|--------|-------------|
| Magic | 2 bytes | "FT" fixed value |
| ProtoID | 4 bytes | Protocol number (BigEndian) |
| SerialNo | 4 bytes | Serial number (BigEndian) |
| BodyLen | 4 bytes | Body length (BigEndian) |
| Body | Variable | Protobuf serialized data |

---

## API Implementation Pattern

All APIs follow a unified pattern, using `GetKL` as example:

### 1. Define Request/Response Structs

```go
// Request struct
type GetKLRequest struct {
    Security  *qotcommon.Security
    RehabType int32
    KLType    int32
    ReqNum    int32
}

// Response struct - custom, easy to use
type GetKLResponse struct {
    Security *qotcommon.Security
    Name     string
    KLList   []*KLine
}
```

### 2. Implement API Function

```go
func GetKL(c *futuapi.Client, req *GetKLRequest) (*GetKLResponse, error) {
    c2s := &qotgetkl.C2S{
        Security:  req.Security,
        RehabType: &req.RehabType,
        KlType:    &req.KLType,
        ReqNum:    &req.ReqNum,
    }
    pkt := &qotgetkl.Request{C2S: c2s}

    body, err := proto.Marshal(pkt)
    if err != nil {
        return nil, err
    }

    serialNo := c.NextSerialNo()
    if err := c.Conn().WritePacket(ProtoID_GetKL, serialNo, body); err != nil {
        return nil, err
    }

    pktResp, err := c.Conn().ReadPacket()
    if err != nil {
        return nil, err
    }

    var rsp qotgetkl.Response
    if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
        return nil, err
    }

    if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
        return nil, fmt.Errorf("GetKL failed: retType=%d, retMsg=%s",
            rsp.GetRetType(), rsp.GetRetMsg())
    }

    s2c := rsp.GetS2C()
    if s2c == nil {
        return nil, fmt.Errorf("GetKL: s2c is nil")
    }

    result := &GetKLResponse{
        Security: s2c.GetSecurity(),
        Name:     s2c.GetName(),
        KLList:   make([]*KLine, 0, len(s2c.GetKlList())),
    }

    for _, kl := range s2c.GetKlList() {
        result.KLList = append(result.KLList, &KLine{
            Time:       kl.GetTime(),
            ClosePrice: kl.GetClosePrice(),
        })
    }

    return result, nil
}
```

### 3. Key Notes

- **Pointer fields**: Protobuf-generated struct fields are pointers, must use address-of
- **Naming differences**: Protobuf getter names match proto field names (field name is `OrderCount`, getter is `GetOrderCount`)
- **Error handling**: Always check RetType for success value

---

## Code Standards

### Protobuf Field Handling

```go
// Correct: pointer assignment
c2s := &qotgetkl.C2S{
    Security:  req.Security,
    RehabType: &req.RehabType,
    KlType:    &req.KLType,
    ReqNum:    &req.ReqNum,
}

// Wrong: value passed directly
c2s := &qotgetkl.C2S{
    RehabType: req.RehabType, // compilation error
}
```

### String Fields

```go
beginTime := "2024-01-01"
c2s := &xxx.C2S{
    BeginTime: &beginTime,
}
```

### ProtoID Constants

Each API package defines ProtoID constants in `pkg/qot/quote.go` (for Qot), `pkg/trd/trade.go` (for Trd), and `pkg/sys/system.go`:

```go
// pkg/qot/quote.go
const (
    ProtoID_GetBasicQot = 3004
    ProtoID_GetKL       = 3006
    // ...
)
```

Key ProtoIDs:
- InitConnect=1001, GetGlobalState=1002, KeepAlive=1004
- Subscribe=3001, GetSubInfo=3002, RegQotPush=3003
- GetBasicQot=3004, GetKL=3006, GetRT=3008, GetTicker=3010, GetOrderBook=3012, GetBroker=3014
- RequestHistoryKL=3103, GetOptionChain=3209, StockFilter=3215

---

## Protobuf Definitions

### File Locations

- **Source definitions**: `api/proto/`
- **Generated code**: `pkg/pb/` (flat structure under local Go module `github.com/shing1211/futuapi4go`)

### Regenerating Protobuf Code

Use the provided script to regenerate all protobuf files:

```powershell
./scripts/regen-all-protos.ps1
```

Or manually:

```bash
cd api/proto
protoc --go_out=../../pkg/pb --go_opt=paths=source_relative \
    -I. Qot_GetKL.proto Qot_Common.proto Common.proto
```

---

## Testing

### Compilation Tests

```bash
go build ./...
go test ./...
go vet ./...
```

### Integration Tests

Requires Futu OpenD running locally:

```bash
go run ./cmd/examples/01_market_data_basic/main.go
```

---

## Debugging

### Enable Connection Debug

```go
cli := futuapi.New()
cli.SetDebug(true)
```

### View Protocol Packets

Add logging in `internal/client/conn.go` WritePacket/ReadPacket:

```go
func (c *Conn) WritePacket(protoID int32, serialNo int32, body []byte) error {
    log.Printf(">>> WritePacket: protoID=%d, serialNo=%d, len=%d",
        protoID, serialNo, len(body))
    // ...
}
```

### Common Troubleshooting

1. **Connection refused**: Confirm Futu OpenD is started, port is correct (default 11111)
2. **Timeout errors**: Check network connection, increase timeout settings
3. **Protobuf deserialization failed**: Check proto version matches OpenD
4. **Return error codes**: Common RetType values:
   - `RetType_Succeed = 0`, `RetType_CommonFail = -1`, `RetType_SystemFail = -2`, `RetType_NoAuth = -3`

---

## Contribution Process

### 1. Implement New API

1. Confirm proto definition exists in `api/proto/`
2. Regenerate Go code: `./scripts/regen-all-protos.ps1`
3. Add to `pkg/qot/quote.go` or `pkg/trd/trade.go`:
   - import statements
   - ProtoID constant
   - Request/response structs
   - API function implementation
4. Compile and verify: `go build ./...`
5. Commit code

### 2. Update Documentation

- Update API status in `README.md`
- Update `DEVELOPER.md` for major architecture changes
- Add to `docs/CHANGELOG.md` and `MIGRATION_GUIDE.md` if needed

### 3. Commit Standards

```
<package>: <short description>

Detailed explanation (optional)
```

Example:

```
qot: implement GetWarrant (3210)

Add support for querying warrant data with comprehensive
filter options including maturity, price, premium, etc.
```

For more detail, see [Conventional Commits](https://www.conventionalcommits.org/).

---

## Module Maintenance

### pb Module

`pkg/pb/` is an independent Go module. The main project uses a `replace` directive:

```go
// go.mod
require (
    github.com/shing1211/futuapi4go v0.6.1
    github.com/shing1211/futuapi4go/pkg/pb v0.0.0
)

replace github.com/shing1211/futuapi4go/pkg/pb => ./pkg/pb
```

### Dependency Updates

```bash
# Main project
go get -u github.com/shing1211/futuapi4go

# pb module
cd pkg/pb && go get -u google.golang.org/protobuf
```
