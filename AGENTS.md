# futuapi4go Operational Guide

This is the main reference for AI agent development sessions. Follow these instructions exactly for consistent, high-quality work.

---

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

---

## Build & Verify Commands

```bash
# Basic build (always run first)
go build ./...

# Lint (must be clean before commit)
go vet ./...

# Full test suite with race detection
go test -race ./...

# Run specific tests only
go test -race ./pkg/trd/... -run PlaceOrder

# Benchmark (for performance items)
go test -bench=. -benchmem ./internal/client/...
```

---

## Adding a New API

1. Confirm the proto in `api/proto/`
2. Run `./scripts/regen-all-protos.sh`
3. Add the wrapper function in `pkg/qot/` or `pkg/trd/`:
   - Context as FIRST parameter
   - Input validation at entry
   - Use `RequestContext()` pattern
   - Use `wrapError()` for proto errors
   - Nil guards on all list iteration
4. Add a public helper in `client/client.go` if it simplifies usage
5. Add unit tests with table-driven edge cases
6. Update `docs/CHANGELOG.md` under `[Unreleased]`
7. Verify: `go build ./... && go vet ./... && go test -race ./...`

---

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

---

## Enhanced Code Review Checklist

### Concurrency & Thread Safety
- [ ] All shared state access protected by `sync.Mutex` or `sync.RWMutex`
- [ ] `defer mu.Unlock()` pattern used consistently
- [ ] No goroutine leaks (all goroutines have exit mechanism via `done` channel or `WaitGroup`)
- [ ] Race detection tests pass: `go test -race ./...`
- [ ] Connection pool `Get()`/`Put()` are thread-safe

### Input Validation
- [ ] Every public function validates inputs at entry
- [ ] `nil` request check: `if req == nil { return error }`
- [ ] Zero value checks: `AccID != 0`, `len(Code) > 0`
- [ ] Boundary checks: string lengths, price/qty ranges
- [ ] Negative values rejected where appropriate

### Proto Safety
- [ ] **NO** `GetXxx()` method calls on proto messages
- [ ] All field access uses direct nil checks: `if kl.Time != nil { val = *kl.Time }`
- [ ] List iteration has nil guard: `if item == nil { continue }`
- [ ] `RetType` always checked BEFORE accessing `S2C`
- [ ] `S2C` nil check before field access

### Context Usage
- [ ] `context.Context` is FIRST parameter to all public APIs
- [ ] Context passed through to `RequestContext()` call
- [ ] Context cancellation respected in all I/O paths
- [ ] No `context.Background()` used inside library functions (accept context from caller)

### Error Handling
- [ ] Errors are never swallowed with `_`
- [ ] Use `wrapError()` helper for all proto API errors
- [ ] Error messages include function name for debugging
- [ ] `FutuError` type used with `Unwrap()` support
- [ ] **Do NOT use `retry.Do()` with trading operations** (`PlaceOrder`, `ModifyOrder`, `CancelOrder`) — non-idempotent order operations must never be retried

### Memory & Performance
- [ ] Buffered I/O used for packet reads/writes
- [ ] `sync.Pool` used for hot-path allocations (Phase 3+)
- [ ] No unnecessary allocations in tight loops
- [ ] Max packet size checks prevent overflow

### Security
- [ ] Sensitive fields (`PwdMD5`) use `SensitiveString` type
- [ ] No sensitive data logged or printed
- [ ] Input validation prevents injection attacks

### Documentation
- [ ] New public functions have GoDoc comments
- [ ] CHANGELOG.md updated under `[Unreleased]`
- [ ] All changes committed with descriptive messages

---

## Key Entry Points

| File | Purpose | Watch For |
|------|---------|-----------|
| `client/client.go` | Public API wrapper | Context passing, enum types |
| `internal/client/client.go` | Connection, serial numbers, reconnect | Goroutine safety, state management |
| `internal/client/conn.go` | Raw TCP packet I/O | Buffering, overflow, timeouts |
| `internal/client/pool.go` | Connection pool | Mutex protection, race conditions |
| `pkg/qot/quote.go` | All market data APIs | Input validation, nil guards |
| `pkg/trd/trade.go` | All trading APIs | Sensitive data, validation |
| `pkg/sys/system.go` | System APIs | Error handling, context |
| `pkg/push/qot_push.go` | Push notification parsers | Goroutine cleanup |
| `pkg/push/chan/chan.go` | Channel-based push | Done channels, WaitGroup, leaks |

---

## Release Process

