# Tasks

## 1. Correct the build-compliance spec

- [x] 1.1 Confirm the delta binds to the existing requirement rather than creating a duplicate: the delta heading `### Requirement: Required status checks gate merges into main` matches `openspec/specs/build-compliance/spec.md:24` byte-for-byte
- [x] 1.2 Add the enforcement clause to the requirement text, stating that the checks MUST be enforced for administrators as well so no actor bypasses the gate
- [x] 1.3 Correct the *A direct push is attempted* scenario to assert rejection with `GH006: Protected branch update failed for refs/heads/main`, applying to administrators, leaving a pull request merge as the only supported path
- [x] 1.4 Add an *Enforcement is weakened* scenario naming `enforce_admins`, the resulting bypass, and that it is a regression against the requirement
- [x] 1.5 Leave the other three requirements and the gofmt scenario untouched, since the gofmt scenario's claim that an unformatted commit cannot reach `main` became true when enforcement was enabled: `git diff` shows changes only within the one modified requirement

## 2. Rewrite the AGENTS.md release process

- [x] 2.1 Replace the `git push origin main && git push gitee main` step, which now fails with `GH006`, with a branch-then-pull-request flow: `git switch -c release/vx.y.z`, commit, `git push -u origin release/vx.y.z`
- [x] 2.2 Add the PR gate step with the commands verified present in the installed `gh`: `gh pr create`, `gh pr checks <n> --watch`, `gh pr merge <n> --merge --delete-branch`, then `git switch main && git pull --ff-only`
- [x] 2.3 Fix the ordering defect: the previous block pushed `origin main` before tagging, and synced the gitee mirror alongside the tag; move the mirror sync to after the merge, since `main` now advances only through a merged PR
- [x] 2.4 State that the gitee mirror is unprotected and accepts direct pushes, so it is synced deliberately after the merge rather than incidentally
- [x] 2.5 Keep the existing `release.yml` note and the `gh release create` fallback intact, since both remain accurate
- [x] 2.6 Add one line to the Session Checklist, `Delivered via PR; all four required status checks green before merge`, and leave the Documentation checklist at line 125 unchanged to avoid duplicating the idea
- [x] 2.7 Update the footer `*Last updated: 2026-09-17*` to the date of this edit, since the file's own metadata would otherwise be wrong the moment it is edited

## 3. Verify, deliver, and sync

- [x] 3.1 Confirm the local gates still pass: `gofmt -l .` produces no output and `make docs-check` reports 6 languages consistent, since `docs guards` runs only the i18n check and `AGENTS.md` is not a translated README
- [x] 3.2 Confirm the change validates as strict: `openspec validate align-docs-with-enforced-gate --strict`
- [x] 3.3 Push the branch and open a pull request, which is now the only supported delivery path for any change to `main`
- [x] 3.4 Wait for all four required status checks — `build & test`, `docs guards`, `scan`, `analyze` — to report success before merging
- [x] 3.5 After the merge, sync the unprotected mirror with `git push gitee main`, then verify local, `origin/main`, and `gitee/main` all agree using `git rev-parse HEAD` and `git ls-remote` rather than trusting push output
- [x] 3.6 Archive the change so the MODIFIED delta syncs into `openspec/specs/build-compliance/spec.md`, then confirm with `openspec validate --specs --strict` that the main spec is still valid
- [x] 3.7 Confirm the mirror's two historical `git push origin main` references in `docs/PHASE4_API_COVERAGE_PLAN.md` and `docs/PHASE5_BUGFIX_HARDENING_PLAN.md` were left untouched, since both are COMPLETE point-in-time records: `grep -rn "git push origin main" --include=*.md .` returns only those two
