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

## Session Workflow (START HERE EVERY TIME)

### Step 1: Open the Implementation Plan
```bash
# Read the current state
cat docs/IMPLEMENTATION_PLAN.md
```

This document contains **24 enhancement items** across 5 phases. Each item has:
- Unique ID (P1-1, P1-2, etc.)
- Severity (CRITICAL/HIGH/MEDIUM/LOW)
- Status (Pending/In Progress/Done)
- Exact file location
- Code example showing before/after
- Definition of Done checklist

### Step 2: Select the Next Item
**Always work in phase order.** Within a phase, prioritize by severity:
1. **CRITICAL** → Fix first (security, crash bugs)
2. **HIGH** → Next (performance, major usability)
3. **MEDIUM** → Then
4. **LOW** → Last

**Phase 1 (Critical Security & Correctness) - ✅ COMPLETE**
All Phase 1 items have been implemented and verified:
- ✅ P1-1: Connection Pool Race Condition - `internal/client/pool.go`
- ✅ P1-2: Packet Length Overflow Check - `internal/client/conn.go`
- ✅ P1-3: Sensitive Data Logging Protection - `pkg/trd/trade.go`
- ✅ P1-4: Goroutine Leaks in Push Subscription - `pkg/chanpkg/chan.go`
- ✅ P1-5: Buffered I/O for Packet Reading - `internal/client/conn.go`
- ✅ P1-6: Input Validation on All Public APIs - All
- ✅ P1-7: Proto Field Nil Checks - All response parsing

**Current Priority Queue (Phase 2):**
- P2-1: Typed Enum Parameters Everywhere (In Progress)
- P2-2: Builder Pattern for Requests (Done)
- P2-3: Convenience Wrappers for Common Operations (Done)
- P2-4: Market Auto-Detection Helper (Done)

### Step 3: Update Status Before Starting
Before writing any code:
1. Change the item's status from `⚪ Pending` → `🔄 In Progress`
2. Add your agent name to **Assignee** (if applicable)
3. Commit this change to the plan file first:
   ```bash
   git add docs/IMPLEMENTATION_PLAN.md
   git commit -m "docs: start P1-1: Connection Pool Race Condition"
   ```

