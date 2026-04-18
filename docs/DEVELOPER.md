# futuapi4go Developer Guide

This guide is for SDK developers, covering project architecture, code standards, and contribution process.

## Table of Contents

1. [Project Architecture](#project-architecture)
2. [Protocol Format](#protocol-format)
3. [Protobuf Definitions](#protobuf-definitions)
4. [Adding a New API](#adding-a-new-api)
5. [Debugging](#debugging)

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

## Protocol Format

Each packet: 46-byte header (Magic "FT" + ProtoID + SerialNo + BodyLen, all BigEndian) + Protobuf body.
See `internal/client/conn.go` for the full implementation. Key ProtoIDs: InitConnect=1001, GetGlobalState=1002, Subscribe=3001, GetBasicQot=3004, GetKL=3006, PlaceOrder=5001.

---

## Protobuf Definitions

- **Source**: `api/proto/` — original `.proto` files
- **Generated**: `pkg/pb/` — `github.com/shing1211/futuapi4go` Go module

To regenerate:

```powershell
./scripts/regen-all-protos.ps1
```

Proto fields are pointers (`*int32`, `*string`) — use address-of (`&val`) in struct literals.
Always check `RetType` for success before accessing `S2C`.

---

## Adding a New API

1. Confirm proto definition in `api/proto/`
2. Run `./scripts/regen-all-protos.ps1`
3. Add wrapper in `pkg/qot/quote.go` or `pkg/trd/trade.go`:
   - ProtoID constant
   - Request/response structs
   - Function using existing patterns (see `pkg/qot/` and `pkg/trd/` for examples)
4. Verify: `go build ./...`
5. Add test in `pkg/qot/` or `pkg/trd/`
6. Update `docs/CHANGELOG.md`

---

## Debugging

### Enable Debug Logging

```go
cli.SetLogger(log.New(os.Stderr, "", 0))
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

### Common RetType Values

- `0` = Success
- `-1` = Common fail
- `-2` = System fail
- `-3` = No auth

### Common Issues

1. **Connection refused** — Futu OpenD not running or wrong port (default: 11111)
2. **Timeout** — check network, increase timeout
3. **Protobuf deserialization failed** — proto version mismatch with OpenD
