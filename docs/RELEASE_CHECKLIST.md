# FutuAPI4Go Release Checklist / 發佈檢查清單

## Pre-Release Verification / 發佈前驗證

### Code Quality / 代碼質量
- [ ] `go vet ./...` passes with no warnings
- [ ] `go test -race ./...` passes with no races detected
- [ ] `go test ./...` passes with ≥80% coverage
- [ ] No `panic()` calls in production code
- [ ] All errors wrapped with `%w` for proper error chaining
- [ ] All public functions have Go doc comments

### API Compatibility / API 兼容性
- [ ] All exported functions work with real Futu OpenD
- [ ] No breaking changes to public API signatures (or documented in CHANGELOG)
- [ ] Backward compatibility verified (if applicable)

### Testing / 測試
- [ ] Unit tests pass for all packages
- [ ] Integration tests pass against real OpenD (with FUTU_INTEGRATION_TESTS=1)
- [ ] Example programs compile successfully
- [ ] Simulator compiles and runs correctly
- [ ] No flaky tests (tests pass consistently on multiple runs)

### Documentation / 文檔
- [ ] README.md updated with current version and features
- [ ] CHANGELOG.md updated with release notes
- [ ] USER_GUIDE.md updated if API changed
- [ ] PRODUCTION_PLAN.md updated with current progress
- [ ] All new public functions have Go doc comments
- [ ] Examples updated if API changed

### Build / 構建
- [ ] `go build ./...` succeeds
- [ ] `go build ./cmd/...` succeeds for all example programs
- [ ] Cross-compilation tested (if applicable)
- [ ] Version string updated in `internal/client/version.go`

### Security / 安全
- [ ] No hardcoded credentials or API keys
- [ ] RSA encryption works if configured
- [ ] No debug logging in production code
- [ ] No `fmt.Printf` statements for sensitive data

### Performance / 性能
- [ ] No goroutine leaks (verified with `go test -race`)
- [ ] Connection pool works correctly
- [ ] Metrics tracking works (GetMetrics returns valid data)
- [ ] Reconnection works after network interruption

---

## Release Process / 發佈流程

1. **Update version** in `internal/client/version.go`
2. **Update CHANGELOG.md** with release notes
3. **Run full test suite**: `go test -race -cover ./...`
4. **Build examples**: `go build ./cmd/examples/...`
5. **Tag release**: `git tag -a v0.4.0 -m "Release v0.4.0"`
6. **Push tag**: `git push origin v0.4.0`
7. **Create Gitee release** with release notes

---

## Post-Release / 發佈後

- [ ] Monitor for issues/bugs reported
- [ ] Update PRODUCTION_PLAN.md progress
- [ ] Announce release to users (if applicable)

---

**Last Updated**: 2026-04-08
**Version**: 0.4.0-dev