### Step 4: Implement the Fix
Follow these rules:
- **Copy code patterns** exactly from the implementation plan (they're pre-reviewed)
- **Write unit tests** BEFORE implementation (TDD style) when possible
- **Follow existing conventions** in the file (naming, spacing, error patterns)
- **Use the existing `wrapError` pattern** for all API errors
- **Apply nil guards** to all list iterations and proto field access

### Step 5: Verify (Must Do This!)
Run these commands before marking anything complete:
```bash
go build ./...    # MUST PASS - no build errors
go vet ./...      # MUST PASS - no linter issues
go test -race ./...  # MUST PASS - no race conditions
```

If any of these fail, fix them before proceeding.

### Step 6: Update Status & Documentation
1. Change the item's status from `🔄 In Progress` → `✅ Done`
2. Add completed date to Assignee field
3. Update `docs/CHANGELOG.md` under `[Unreleased]`
4. Commit all changes with descriptive message:
   ```bash
   git add docs/IMPLEMENTATION_PLAN.md docs/CHANGELOG.md internal/client/pool.go internal/client/pool_test.go
   git commit -m "fix: add mutex protection to connection pool (P1-1)"
   ```

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

## Phase Summary & Targets

| Phase | Name | Items | Effort | Target Version | Breaking? |
|-------|------|-------|--------|----------------|-----------|
| **Phase 1** | Critical Security & Correctness | 7 | 15-20 hrs | v0.2.1 | No |
| **Phase 2** | Ease of Use - Type Safety | 4 | 20-25 hrs | v0.3.0 | **YES** |
| **Phase 3** | Infrastructure Improvements | 4 | 15-20 hrs | v0.3.1 | No |
| **Phase 4** | Testing & Validation | 4 | 15-20 hrs | v0.3.2 | No |
| **Phase 5** | Polish & Documentation | 5 | 10-15 hrs | v0.4.0 | Partial |
| **Phase 6** | World-Class SDK | 9 | 13-15 hrs | v0.5.1 | Partial |
| **TOTAL** | | **24** | **75-100 hrs** | | |

### Phase Gates
Before starting a new phase:
1. All items in previous phase marked `✅ Done`
2. All tests pass with `-race` flag
3. `go build ./...` and `go vet ./...` pass
4. CHANGELOG.md updated with all completed items
5. Demo project examples updated (if breaking changes)

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
- [ ] `FutuError` type used with `Unwrap()` support (Phase 3+)

### Memory & Performance
- [ ] Buffered I/O used for packet reads/writes
- [ ] `sync.Pool` used for hot-path allocations (Phase 3+)
- [ ] No unnecessary allocations in tight loops
- [ ] Max packet size checks prevent overflow

### Security
- [ ] Sensitive fields (`PwdMD5`) use `SensitiveString` type
- [ ] No sensitive data logged or printed
- [ ] TLS option available for connections (Phase 3+)
- [ ] Input validation prevents injection attacks

### Documentation
- [ ] New public functions have GoDoc comments
- [ ] CHANGELOG.md updated under `[Unreleased]`
- [ ] IMPLEMENTATION_PLAN.md status updated
- [ ] MIGRATION_GUIDE.md updated if breaking changes

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
| `pkg/chanpkg/chan.go` | Channel-based push | Done channels, WaitGroup, leaks |

---

## Adding a New API

1. Confirm the proto in `api/proto/`
2. Run `./scripts/regen-all-protos.ps1`
3. Add the wrapper function in `pkg/qot/` or `pkg/trd/`:
   - Context as FIRST parameter
   - Input validation at entry
   - Use `RequestContext()` pattern
   - Use `wrapError()` for proto errors
   - Nil guards on all list iteration
4. Add a public helper in `client/client.go` if it simplifies usage
5. Add unit tests with table-driven edge cases
6. Update `docs/CHANGELOG.md` under `[Unreleased]`
7. Update `docs/IMPLEMENTATION_PLAN.md` if this was a planned item
8. Verify: `go build ./... && go vet ./... && go test -race ./...`

---

## Breaking Change Handling Process

For Phase 2 (typed enums) and other breaking changes:

1. **Update MIGRATION_GUIDE.md** with:
   - Before/after code examples
   - Search/replace patterns users can apply
   - List of all affected functions

2. **Update demo project FIRST** (before SDK changes) to:
   - Have working code with old API (baseline)
   - Apply changes incrementally
   - Verify all examples still work

3. **Make SDK changes** and update all internal callers
4. **Update package version** in documentation (v0.3.0 for Phase 2)
5. **Update both CHANGELOG.md files** (SDK and demo)

---

## Demo Project Coordination

The demo project at `../futuapi4go-demo` must be kept in sync:

| When | Action |
|------|--------|
| After Phase 1 | No action needed (no breaking changes) |
| After Phase 2 (typed enums) | **REQUIRED** - Update all examples to remove `int32()` casts |
| After Phase 3 | Update examples using new convenience wrappers |
| After Phase 4 | Update CI integration tests |
| After Phase 5 | Add new tutorial examples |

### Demo Project Quick Check
```bash
cd ../futuapi4go-demo
go build ./...    # Must pass after all changes
go vet ./...      # Must pass
```

---

## Status Emoji Conventions

Use these consistently in IMPLEMENTATION_PLAN.md:

| Emoji | Meaning |
|-------|---------|
| `⚪` | Pending — not started |
| `🔄` | In Progress — actively working on this |
| `⚠️` | Blocked — needs input or dependency |
| `✅` | Done — implemented, tested, verified |
| `❌` | Rejected — will not implement |

---

## Official Documentation References

- **Implementation Plan:** `docs/IMPLEMENTATION_PLAN.md` (MAIN WORK TRACKER)
- **Changelog:** `docs/CHANGELOG.md`
- **Migration Guide:** `docs/MIGRATION_GUIDE.md`
- **API Reference:** `docs/API_REFERENCE.md`
- **Proto Reference:** https://openapi.futunn.com/mds/Futu-API-Doc-zh-Proto.md
- **Go module:** `github.com/shing1211/futuapi4go` (current: v0.5.1)

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
1. Did you regenerate protos after proto file change? `./scripts/regen-all-protos.ps1`
2. Check import paths match package structure

---

## Session Checklist (Review Before Finishing)

Before ending a work session, confirm:

- [ ] All changes build: `go build ./...` ✅
- [ ] All changes lint: `go vet ./...` ✅
- [ ] All tests pass with race detection: `go test -race ./...` ✅
- [ ] IMPLEMENTATION_PLAN.md status updated (✅ Done or 🔄 In Progress)
- [ ] CHANGELOG.md updated with completed items
- [ ] All changes committed with descriptive messages
- [ ] Demo project updated (if breaking changes)
