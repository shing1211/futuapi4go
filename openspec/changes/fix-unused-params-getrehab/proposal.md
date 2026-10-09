# Proposal

## Why

`GetRehab` in `client/quote_api.go:1835-1838` suppresses four parameters (`ctx`, `c`, `market`, `code`) with `_ = var` assignments. These suppressions violate the AGENTS.md error-handling rule: errors (and unused values) are never silently swallowed. Since `GetRehab` is a deprecated stub that always returns an error, the parameters are dead code.

## What Changes

- Remove `_ = ctx`, `_ = c`, `_ = market`, `_ = code` suppression lines from `GetRehab`
- No behavioral change — function still returns the same error
- No API change — parameter signature unchanged for backward compatibility

## Capabilities

### New Capabilities
None — this is a pure refactor with no new behavior.

### Modified Capabilities
None — no spec-level behavior changes.

## Impact

- **File:** `client/quote_api.go`
- **Proto affected:** none
- **API compatibility:** `client/` API unchanged (parameter signature preserved)
