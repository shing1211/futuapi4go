# Contributing to futuapi4go

> Thank you! Even fixing a typo in a comment counts as a valuable contribution.

## Code of Conduct

Treat everyone with respect. See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## Development Setup

```bash
git clone https://github.com/shing1211/futuapi4go.git
cd futuapi4go
go mod download
go build ./...
```

## Build, Lint & Test

```bash
# Build everything
go build ./...

# Lint
go vet ./...

# Run unit tests (no OpenD required)
go test ./client/... ./pkg/... ./internal/client/

# Run all tests (some require OpenD)
go test ./...

# Race detector
go test -race ./client/... ./pkg/... ./internal/client/
```

## Code Standards

| Rule | Why |
|------|-----|
| Run `go fmt ./...` before committing | Consistent formatting |
| `go vet ./...` must pass | Catch common mistakes |
| Never ignore errors with `_` | Errors are values — handle them |
| Use `context.Context` for all APIs | Enables cancellation and timeouts |
| New features need tests | Bug fixes need regression tests |
| Document exported functions | Your future self will thank you |

## Commit Messages

Use this format (Conventional Commits):

```
<type>: <short description>

Optional longer explanation.
```

Types: `fix:`, `feat:`, `docs:`, `test:`, `chore:`, `refactor:`

Examples:
```
fix: restore nil guard in ParseUpdateOrderBook
feat: add GetLoginUserID and IsEncrypt helpers
docs: update CHANGELOG for v0.9.0
client: add CanSendProto connection state check
```

## Pull Request Process

1. Fork and create a branch: `git checkout -b fix/your-issue`
2. Make your changes
3. Ensure `go build ./...` and `go vet ./...` both pass
4. Update `docs/CHANGELOG.md` under `[Unreleased]` for user-facing changes
5. Open a PR with a clear description — link any related issue
6. A maintainer will review and may request changes

## What to Contribute

- Bug fixes (bonus: include a test that would have caught it)
- Missing API wrappers — check `pkg/qot/` and `pkg/trd/` for patterns
- Performance improvements
- Documentation improvements (this README, code comments, examples)
- New examples in `cmd/examples/`

## Related Docs

- [docs/DEVELOPER.md](docs/DEVELOPER.md) — Architecture and implementation patterns
- [docs/TESTING.md](docs/TESTING.md) — Testing guide and fixtures
- [docs/API_REFERENCE.md](docs/API_REFERENCE.md) — Complete API reference
- [ROADMAP.md](ROADMAP.md) — What's coming next

## License

By contributing, you agree your work will be licensed under Apache License 2.0.
