# futuapi4go — Agent Guide

Go SDK for the Futu OpenD / OpenAPI protocol. Pure library (no `main`
package), module `github.com/shing1211/futuapi4go`, Go 1.26.6+, Apache-2.0.
Latest release: **v0.20.0** (Futu protocol v10.10.7008, 184 protos).

**Read [AGENTS.md](./AGENTS.md) first** — it is the primary agent guide
(architecture, CI gates, release process, troubleshooting).

Hard rules:
- `context.Context` is the first parameter of every public API.
- Do not auto-retry order-mutation calls.
- Run `gitnexus_impact` before editing a symbol and `gitnexus_detect_changes`
  before committing.

Commands: `make check` (gofmt-fix + vet + build), `make test` (race),
`make fmt` (format check), `make docs-check` (translations).

---

<!-- gitnexus:start -->
# GitNexus — Code Intelligence

This project is indexed by GitNexus as **futuapi4go** (20427 symbols, 51795 relationships, 186 execution flows). Use the GitNexus MCP tools to understand code, assess impact, and navigate safely.

> If any GitNexus tool warns the index is stale, run `npx gitnexus analyze` in terminal first.

## Always Do

- **MUST run impact analysis before editing any symbol.** Before modifying a function, class, or method, run `gitnexus_impact({target: "symbolName", direction: "upstream"})` and report the blast radius (direct callers, affected processes, risk level) to the user.
- **MUST run `gitnexus_detect_changes()` before committing** to verify your changes only affect expected symbols and execution flows.
- **MUST warn the user** if impact analysis returns HIGH or CRITICAL risk before proceeding with edits.
- When exploring unfamiliar code, use `gitnexus_query({query: "concept"})` to find execution flows instead of grepping. It returns process-grouped results ranked by relevance.
- When you need full context on a specific symbol — callers, callees, which execution flows it participates in — use `gitnexus_context({name: "symbolName"})`.

## Never Do

- NEVER edit a function, class, or method without first running `gitnexus_impact` on it.
- NEVER ignore HIGH or CRITICAL risk warnings from impact analysis.
- NEVER rename symbols with find-and-replace — use `gitnexus_rename` which understands the call graph.
- NEVER commit changes without running `gitnexus_detect_changes()` to check affected scope.

## Resources

| Resource | Use for |
|----------|---------|
| `gitnexus://repo/futuapi4go/context` | Codebase overview, check index freshness |
| `gitnexus://repo/futuapi4go/clusters` | All functional areas |
| `gitnexus://repo/futuapi4go/processes` | All execution flows |
| `gitnexus://repo/futuapi4go/process/{name}` | Step-by-step execution trace |

## CLI

| Task | Read this skill file |
|------|---------------------|
| Understand architecture / "How does X work?" | `.claude/skills/gitnexus/gitnexus-exploring/SKILL.md` |
| Blast radius / "What breaks if I change X?" | `.claude/skills/gitnexus/gitnexus-impact-analysis/SKILL.md` |
| Trace bugs / "Why is X failing?" | `.claude/skills/gitnexus/gitnexus-debugging/SKILL.md` |
| Rename / extract / split / refactor | `.claude/skills/gitnexus/gitnexus-refactoring/SKILL.md` |
| Tools, resources, schema reference | `.claude/skills/gitnexus/gitnexus-guide/SKILL.md` |
| Index, status, clean, wiki CLI commands | `.claude/skills/gitnexus/gitnexus-cli/SKILL.md` |

<!-- gitnexus:end -->
