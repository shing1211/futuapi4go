# Documentation Index

Complete guide to all futuapi4go documentation.

## 📚 Documentation Overview

| Document | Audience | Purpose | When to Read |
|----------|----------|---------|--------------|
| [README.md](README.md) | Everyone | Project overview, quick start, features | **Start here** |
| [API_REFERENCE.md](API_REFERENCE.md) | Developers | Complete SDK function reference | When using APIs |
| [USER_GUIDE.md](USER_GUIDE.md) | Traders | User-facing guide for trading | When getting started |
| [DEVELOPER.md](DEVELOPER.md) | Contributors | SDK development architecture | When contributing |
| [TESTING.md](TESTING.md) | QA/Developers | Test suite guide with examples | When writing tests |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Contributors | How to contribute | Before making PRs |
| [SIMULATOR.md](SIMULATOR.md) | Testers | Mock OpenD server docs | When testing without OpenD |
| [CHANGELOG.md](CHANGELOG.md) | Everyone | Version history | When upgrading |
| [PROJECT_STATUS.md](PROJECT_STATUS.md) | Maintainers | Current development status | Project management |
| [docs/RELEASE_CHECKLIST.md](docs/RELEASE_CHECKLIST.md) | Maintainers | Pre-release verification | When releasing |
| [cmd/examples/README.md](cmd/examples/README.md) | Users | Example programs guide | When learning |
| [cmd/examples/ALGO_README.md](cmd/examples/ALGO_README.md) | Traders | Algorithmic trading examples | When building strategies |
| [docs/Futu-API-Doc-hk-Proto.md](docs/Futu-API-Doc-hk-Proto.md) | Protocol devs | Official protocol specification | Deep protocol work |

---

## 🎯 Documentation by Role

### 👨‍💻 For Developers

**Getting Started:**
1. [README.md](README.md) - Quick start and installation
2. [API_REFERENCE.md](API_REFERENCE.md) - Complete API reference with examples
3. [USER_GUIDE.md](USER_GUIDE.md) - Detailed usage guide

**Working with Code:**
1. [DEVELOPER.md](DEVELOPER.md) - Architecture and development setup
2. [CONTRIBUTING.md](CONTRIBUTING.md) - Contribution guidelines
3. [TESTING.md](TESTING.md) - Testing strategies

**Examples:**
1. [cmd/examples/README.md](cmd/examples/README.md) - All examples
2. [cmd/examples/ALGO_README.md](cmd/examples/ALGO_README.md) - Algorithmic strategies

### 📊 For Traders

**Using the SDK:**
1. [README.md](README.md) - Overview and quick start
2. [USER_GUIDE.md](USER_GUIDE.md) - Complete trading guide
3. [API_REFERENCE.md](API_REFERENCE.md) - API reference

**Strategies:**
1. [cmd/examples/ALGO_README.md](cmd/examples/ALGO_README.md) - Trading algorithms
2. [TESTING.md](TESTING.md) - Testing with HSI data

### 🧪 For QA/Testers

**Testing:**
1. [TESTING.md](TESTING.md) - Complete testing guide
2. [SIMULATOR.md](SIMULATOR.md) - Mock server usage
3. [test/fixtures/](test/fixtures/) - Test data fixtures

### 🏗️ For Maintainers

**Project Management:**
1. [PROJECT_STATUS.md](PROJECT_STATUS.md) - Current status
2. [docs/RELEASE_CHECKLIST.md](docs/RELEASE_CHECKLIST.md) - Release process
3. [CHANGELOG.md](CHANGELOG.md) - Version history

---

## 📖 Documentation by Topic

### Installation & Setup

| Document | Sections |
|----------|----------|
| [README.md](README.md) | Installation, Quick Start, Prerequisites |
| [USER_GUIDE.md](USER_GUIDE.md) | Connection Management, Basic Usage |
| [DEVELOPER.md](DEVELOPER.md) | Development Environment, Build Process |

### Market Data

| Document | Sections |
|----------|----------|
| [API_REFERENCE.md](API_REFERENCE.md) | Market Data (qot) - All 37 APIs |
| [USER_GUIDE.md](USER_GUIDE.md) | Market Data Queries, Subscriptions |
| [cmd/examples/README.md](cmd/examples/README.md) | Market Data Examples (11 examples) |

### Trading

| Document | Sections |
|----------|----------|
| [API_REFERENCE.md](API_REFERENCE.md) | Trading (trd) - All 16 APIs |
| [USER_GUIDE.md](USER_GUIDE.md) | Trading Operations, Order Management |
| [cmd/examples/README.md](cmd/examples/README.md) | Trading Examples (7 examples) |

### Testing

