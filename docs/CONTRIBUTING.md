# Contributing to futuapi4go

Thank you for your interest in contributing to futuapi4go! This guide will help you get started.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Code Standards](#code-standards)
- [Testing Requirements](#testing-requirements)
- [Pull Request Process](#pull-request-process)
- [Commit Message Guidelines](#commit-message-guidelines)
- [Architecture Overview](#architecture-overview)
- [Areas for Contribution](#areas-for-contribution)
- [Getting Help](#getting-help)

---

## 🤝 Code of Conduct

### Our Pledge

We pledge to make participation in this project a harassment-free experience for everyone, regardless of age, body size, disability, ethnicity, gender identity and expression, level of experience, nationality, personal appearance, race, religion, or sexual identity and orientation.

### Our Standards

**Positive behavior:**
- Using welcoming and inclusive language
- Being respectful of differing viewpoints
- Accepting constructive criticism gracefully
- Focusing on what is best for the community
- Showing empathy towards other contributors

**Unacceptable behavior:**
- Trolling, insulting/derogatory comments, or personal attacks
- Public or private harassment
- Publishing others' private information without permission
- Other conduct which could reasonably be considered inappropriate

---

## 🚀 Getting Started

### Prerequisites

- **Go 1.21+** - [Download](https://golang.org/dl/)
- **Git** - Version control
- **Futu OpenD** (optional) - For integration testing
- **Protobuf compiler** (optional) - For protocol modifications

### Setup Development Environment

```bash
# 1. Fork the repository on GitHub

# 2. Clone your fork
git clone https://github.com/YOUR_USERNAME/futuapi4go.git
cd futuapi4go

# 3. Add upstream remote
git remote add upstream https://github.com/shing1211/futuapi4go.git

# 4. Install dependencies
go mod download

# 5. Run tests to verify setup
go test ./...

# 6. Build examples
cd cmd/examples && go build ./...
```

### Project Structure

```
futuapi4go/
├── api/proto/           # Protobuf definitions
├── cmd/
│   ├── examples/        # Example programs
│   └── simulator/       # Mock OpenD server
├── internal/
│   └── client/          # Core client implementation (TCP + WebSocket)
├── pkg/
│   ├── qot/             # Market Data APIs
│   ├── trd/             # Trading APIs
│   ├── sys/             # System APIs
│   ├── push/            # Push handlers
│   └── pb/              # Generated protobuf code
├── test/                # Test suite
└── docs/                # Documentation
```

See [DEVELOPER.md](DEVELOPER.md) for detailed architecture documentation.

---

## 🔄 Development Workflow

### 1. Create Feature Branch

```bash
# Sync with upstream
git fetch upstream
git checkout main
git merge upstream/main

# Create feature branch
git checkout -b feature/your-feature-name
```

### 2. Make Changes

```bash
# Edit code
# ...

# Format code
gofmt -w .

# Run linter
go vet ./...

# Run tests
go test ./...
```

### 3. Commit Changes

```bash
git add .
git commit -m "feat: add your feature description"
```

See [Commit Message Guidelines](#commit-message-guidelines) below.

### 4. Push and Create PR

```bash
# Push to your fork
git push origin feature/your-feature-name

# Create Pull Request on GitHub
# Navigate to your fork and click "Create Pull Request"
```

---

## 📝 Code Standards

### Go Code Style

**Use `gofmt`:**
```bash
# Format all Go files
gofmt -w .

# Check formatting
gofmt -l .
```

**Naming Conventions:**
- **Exported types**: `PascalCase` (e.g., `BasicQot`, `GetKLRequest`)
- **Unexported types**: `camelCase` (e.g., `clientOptions`)
- **Constants**: `PascalCase` for exported, `camelCase` for unexported
- **Interfaces**: `-er` suffix (e.g., `Reader`, `Writer`)

### Code Organization

**Package structure:**
```go
// Package qot provides market data APIs.
package qot

import (
    // Standard library first
    "fmt"
    "time"
    
    // Third-party imports
    "google.golang.org/protobuf/proto"
    
    // Internal imports last
    "github.com/shing1211/futuapi4go/internal/client"
    "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

// Exported types and functions
type GetKLRequest struct {
    // Fields...
}

// GetKL retrieves K-line data.
func GetKL(c *client.Client, req *GetKLRequest) (*GetKLResponse, error) {
    // Implementation...
}
```

**Function documentation:**
```go
// GetBasicQot retrieves real-time quotes for one or more securities.
//
// Example:
//
//     securities := []*qotcommon.Security{
//         {Market: &market, Code: ptrStr("00700")},
//     }
//     quotes, err := qot.GetBasicQot(cli, securities)
//
func GetBasicQot(c *client.Client, securityList []*qotcommon.Security) ([]*BasicQot, error) {
    // Implementation...
}
```

### Error Handling

**Use wrapped errors:**
```go
if err != nil {
    return fmt.Errorf("GetBasicQot failed: %w", err)
}
```

**Define sentinel errors:**
```go
var (
    ErrNotConnected     = errors.New("not connected")
    ErrRequestTimeout   = errors.New("request timeout")
)
```

**Check errors with `errors.Is`/`errors.As`:**
```go
if errors.Is(err, ErrNotConnected) {
    // Handle not connected
}
```

### Context Usage

**Accept context as first parameter:**
```go
func (c *Client) WithContext(ctx context.Context) *Client {
    // Returns new client with context
}
```

**Respect context cancellation:**
```go
select {
case pkt := <-ch:
    return pkt, nil
case <-ctx.Done():
    return nil, ctx.Err()
}
```

---

## 🧪 Testing Requirements

### Mandatory Tests

All PRs must pass:

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...

# Run linter
go vet ./...
```

### Writing Tests

**Test file location:**
- Unit tests: Same package as code (`*_test.go`)
- Integration tests: `test/integration/`
- Benchmarks: `test/benchmark/`

**Test naming:**
```go
func TestGetBasicQot_HSI(t *testing.T)           // API test with HSI
func TestPlaceOrder_InsufficientFunds(t *testing.T)  // Error path test
func BenchmarkGetBasicQot_Mock(b *testing.B)     // Benchmark
```

**Use table-driven tests:**
```go
func TestGetKL(t *testing.T) {
    tests := []struct {
        name    string
        klType  int32
        wantErr bool
    }{
        {"Day", int32(KLType_Day), false},
        {"Week", int32(KLType_Week), false},
        {"Invalid", -1, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test logic...
        })
    }
}
```

**Use test fixtures:**
```go
import "github.com/shing1211/futuapi4go/test/fixtures"

// Always use realistic fixtures
quote := fixtures.HSIQuote()
```

See [TESTING.md](TESTING.md) for complete testing guide.

### Code Coverage

**Minimum coverage:** 80% for new code

```bash
# Check coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## 📤 Pull Request Process

### Before Submitting

1. **Update documentation** for new features
2. **Add tests** for new functionality
3. **Run all tests** and ensure they pass
4. **Update CHANGELOG.md** with your changes
5. **Squash commits** if multiple small commits

### PR Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix (non-breaking change)
- [ ] New feature (non-breaking change)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Refactoring (no functional changes)

## Testing
- [ ] All tests pass (`go test ./...`)
- [ ] Race detector passes (`go test -race ./...`)
- [ ] Linter passes (`go vet ./...`)
- [ ] Added tests for new functionality
- [ ] Tested with live OpenD (if applicable)

## Documentation
- [ ] Updated README.md (if needed)
- [ ] Updated API_REFERENCE.md (if needed)
- [ ] Added Go doc comments
- [ ] Updated CHANGELOG.md

## Additional Notes
Any additional information, screenshots, or context
```

### Review Process

1. **Automated checks** must pass
2. **Code review** by maintainers (1-2 reviewers)
3. **Address feedback** with follow-up commits
4. **Approval** required before merge
5. **Squash and merge** by maintainer

### Response Time

- **Initial review**: Within 3 business days
- **Follow-up**: Within 1 business day
- **Merge**: Within 1 day after approval

---

## ✍️ Commit Message Guidelines

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

| Type | Description | Example |
|------|-------------|---------|
| `feat` | New feature | `feat: add GetOptionChain API` |
| `fix` | Bug fix | `fix: correct OrderBook field names` |
| `docs` | Documentation | `docs: update API reference` |
| `style` | Formatting | `style: format with gofmt` |
| `refactor` | Code restructuring | `refactor: simplify client initialization` |
| `test` | Tests | `test: add HSI fixtures` |
| `chore` | Maintenance | `chore: update dependencies` |

### Examples

**Good:**
```
feat(qot): add GetOptionChain API

Implement option chain retrieval API (ProtoID 3209).
Includes support for call/put filtering and expiry date range.

Closes #123
```

**Bad:**
```
update stuff
```

### Rules

- **Subject line**: Max 72 characters, imperative mood
- **Body**: Explain WHAT changed and WHY (not HOW)
- **Footer**: Reference issues/PRs (`Closes #123`, `Fixes #456`)
- **Scope**: Optional, use package name (e.g., `qot`, `trd`, `test`)

---

## 🏗️ Architecture Overview

### Core Components

**Client Layer** (`internal/client/`):
- TCP connection management
- Request/response matching (serial numbers)
- Auto-reconnect with exponential backoff
- Keep-alive heartbeat
- Push notification dispatch

**API Layer** (`pkg/qot/`, `pkg/trd/`, `pkg/sys/`):
- High-level typed functions
- Protobuf marshaling/unmarshaling
- Error handling and validation
- Response transformation

**Protocol Layer** (`api/proto/`, `pkg/pb/`):
- Protobuf definitions (74 files)
- Generated Go code (74 packages)
- Custom binary framing (44-byte header)

### Design Principles

1. **Type safety**: Strong typing throughout
2. **Error handling**: No panics, all errors returned
3. **Thread safety**: Mutex and atomic operations
4. **Context support**: Cancellation and timeouts
5. **Minimal dependencies**: Only protobuf library
6. **Production-ready**: Comprehensive testing

See [DEVELOPER.md](DEVELOPER.md) for detailed architecture.

---

## 🎯 Areas for Contribution

### High Priority

- [ ] **WebSocket transport**: Complete implementation
- [ ] **More examples**: Advanced trading strategies
- [ ] **Performance**: Optimize hot paths
- [ ] **Documentation**: More tutorials and guides

### Medium Priority

- [ ] **Fuzz testing**: Protobuf parsing edge cases
- [ ] **Mock server**: Complete all handler implementations
- [ ] **Metrics**: OpenTelemetry integration
- [ ] **Rate limiting**: Client-side rate limiting

### Good First Issues

- [ ] **Documentation**: Fix typos, improve examples
- [ ] **Tests**: Add edge case coverage
- [ ] **Benchmarks**: Performance profiling
- [ ] **Code quality**: Linter suggestions

### Future Enhancements

- [ ] **GraphQL interface**: Alternative API
- [ ] **gRPC support**: Modern RPC protocol
- [ ] **Plugin system**: Custom strategies
- [ ] **Dashboard**: Web UI for monitoring

---

## ❓ Getting Help

### Resources

- **[README.md](README.md)**: Project overview
- **[API_REFERENCE.md](API_REFERENCE.md)**: Complete API docs
- **[TESTING.md](TESTING.md)**: Testing guide
- **[DEVELOPER.md](DEVELOPER.md)**: Development guide
- **[CHANGELOG.md](CHANGELOG.md)**: Version history

### Communication

- **Issues**: [GitHub Issues](https://github.com/shing1211/futuapi4go/issues)
- **Discussions**: [GitHub Discussions](https://github.com/shing1211/futuapi4go/discussions)
- **Email**: shing1211@users.noreply.github.com

### FAQs

**Q: How do I report a bug?**

A: Open an issue with:
- Description of the bug
- Steps to reproduce
- Expected vs actual behavior
- Go version and OS
- Code example (if applicable)

**Q: How do I request a feature?**

A: Open an issue with:
- Clear description of the feature
- Use cases and examples
- Benefits to the project
- Willingness to implement (optional)

**Q: How do I get help with my code?**

A: 
1. Check documentation first
2. Search existing issues
3. Open a discussion on Gitee
4. Provide code examples

---

## 📜 License

By contributing to futuapi4go, you agree that your contributions will be licensed under the Apache License 2.0. See [LICENSE](LICENSE) for details.

---

## 🙏 Thank You

Your contributions make futuapi4go better for everyone. We appreciate your time and effort!

Happy coding! 🎉