```bash
# 1. Update CHANGELOG.md: move [Unreleased] → [x.y.z] - YYYY-MM-DD
# 2. Commit changes
git add -A && git commit -m "release vx.y.z: summary"

# 3. Tag and push to both remotes
git tag vx.y.z
git push origin vx.y.z
git push github vx.y.z

# 4. Create GitHub release
make release          # macOS/Linux — requires goreleaser
# or:
gh release create vx.y.z --title "vx.y.z" --notes "see CHANGELOG.md"
```

Prerequisites (macOS/Linux): `go install github.com/goreleaser/goreleaser@latest`

---

## Official Documentation References

- **API Reference:** See README.md "Full API Reference" section
- **Changelog:** `docs/CHANGELOG.md`
- **Developer Guide:** This file (AGENTS.md)
- **Testing Guide:** See README.md "Testing" section
- **Enhancement Plan:** `ENHANCEMENT_PLAN.md` (advanced features — application-level, not core SDK)
- **Proto Reference:** https://openapi.futunn.com/futu-api-doc/en/
- **Go module:** `github.com/shing1211/futuapi4go` (current: v0.11.0)

---

## Troubleshooting Common Issues

### Build Failures After Changes
1. Did you forget to add new file to git? `git status`
2. Did you run `go mod tidy` if adding new imports?
3. Check file naming: Go ignores `*_test.go` in non-test builds

### Test Failures
1. Run with `-v` flag to see verbose output
2. Run specific test: `go test -run TestName ./...`
3. Race detection: Always run with `-race` flag for concurrent code

### Proto Issues
1. Did you regenerate protos after proto file change? `./scripts/regen-all-protos.sh`
2. Check import paths match package structure

---

## Session Checklist (Review Before Finishing)

Before ending a work session, confirm:

- [ ] All changes build: `go build ./...` ✅
- [ ] All changes lint: `go vet ./...` ✅
- [ ] All tests pass with race detection: `go test -race ./...` ✅
- [ ] CHANGELOG.md updated with completed items under `[Unreleased]`
- [ ] All changes committed with descriptive messages
- [ ] Demo project updated (if breaking changes)

---

*Last updated: 2026-05-21*

<!-- gitnexus:start -->
# GitNexus — Code Intelligence

This project is indexed by GitNexus as **futuapi4go** (20427 symbols, 51795 relationships, 186 execution flows). Use the GitNexus MCP tools to understand code, assess impact, and navigate safely.

> If any GitNexus tool warns the index is stale, run `npx gitnexus analyze` in terminal first.

## Always Do

- **MUST run impact analysis before editing any symbol.** Before modifying a function, class, or method, run `gitnexus_impact({target: "symbolName", direction: "upstream"})` and report the blast radius (direct callers, affected processes, risk level) to the user.
- **MUST run `gitnexus_detect_changes()` before committing** to verify your changes only affect expected symbols and execution flows.
- **MUST warn the user** if impact analysis returns HIGH or CRITICAL risk before proceeding with edits.
- When exploring unfamiliar code, use `gitnexus_query({query: "concept"})` to find execution flows instead of grepping. It returns process-grouped results ranked by relevance.
- When you need full context on a specific symbol — callers, callees, which execution flows it participates in — use `gitnexus_context({name: "symbolName"})`.

## Never Do

- NEVER edit a function, class, or method without first running `gitnexus_impact` on it.
- NEVER ignore HIGH or CRITICAL risk warnings from impact analysis.
- NEVER rename symbols with find-and-replace — use `gitnexus_rename` which understands the call graph.
- NEVER commit changes without running `gitnexus_detect_changes()` to check affected scope.

## Resources

| Resource | Use for |
|----------|---------|
| `gitnexus://repo/futuapi4go/context` | Codebase overview, check index freshness |
| `gitnexus://repo/futuapi4go/clusters` | All functional areas |
| `gitnexus://repo/futuapi4go/processes` | All execution flows |
| `gitnexus://repo/futuapi4go/process/{name}` | Step-by-step execution trace |

## CLI

| Task | Read this skill file |
|------|---------------------|
| Understand architecture / "How does X work?" | `.claude/skills/gitnexus/gitnexus-exploring/SKILL.md` |
| Blast radius / "What breaks if I change X?" | `.claude/skills/gitnexus/gitnexus-impact-analysis/SKILL.md` |
| Trace bugs / "Why is X failing?" | `.claude/skills/gitnexus/gitnexus-debugging/SKILL.md` |
| Rename / extract / split / refactor | `.claude/skills/gitnexus/gitnexus-refactoring/SKILL.md` |
| Tools, resources, schema reference | `.claude/skills/gitnexus/gitnexus-guide/SKILL.md` |
| Index, status, clean, wiki CLI commands | `.claude/skills/gitnexus/gitnexus-cli/SKILL.md` |

<!-- gitnexus:end -->
