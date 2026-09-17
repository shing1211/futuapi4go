# Governance

## Model

futuapi4go is a **single-maintainer** open-source project at present. The
maintainer has final say on design, scope, and releases. This section documents
how decisions are made so contributors know what to expect.

## Decision making

- **Small changes** (docs, tests, bug fixes): decided through pull-request review.
- **Design changes** (public API, architecture, dependencies): discussed in an
  issue and recorded in [DESIGN.md](./DESIGN.md) or a plan document before or
  alongside the code.
- **Breaking changes**: follow the release process in [AGENTS.md](./AGENTS.md).

## Roles

| Role | Responsibility |
|------|----------------|
| Maintainer | final decisions, releases, security response, CODEOWNERS |
| Translation maintainer | keeps one language's README in sync (see [TRANSLATING.md](./TRANSLATING.md)) |
| Contributor | issues, pull requests, reviews (see [CONTRIBUTING.md](./CONTRIBUTING.md)) |

Contributors who make sustained, high-quality contributions may be invited to
become reviewers, translation maintainers, or co-maintainers.

### Translations

The English [README](./README.md) is canonical. Each translated README names its
own maintainer in pull-request history. Translation changes are reviewed like any
other contribution and must pass `scripts/check_i18n.py`.

## Bus factor

The project currently has **one maintainer**. If you depend on it, consider
contributing. To reduce single-point risk, all decisions are recorded in the
repository (design docs, issues, PRs) rather than in private channels.

## Conduct

All participation is governed by the [Code of Conduct](./CODE_OF_CONDUCT.md).

## Changes to this document

Changes to governance are made by pull request and require maintainer approval.
