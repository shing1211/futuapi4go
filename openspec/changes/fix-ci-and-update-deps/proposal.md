# Proposal

## Why

`main` has carried a failing CI build since v0.20.0 (commit `d435f9b`, 2026-09-18) and nothing has fixed it. The `build & test` workflow has failed on every commit for 10+ days, including the four dependabot PRs merged since. The cause is three files that are not `gofmt`-clean, introduced by the v0.20.0 WebSocket work.

The gate caught this correctly, but could not act: `main` has `enforce_admins: false`, so direct pushes bypass all four required status checks. GitHub confirms the bypass on every push (`Bypassed rule violations for refs/heads/main: 4 of 4 required status checks are expected`). The result is a release where the shipped tag's build is red, and no mechanism to stop the next one.

We fix the formatting, update the four outdated indirect dependencies, and route this change through a PR so the status checks gate the merge for the first time.

## What Changes

- **`gofmt` repair** — reformat three files that v0.20.0 left misaligned, in the `ClientOptions` struct field block and its tests. Whitespace only; no identifier, type, or behavior change.
- **Indirect dependency updates** — four modules reach their latest versions: `prometheus/client_model` v0.6.2→v0.6.3, `prometheus/common` v0.70.1→v0.72.0, `prometheus/procfs` v0.21.1→v0.22.0, `golang.org/x/sys` v0.47.0→v0.48.0.
- **Release `v0.21.0`** — move the `[Unreleased]` CHANGELOG section to `## [0.21.0] - 2026-09-29`, and update the version badge and `go get` line in the English `README.md`. The remaining 18 files carrying v0.20.0 are swept in a post-release commit, matching the precedent set by `e8b91e7`, `55a4c35`, and `746120e`.
- **Delivery via PR** — the work lands on `fix/ci-and-deps-v0.21.0` and merges only once all four required status checks pass, so the gate has real authority for the first time.

No breaking changes.

## Capabilities

### New Capabilities

- `build-compliance`: The repository's build gates hold on `main`, and a release is only cut from a tree where they pass. Covers the requirement that `gofmt -l .` is empty, that all four required status checks gate merges, and that release tags are cut from verified-green commits.

### Modified Capabilities

None. This is the first spec in the project; there are no existing capabilities to amend.

## Impact

**Affected code**

- `internal/client/client.go` — `ClientOptions` field alignment in the `// WebSocket` block (`WSSecretKey`, `WSReconnect` carry stray extra spaces)
- `internal/client/client_test.go` — same misalignment
- `pkg/constant/errors_test.go` — same misalignment

**Public API compatibility: none.** All three edits are whitespace within existing struct definitions. No exported identifier, signature, or type changes in `client/` or `pkg/`. The `WSSecretKey` and `WSReconnect` fields added in v0.20.0 keep their names, types, and semantics.

**Protos affected: none.** No change touches `api/proto/`, `pkg/pb/`, or any wrapper. All 184 protos and their typed wrappers are untouched.

**Dependencies** — four indirect modules move; `go.sum` grows from 47 to 55 lines, purely additive. All six direct requires are already current: `gorilla/websocket` v1.5.3, `prometheus/client_golang` v1.24.1, the three `go.opentelemetry.io/otel` v1.46.0 modules, and `google.golang.org/protobuf` v1.36.12. These modules back Prometheus metrics and OS syscalls; they are not on any proto or trading path.

**Verification** — all six gates pass locally on the merged result: `gofmt -l .` empty, `go build ./...`, `go vet ./...`, `go test -race -count=1 ./...` (29/29 packages), `govulncheck ./...` (no vulnerabilities), `make docs-check`.