| Document | Sections |
|----------|----------|
| [TESTING.md](TESTING.md) | Complete testing guide (all sections) |
| [SIMULATOR.md](SIMULATOR.md) | Mock server setup and usage |
| [test/fixtures/hsi_fixtures.go](test/fixtures/hsi_fixtures.go) | HSI test data reference |

### Architecture

| Document | Sections |
|----------|----------|
| [DEVELOPER.md](DEVELOPER.md) | Architecture, Design Patterns, Code Structure |
| [README.md](README.md) | Architecture Diagram |
| [docs/Futu-API-Doc-hk-Proto.md](docs/Futu-API-Doc-hk-Proto.md) | Protocol Specification |

### Protocol Reference

| Document | Description |
|----------|-------------|
| [docs/Futu-API-Doc-hk-Proto.md](docs/Futu-API-Doc-hk-Proto.md) | Official Futu protocol (575KB, 13K+ lines) |
| [api/proto/](api/proto/) | Protocol buffer definitions (74 files) |
| [pkg/pb/](pkg/pb/) | Generated Go code (74 packages) |

---

## 🚀 Quick Reference

### Most Common Tasks

**Get real-time quote:**
- See: [README.md - Quick Start](README.md#-quick-start)
- See: [API_REFERENCE.md - GetBasicQot](API_REFERENCE.md#getbasicqot)
- Example: [cmd/examples/qot_get_basic_qot/](cmd/examples/qot_get_basic_qot/)

**Place trading order:**
- See: [USER_GUIDE.md - Place Order](USER_GUIDE.md#place-order)
- See: [API_REFERENCE.md - PlaceOrder](API_REFERENCE.md#placeorder)
- Example: [cmd/examples/trd_place_order/](cmd/examples/trd_place_order/)

**Subscribe to push:**
- See: [TESTING.md - Integration Tests](TESTING.md#-integration-tests)
- See: [API_REFERENCE.md - Subscribe](API_REFERENCE.md#subscribe)
- Example: [cmd/examples/04_push_subscriptions/](cmd/examples/04_push_subscriptions/)

**Run tests:**
- See: [TESTING.md - Quick Start](TESTING.md#-quick-start)
- See: [CONTRIBUTING.md - Testing Requirements](CONTRIBUTING.md#testing-requirements)

**Contribute code:**
- See: [CONTRIBUTING.md](CONTRIBUTING.md) (complete guide)
- See: [DEVELOPER.md](DEVELOPER.md) (architecture)

---

## 📊 Documentation Statistics

| Metric | Count |
|--------|-------|
| **Total Documents** | 13 |
| **Total Lines** | ~6,500 |
| **Code Examples** | 100+ |
| **APIs Documented** | 71 |
| **Test Examples** | 46 |
| **Languages** | English (primary), Chinese (protocol reference) |

---

## 🔄 Documentation Maintenance

### When Adding Features

1. Update [README.md](README.md) features section
2. Add to [API_REFERENCE.md](API_REFERENCE.md)
3. Update [USER_GUIDE.md](USER_GUIDE.md) with examples
4. Add tests to [TESTING.md](TESTING.md) scope
5. Update [CHANGELOG.md](CHANGELOG.md)
6. Add/update examples in cmd/examples/

### Before Releases

1. Review [docs/RELEASE_CHECKLIST.md](docs/RELEASE_CHECKLIST.md)
2. Update version in [README.md](README.md) badges
3. Update [CHANGELOG.md](CHANGELOG.md) with release date
4. Update [PROJECT_STATUS.md](PROJECT_STATUS.md)
5. Verify all links work
6. Review documentation for accuracy

### Regular Maintenance

- **Monthly**: Verify all links work
- **Per release**: Update version numbers
- **Per feature**: Add documentation
- **As needed**: Improve examples and clarity

---

## 📝 Documentation Standards

### File Naming

- **ALL_CAPS** for root documents (README.md, LICENSE, etc.)
- **snake_case** for subdirectory documents
- **ALGO_README.md** for specialized example docs

### Content Guidelines

- **Code examples**: Must compile and run
- **API references**: Include request/response examples
- **Guides**: Step-by-step with expected output
- **Architecture**: Diagrams where helpful

### Cross-Referencing

- Link to other documents liberally
- Use relative paths (not absolute)
- Include section anchors when helpful

---

## ❓ Need Help?

- **Can't find what you need?** Open an issue: [Gitee Issues](https://gitee.com/shing1211/futuapi4go/issues)
- **Documentation improvement?** Submit a PR following [CONTRIBUTING.md](CONTRIBUTING.md)
- **Questions?** Start a discussion: [Gitee Discussions](https://gitee.com/shing1211/futuapi4go/discussions)

---

**Last Updated**: 2026-04-11  
**Total Documents**: 13  
**Status**: ✅ Complete and current
