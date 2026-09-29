# Spec Delta

## MODIFIED Requirements

### Requirement: Required status checks gate merges into main

`main` MUST be configured with the required status checks `build & test`, `docs guards`, `scan`, and `analyze`. Merges into `main` MUST be refused until all four pass. These checks MUST also be enforced for administrators, so that no actor bypasses the gate. A push that skips the gate MUST be treated as a defect, not a shortcut.

#### Scenario: A change is delivered to main

- **WHEN** a change is delivered to `main`
- **THEN** it arrives through a pull request rather than a direct push
- **AND** the pull request is merged only after all four required status checks report success

#### Scenario: A direct push is attempted

- **WHEN** a commit is pushed directly to `main` instead of merged through a pull request
- **THEN** the push is rejected with `GH006: Protected branch update failed for refs/heads/main`
- **AND** this rejection applies to administrators as well as everyone else
- **AND** the four required status checks gate the only supported path, which is a pull request merge

#### Scenario: Enforcement is weakened

- **WHEN** `enforce_admins` is set to `false` on `main`
- **THEN** administrators can once again push to `main` without passing the required status checks
- **AND** this is a regression against this requirement, because the checks no longer constrain every actor
