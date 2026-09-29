# Tasks

## 1. Repair the v0.20.0 formatting regression

- [x] 1.1 Confirm the defect on the unmodified tree, so the fix is provably a repair and not a no-op: `gofmt -l .` lists `internal/client/client.go`, `internal/client/client_test.go`, `pkg/constant/errors_test.go`
- [x] 1.2 Reformat the three files with `gofmt -w internal/client/client.go internal/client/client_test.go pkg/constant/errors_test.go`
- [x] 1.3 Review `git diff` to confirm the change is whitespace-only within existing struct definitions — no identifier, type, or exported-name change — then commit as `fix(format): restore gofmt compliance broken in v0.20.0`
- [x] 1.4 Verify the Format check now passes: `gofmt -l .` produces no output

## 2. Update the four indirect dependencies

- [x] 2.1 Apply the update with `go get -u ./...`, expecting exactly four indirect moves: `prometheus/client_model` v0.6.2→v0.6.3, `prometheus/common` v0.70.1→v0.72.0, `prometheus/procfs` v0.21.1→v0.22.0, `golang.org/x/sys` v0.47.0→v0.48.0
- [x] 2.2 Confirm `go.mod` shows only those four `// indirect` lines changed and `go.sum` grew from 47 to 55 lines with no removals: `git diff --stat go.mod go.sum`
- [x] 2.3 Confirm no direct require moved, since all six were already current: `git diff go.mod` shows no change to the first require block
- [x] 2.4 Tidy the module graph with `go mod tidy`
- [x] 2.5 Commit as `chore(deps): update 4 indirect Prometheus and x/sys modules`

## 3. Prepare the v0.21.0 release

- [x] 3.1 Convert the `[Unreleased]` section of `CHANGELOG.md` to `## [0.21.0] - 2026-09-29`, using the unprefixed heading style of the 39 existing unprefixed entries
- [x] 3.2 Add a `Fixed` entry naming `d435f9b` as the source of the formatting regression, so the red v0.20.0 build is traceable rather than silently erased
- [x] 3.3 Add a `Changed` entry listing the four dependency bumps and stating that no protos and no public API in `client/` or `pkg/` are affected
- [x] 3.4 Update the version badge and the `go get` line in `README.md` only, matching the file scope of the v0.20.0 release commit `d435f9b`
- [x] 3.5 Leave the other 18 files carrying v0.20.0 untouched — they are swept post-release, per task 5.1
- [x] 3.6 Confirm the release workflow can extract notes for this version: `awk '/^## \[0\.21\.0\]/{f=1;next} f&&/^## \[/{exit} f' CHANGELOG.md` produces non-empty output

## 4. Verify and deliver through a pull request

- [x] 4.1 Run the full local gate and require all six to pass: `gofmt -l .` empty, `go build ./...`, `go vet ./...`, `go test -race -count=1 ./...` at 29/29 packages, `govulncheck ./...` reporting no vulnerabilities, and `make docs-check` clean
- [x] 4.2 Push the branch: `git push -u origin fix/ci-and-deps-v0.21.0`
- [x] 4.3 Open the pull request with `gh pr create`, summarizing the formatting repair, the four dependency bumps, and the verified gate results
- [x] 4.4 Confirm all four required status checks report success — `build & test`, `docs guards`, `scan`, `analyze` — before merging, since this is the first delivery where the gate actually blocks
- [x] 4.5 Merge the pull request into `main` and confirm the working tree is clean: `git status --porcelain` is empty

## 5. Release and follow up

- [x] 5.1 After the merge, sweep the remaining version references from v0.20.0 to v0.21.0 in a separate commit touching the other 18 files plus `docs/index.html`, following the precedent of `e8b91e7` and `55a4c35`
- [x] 5.2 Create the annotated tag: `git tag -a v0.21.0 -m "v0.21.0: gofmt compliance repair, 4 indirect dependency updates"`
- [x] 5.3 Push the tag to both remotes: `git push origin v0.21.0 && git push gitee v0.21.0`
- [x] 5.4 Confirm `release.yml` published a GitHub release whose body came from the CHANGELOG section rather than generated notes: `gh release view v0.21.0`
- [x] 5.5 Verify all three refs agree, rather than trusting push output: `git rev-parse HEAD`, `git ls-remote origin refs/heads/main`, and `git ls-remote gitee refs/heads/main` all report the merge commit
- [x] 5.6 Set `enforce_admins: true` on `main`, so direct pushes can no longer bypass the required status checks — applied after the release, on approval. Verified by an actual direct push, which was rejected with `GH006: Protected branch update failed`; `enforce_admins` was the only key changed, all other protection settings left identical
