# futuapi4go Operational Guide

## Architecture

```
Application
  └── client/Client         (Public wrapper API)
       └── pkg/*            (qot, trd, sys — business logic)
            └── internal/client/Client   (Connection management)
                 └── internal/client/Conn  (TCP I/O, packet framing)
                      └── Futu OpenD (TCP socket)
```

**Key constraint:** All communication is via Protocol Buffers over TCP. No JSON by default.

## Connection Lifecycle

1. `client.New()` — creates a client with options
2. `cli.Connect(addr)` — TCP dial → InitConnect handshake → AES key exchange
3. `cli.Close()` — sends close signal, drains goroutines, closes socket

During connect, OpenD returns: `connID`, `loginUserID`, `aesKey`, `serverVer`, `keepAliveInterval`. These are stored and accessible via:
- `cli.GetConnID()` → `uint64`
- `cli.GetLoginUserID()` → `uint64` (Futu/NiuNiu user ID)
- `cli.IsEncrypt()` → `bool` (was RSA key provided?)
- `cli.GetServerVer()` → `int32`
- `cli.CanSendProto(protoID)` → `bool` (connection state check)

## Build & Verify

```bash
go build ./...   # Must always pass
go test ./...    # Run full test suite
go vet ./...     # Lint — must pass before commit
```

## Code Review Focus

| Area | What to Watch |
|------|--------------|
| Concurrency | Mutex usage in `internal/client/conn.go` and `client.go` |
| Pool safety | `Pool.Put()` / `Pool.Get()` thread safety |
| Context | All public APIs should accept `context.Context` |
| Errors | Never swallow errors with `_` |
| Proto | Always check `RetType` before accessing `S2C` |

## Key Entry Points

| File | Purpose |
|------|---------|
| `client/client.go` | Public API wrapper |
| `internal/client/client.go` | Connection, serial numbers, reconnect |
| `internal/client/conn.go` | Raw TCP packet I/O |
| `internal/client/pool.go` | Connection pool |
| `pkg/qot/quote.go` | All market data APIs |
| `pkg/trd/trade.go` | All trading APIs |
| `pkg/sys/system.go` | System APIs |
| `pkg/push/qot_push.go` | Push notification parsers |

## Adding a New API

1. Confirm the proto in `api/proto/`
2. Run `./scripts/regen-all-protos.ps1`
3. Add the wrapper function in `pkg/qot/` or `pkg/trd/`
4. Add a public helper in `client/client.go` if it simplifies usage
5. Add a test
6. Update `docs/CHANGELOG.md` under `[Unreleased]`
7. Verify: `go build ./... && go vet ./...`

## Official Documentation

- Proto Reference: https://openapi.futunn.com/mds/Futu-API-Doc-zh-Proto.md
- Go module: `github.com/shing1211/futuapi4go` (v0.2.0)
