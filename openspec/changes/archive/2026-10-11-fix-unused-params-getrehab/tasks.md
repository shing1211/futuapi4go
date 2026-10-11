# Tasks

## 1. Implement

- [ ] 1.1 Remove `_ = ctx`, `_ = c`, `_ = market`, `_ = code` suppression lines from `client/quote_api.go`
  - Verify: `go vet ./client/...` passes with no new errors

## 2. Verify

- [ ] 2.1 Run `go vet ./...` and verify no errors
  - Verify: `go vet ./...` exits 0
- [ ] 2.2 Run `go build ./...` and verify build succeeds
  - Verify: `go build ./...` exits 0
- [ ] 2.3 Create PR with `Closes #43`, request Mary (shing1211) review

## Workflow follow-up

- Merge PR after approval
- Archive change: `openspec archive fix-unused-params-getrehab`
