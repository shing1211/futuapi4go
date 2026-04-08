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
├── client/           # Core client
│   ├── conn.go       # TCP connection and binary protocol encapsulation
│   ├── client.go     # Main client, connection management
│   └── errors.go     # Error type definitions
├── qot/              # Market Data API
│   └── quote.go      # All Qot API implementations
├── trd/              # Trading API
│   └── trade.go      # All Trd API implementations
├── sys/              # System API
│   └── system.go     # System-level APIs
├── push/             # Push notification handling
│   ├── qot_push.go   # Qot push parsers
│   └── trd_push.go   # Trd push parsers
├── pb/               # Protobuf generated Go code (local module)
├── proto/            # Original Protobuf definition files
└── examples/         # Usage examples
```

### Core Components

#### 1. Connection Layer (client/conn.go)

TCP connection and custom binary protocol encapsulation:

- **Protocol header** (46 bytes): Magic "FT" + ProtoID + SerialNo + BodyLen
- **Heartbeat mechanism**: Auto-send KeepAlive for keep-alive
- **Packet read/write**: WritePacket / ReadPacket

```go
// Connection structure
type Conn struct {
    conn   net.Conn
    protoID int32
    serialNo int32
    mu       sync.Mutex
}

// Write packet
func (c *Conn) WritePacket(protoID int32, serialNo int32, body []byte) error

