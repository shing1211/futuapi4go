# AGENTS.md — futuapi4go

## Build
```bash
go build ./...
go test ./...
```

## Key Files
- `client/client.go` — main client
- `internal/client/` — connection handling
- `api/proto/` — 74 protobuf definitions

## Current Protobuf Version
- v10.3.6308 — **forward-compatible with FutuOpenD 10.3.6308**

## Entry Points
- `client.New()` — create client
- `cli.ConnectWithRSA()` — connect with trading auth
- `cli.GetGlobalState()` — verify connection

## Gotchas
- Uses Protobuf (not JSON) by default with OpenD
- Connection pool in `internal/client/pool.go`
- `APIConnector` interface enables mocking for tests

## Recent Fixes (v0.6.1)
- Push parse functions (`pkg/push/qot_push.go`) unmarshal `S2C` directly — OpenD sends raw S2C bytes, not Response wrapper
- `logf()` nil logger panic fixed — eager `log.Default()` at package level
- Connection state race fixed — `connected int32` with atomic operations
- `client/client_test.go` has `//go:build skip` — pending redesign
- `TestPoolConnReuse` times out — pre-existing, requires real OpenD