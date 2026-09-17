# Testing Guide

## Layout

| Location | What it tests | Needs OpenD |
|----------|---------------|-------------|
| `*_test.go` beside source | Unit tests (parsing, validation, constants) | No |
| `test/qot_api/` | Quote-API request/response wrapping via the mock | No |
| `test/trd_api/` | Trading-API request/response wrapping via the mock | No |
| `test/util/` | The in-process `MockServer` used by the API tests | No |
| `test/fixtures/` | HSI fixture data | No |
| `test/integration/` | End-to-end against a live OpenD | **Yes** |
| `test/benchmark/` | Benchmarks (connection pool, hot paths) | No |
| `test/examples/` | Example snippets | No |

## Running

```bash
make test          # go test -race ./...
make test-cover    # go test -cover ./...
make bench         # go test -bench=. -benchmem -count=3 ./...
go test -race -run TestName ./pkg/trd/...
```

Always use `-race` — the SDK is concurrent.

## Integration tests

Integration tests are gated by an **environment variable**, not a build tag:

```bash
FUTU_INTEGRATION_TESTS=1 FUTU_OPEND_ADDR=127.0.0.1:11111 \
  go test -race ./test/integration/...
```

Without `FUTU_INTEGRATION_TESTS=1` they call `t.Skip`. `make test-integration`
sets the variable for you.

## Mock server

`test/util.MockServer` implements the OpenD wire protocol in-process (handshake,
AES key exchange, 44-byte framing, SHA-1 body check). Register a handler per
ProtoID and assert which ProtoIDs were exercised:

```go
server := testutil.NewMockServer(t)
server.RegisterHandler(2205, func(req []byte) (proto.Message, error) {
    // return a trdmodifyorder.Response
})
if err := server.Start(); err != nil {
    t.Fatal(err)
}
defer server.Stop()

cli, cleanup := testutil.NewTestClient(t, server)
defer cleanup()

// ... exercise the API ...
server.AssertProtoID(t, 2205)
```

`internal/testutil/mock` is a fuller variant that also supports RSA-encrypted
handshakes.

## CI gates

`.github/workflows/ci.yml` runs on every push and PR: `go build`,
strict `gofmt -l .`, `go vet`, `go test -race -count=1`, a coverage artifact,
and the README-translation guard. `govulncheck.yml` and `codeql.yml` run on
push/PR and weekly. See [CONTRIBUTING.md](../CONTRIBUTING.md).

## Writing tests

- Prefer table-driven tests.
- Unit-test validation and parsing directly (see `pkg/trd/validation_test.go`).
- Use the mock server for request/response wrapping (`test/trd_api`).
- Cover the error paths — including nil `S2C`/`Header` responses and timeouts.
