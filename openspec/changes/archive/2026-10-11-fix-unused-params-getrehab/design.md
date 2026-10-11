# Design

## Context

`GetRehab` in `client/quote_api.go:1835-1841` is a deprecated stub function. It currently suppresses four parameters with `_ = var` assignments:

```go
func GetRehab(ctx context.Context, c *Client, market constant.Market, code string) ([]*RehabInfo, error) {
    _ = ctx
    _ = c
    _ = market
    _ = code
    return nil, fmt.Errorf("GetRehab: removed in Futu v10.6 — use RequestRehab instead")
}
```

See proposal.md for motivation.

## Goals / Non-Goals

**Goals:**
- Remove unused parameter suppressions to comply with AGENTS.md error-handling rule
- Keep parameter signature unchanged for backward API compatibility

**Non-Goals:**
- No behavioral change (function still returns the same error)
- No new tests required (no logic change)

## Decisions

### Decision: Remove suppressions, keep signature

**Choice:** Remove `_ = ctx`, `_ = c`, `_ = market`, `_ = code` lines while keeping the parameter names.

**Rationale:**
- Suppressions hide potential bugs and violate the project's error-handling convention
- Removing parameters entirely would be a **BREAKING** API change
- The function signature must remain `GetRehab(ctx context.Context, c *Client, market constant.Market, code string)` for backward compatibility

**Alternatives considered:**
1. **Remove parameters entirely** — BREAKING API change; rejected
2. **Keep suppressions** — Violates AGENTS.md rule; rejected
3. **Document suppressions** — Does not fix the underlying issue; rejected

## Risks / Trade-offs

| Risk | Mitigation |
|------|------------|
| None | Simple removal of dead code; no behavioral impact |

## Migration Plan

1. Remove 4 suppression lines from `client/quote_api.go`
2. Run `go vet ./...` to verify no new errors
3. Create PR, request review
4. Merge on approval
