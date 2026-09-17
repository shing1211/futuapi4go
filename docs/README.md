# Documentation Index

> Current release: **v0.19.2** · Futu OpenD protocol **v10.10.7008** (184 protos).
> The version authority is [VERSION_MAP.md](VERSION_MAP.md).

This index lists every document and its status, so it is clear which are current
references and which are point-in-time plans.

## Start here

| Document | What it is |
|----------|------------|
| [../README.md](../README.md) | Project overview, install, quick start, package map, common APIs |
| [USAGE.md](USAGE.md) | Detailed setup, environment variables, patterns (English + 中文) |
| [CONFIGURATION.md](CONFIGURATION.md) | All client options, environment variables, and defaults |
| [TESTING.md](TESTING.md) | Test layout, the mock server, integration-test gate, benchmarks |
| [ARCHITECTURE.md](ARCHITECTURE.md) | Package layout, execution flows, concurrency, protobuf layout |
| [../DESIGN.md](../DESIGN.md) | Design decisions, wire protocol, security model, extension points |
| [ERRORS.md](ERRORS.md) | Error codes, categories, recovery hints, retry guidance |
| [VERSION_MAP.md](VERSION_MAP.md) | Futu protocol ↔ SDK tag mapping (authoritative) |

## Architecture decisions

| Document | What it is |
|----------|------------|
| [adr/README.md](adr/README.md) | Index of Architecture Decision Records (ADRs) — durable design decisions and their rationale |

## Operations & policy

| Document | What it is |
|----------|------------|
| [../CONTRIBUTING.md](../CONTRIBUTING.md) | How to contribute, checks, commit format |
| [../GOVERNANCE.md](../GOVERNANCE.md) | Decision-making and roles |
| [../SECURITY.md](../SECURITY.md) | Vulnerability reporting, supported versions, scope |
| [../SUPPORT.md](../SUPPORT.md) | Where to ask for help |
| [../TRANSLATING.md](../TRANSLATING.md) | README translation workflow |
| [../CODE_OF_CONDUCT.md](../CODE_OF_CONDUCT.md) | Community standards |
| [../DISCLAIMER.md](../DISCLAIMER.md) | Legal / no-warranty notice |
| [../CHANGELOG.md](../CHANGELOG.md) | Release history (v0.0.1 → v0.19.2) |

## Status & plans

| Document | Status | Notes |
|----------|--------|-------|
| [IMPLEMENTATION_COMPLETE.md](IMPLEMENTATION_COMPLETE.md) | COMPLETE | API-coverage summary and early-phase history |
| [IMPROVEMENT_PLAN.md](IMPROVEMENT_PLAN.md) | COMPLETE | Robustness audit (FIX-001…FIX-007, L01–L14) |
| [IMPLEMENTATION_PLAN.md](IMPLEMENTATION_PLAN.md) | Historical | Phased plan; superseded by later releases |
| [PHASE3_PROTO_SAFETY_PLAN.md](PHASE3_PROTO_SAFETY_PLAN.md) | Historical | Proto-safety refactor |
| [PHASE4_API_COVERAGE_PLAN.md](PHASE4_API_COVERAGE_PLAN.md) | Historical | API-coverage expansion (v0.14.0) |
| [PHASE5_BUGFIX_HARDENING_PLAN.md](PHASE5_BUGFIX_HARDENING_PLAN.md) | Historical | Bug fixes + hardening |
| [PHASE6_ENUM_AND_PUSH_ALIGNMENT_PLAN.md](PHASE6_ENUM_AND_PUSH_ALIGNMENT_PLAN.md) | Historical | Enum/push alignment (v0.19.0) |
| [UPGRADE_PLAN.md](UPGRADE_PLAN.md) | Historical | Protocol-upgrade template (v10.5 → v10.6) |

> **Historical** documents are point-in-time plans kept for context; their
> counts and version numbers are *not* maintained. For current numbers always
> use [VERSION_MAP.md](VERSION_MAP.md).

## Website

[docs/index.html](index.html) is the GitHub Pages landing page, served at
<https://shing1211.github.io/futuapi4go/>.