// Read packet
func (c *Conn) ReadPacket() (*Packet, error)
```

#### 2. Client (client/client.go)

High-level API, encapsulates connection and serialization:

```go
type Client struct {
    conn *Conn
    serialNo int32
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
|------|------|------|
| Magic | 2 bytes | "FT" fixed value |
| ProtoID | 4 bytes | Protocol number (BigEndian) |
| SerialNo | 4 bytes | Serial number (BigEndian) |
| BodyLen | 4 bytes | Body length (BigEndian) |
| Body | Variable | Protobuf serialized data |

### Protocol Encoding Implementation

```go
// WritePacket in conn.go
func (c *Conn) WritePacket(protoID int32, serialNo int32, body []byte) error {
    // 1. Build protocol header
    header := make([]byte, 14)
    binary.BigEndian.PutUint16(header[0:2], 0x4654)  // "FT"
    binary.BigEndian.PutUint32(header[2:6], uint32(protoID))
    binary.BigEndian.PutUint32(header[6:10], uint32(serialNo))
    binary.BigEndian.PutUint32(header[10:14], uint32(len(body)))
    
    // 2. Send header + body
    _, err := c.conn.Write(append(header, body...))
    return err
}
```

---

## API Implementation Pattern

All APIs follow a unified pattern, using `GetBasicQot` as example:

### 1. Define Request/Response Structs

```go
// Request struct - corresponds to Protobuf C2S
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
    // 1. Build Protobuf request
    c2s := &qotgetkl.C2S{
        Security:  req.Security,
        RehabType: &req.RehabType,
        KlType:    &req.KLType,
        ReqNum:    &req.ReqNum,
    }
    pkt := &qotgetkl.Request{C2S: c2s}

    // 2. Serialize
    body, err := proto.Marshal(pkt)
    if err != nil {
        return nil, err
    }

    // 3. Send request
    serialNo := c.NextSerialNo()
    if err := c.Conn().WritePacket(ProtoID_GetKL, serialNo, body); err != nil {
        return nil, err
    }

    // 4. Read response
    pktResp, err := c.Conn().ReadPacket()
    if err != nil {
        return nil, err
    }

    // 5. Deserialize
    var rsp qotgetkl.Response
    if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
        return nil, err
    }

    // 6. Check return code
    if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
        return nil, fmt.Errorf("GetKL failed: retType=%d, retMsg=%s", 
            rsp.GetRetType(), rsp.GetRetMsg())
    }

    // 7. Transform result
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
            // ... other fields
        })
    }

    return result, nil
}
```

### 3. Key Notes

- **Pointer fields**: Protobuf-generated struct fields are mostly pointers, must assign address
- **Naming differences**: Sometimes Protobuf getter method names differ from field names (e.g., `GetOrederCount` instead of `GetOrderCount`)
- **Required vs Optional**: Required fields must be non-nil, optional fields can be nil
- **Error handling**: Always check if RetType is success value

---

## Code Standards

### Protobuf Field Handling

```go
// Correct: use pointer assignment
c2s := &qotgetkl.C2S{
    Security:  req.Security,
    RehabType: &req.RehabType,  // int32 needs address-of
    KlType:    &req.KLType,
    ReqNum:    &req.ReqNum,
}

// Wrong: pass value directly
c2s := &qotgetkl.C2S{
    Security:  req.Security,
    RehabType: req.RehabType,  // compilation error
}
```

### String Fields

```go
// Correct
beginTime := "2024-01-01"
c2s := &xxx.C2S{
    BeginTime: &beginTime,
}

// Wrong
c2s := &xxx.C2S{
    BeginTime: "2024-01-01",  // compilation error
}
```

### Slice Fields

```go
// Correct
typeList := []int32{1, 2, 3}
c2s := &xxx.C2S{
    TypeList: typeList,
}

// Wrong
c2s := &xxx.C2S{
    TypeList: []int32{1, 2, 3},  // compilation error
}
```

### Protobuf Imports

All protobuf packages use local paths:

```go
import (
    "google.golang.org/protobuf/proto"
    futuapi "gitee.com/shing1211/futuapi4go/client"
    "gitee.com/shing1211/futuapi4go/pb/common"
    "gitee.com/shing1211/futuapi4go/pb/qotcommon"
    "gitee.com/shing1211/futuapi4go/pb/qotgetkl"
)
```

---

## Protobuf Definitions

### File Locations

- **Source definitions**: `proto/` directory
- **Generated code**: `pb/` directory (local Go module)

### Adding New Protobuf

1. If using new proto files, first generate Go code with protoc:

```bash
cd proto
protoc --go_out=../pb --go_opt=paths=source_relative \
    -I. \
    Qot_GetKL.proto Qot_Common.proto Common.proto
```

2. Update `pb/go.mod` to ensure correct module name:

```go
module gitee.com/shing1211/futuapi4go/pb
```

### ProtoID Constant Definitions

Each API defines constants in quote.go:

```go
const (
    ProtoID_GetBasicQot = 2101
    ProtoID_GetKL       = 2102
    // ...
)
```

Protocol numbers reference Futu OpenD API documentation.

---

## Testing

### Compilation Tests

```bash
# Compile all packages
go build ./...

# Run tests
go test ./...

# Static analysis
go vet ./...
```

### Integration Tests

Requires Futu OpenD running locally:

```bash
# Run examples
go run examples/main.go
```

### Common Compilation Errors

#### 1. Protobuf import error

```
could not import google.golang.org/protobuf/reflect/protoreflect
```

Solution: Ensure `go.mod` contains correct dependencies:

```go
require (
    google.golang.org/protobuf v1.32.0
)
```

#### 2. LSP errors but compilation passes

Some LSPs (like gopls) may not resolve local `replace` directives. As long as `go build` passes, it's fine.

---

## Debugging

### Enable Connection Debug

```go
cli := futuapi.New()
cli.SetDebug(true)
```

### View Protocol Packets

Add logging in `conn.go` WritePacket/ReadPacket:

```go
func (c *Conn) WritePacket(protoID int32, serialNo int32, body []byte) error {
    log.Printf(">>> WritePacket: protoID=%d, serialNo=%d, len=%d", 
        protoID, serialNo, len(body))
    // ...
}
```

### Common Troubleshooting

#### 1. Connection refused

- Confirm Futu OpenD is started
- Confirm port number is correct (default 11111)
- Confirm firewall allows connection

#### 2. Timeout errors

- Check network connection
- Increase timeout settings

#### 3. Protobuf deserialization failed

- Check if proto version matches OpenD
- Confirm body is complete (check BodyLen)

#### 4. Return error codes

```go
// Common RetType values
RetType_Succeed      = 0   // Success
RetType_CommonFail   = -1   // Common failure
RetType_SystemFail   = -2   // System error
RetType_NoAuth       = -3   // Unauthorized
```

---

## Contribution Process

### 1. Implement New API

1. Confirm proto definition exists in `proto/`
2. Confirm Go code generated in `pb/`
3. Add in `qot/quote.go`:
   - import statements
   - ProtoID constant
   - Request/response structs
   - API function implementation
4. Compile and verify: `go build ./...`
5. Commit code

### 2. Update Documentation

- Update API status in `README.md`
- Add usage example for new API in `USER_GUIDE.md`
- Update `DEVELOPER.md` for major architecture changes

### 3. Commit Standards

```
<module>: <short description>

Detailed explanation (optional)

Fixed issues: Fixes #xxx
```

Example:

```
qot: implement GetWarrant (2306)

Add support for querying warrant data with comprehensive
filter options including maturity, price, premium, etc.
```

---

## Module Maintenance

### pb Module Management

`pb/` is an independent Go module:

```bash
cd pb
go mod tidy
go build ./...
```

Main project uses `replace` directive for local pb:

```go
// go.mod
require (
    gitee.com/shing1211/futuapi4go/pb v0.0.0
)

replace gitee.com/shing1211/futuapi4go/pb => ./pb
```

### Dependency Updates

```bash
# Main project
go get -u gitee.com/shing1211/futuapi4go/pb

# pb module
cd pb && go get -u google.golang.org/protobuf
```
