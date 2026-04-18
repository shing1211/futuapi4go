# Contributing to futuapi4go

> Thank you for your interest in contributing. This SDK is a Go library for Futu OpenD. Familiarity with Go best practices is expected.

## Code of Conduct

Please be respectful and professional in all interactions. See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## Development Setup

```bash
git clone https://gitee.com/shing1211/futuapi4go.git
cd futuapi4go
go mod download
```

## Building and Testing

```bash
# Build
go build ./...

# Run tests (core unit tests — no OpenD required)
go test ./client/... ./pkg/... ./internal/client/

# Run all tests (includes tests requiring real OpenD)
go test ./...

# Run with race detector
go test -race ./client/... ./pkg/... ./internal/client/

# Run go vet
go vet ./...
```

## Code Standards

- **Format**: Code must be formatted with `gofmt`. Run `go fmt ./...` before committing.
- **Linting**: All `go vet` checks must pass.
- **Error handling**: Do not ignore errors with `_`. All errors must be handled or explicitly propagated.
- **Context**: Use `context.Context` for cancellation and timeouts.
- **Tests**: New features should include tests. Fixes should include a test that would have caught the bug.
- **Comments**: Document all exported functions and types. Internal helpers may be uncommented.

## Commit Message Format

Use this format:

```
<package>: <short description>

Optional longer explanation.
```

Examples:

```
push: fix ParseUpdateKL unmarshal into S2C directly
internal/client: fix logf nil logger panic
docs: update CHANGELOG for v0.6.1
```

Types: `fix:`, `feat:`, `docs:`, `test:`, `chore:`, `refactor:` (Conventional Commits are recommended).

## Pull Request Guidelines

1. Fork the repository and create a feature branch (`git checkout -b fix/your-issue`).
2. Make your changes. Ensure `go build ./...` and `go test` pass.
3. Update [docs/CHANGELOG.md](docs/CHANGELOG.md) under the `[Unreleased]` section for user-facing changes.
4. Open a Pull Request with a clear description of the fix or feature.
5. Link any related issues.

## What to Contribute

- Bug fixes with tests
- Missing API wrappers (check `pkg/qot/` and `pkg/trd/` for patterns)
- Performance improvements
- Documentation improvements
- Additional examples in `cmd/examples/`

## Related Documentation

- [docs/DEVELOPER.md](docs/DEVELOPER.md) — Architecture and implementation patterns
- [docs/TESTING.md](docs/TESTING.md) — Testing guide
- [docs/API_REFERENCE.md](docs/API_REFERENCE.md) — API reference

## License

By contributing, you agree that your contributions will be licensed under the Apache License 2.0.
