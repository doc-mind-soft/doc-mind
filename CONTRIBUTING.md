# Contributing to DocMind

## Development principle

DocMind is developed through small, independently verifiable iterations.

Each iteration must have a narrow goal, a defined validation method, and a reviewable result. Unrelated changes must not be combined merely because they are convenient to implement at the same time.

## Branches

`main` is the only long-lived branch and must remain in a usable state.

The project does not use a `develop` branch.

Every iteration starts from an up-to-date `main`:

```bash
git switch main
git pull --ff-only origin main
git status --short
git switch -c iteration/<number>-<short-description>
```

Branch names use lowercase kebab-case. For example:

```text
iteration/00-07-document-git-workflow
```

One branch represents one iteration. A merged branch must not be reused for later work.

Direct pushes to `main` are not part of the normal workflow. The initial repository commit was the only bootstrap exception.

## Commits

A commit must represent one logical change and must not contain unrelated files.

Commit subjects follow this form:

```text
<type>(<scope>): <imperative description>
```

Common types include:

- `feat` for user-visible functionality;
- `fix` for defect corrections;
- `docs` for documentation;
- `test` for tests;
- `refactor` for behavior-preserving code changes;
- `chore` for repository and tooling maintenance.

The subject uses lowercase after the colon, does not end with a period, and should explain what the commit does.

A commit body should explain relevant context or intent when the subject alone is insufficient.

Before committing, inspect exactly what is staged:

```bash
git diff --cached --check
git diff --cached --name-status
git diff --cached
```

## Pull requests

Every normal iteration is delivered through a Pull Request targeting `main`.

Before creating a Pull Request:

- confirm that the branch contains only the intended changes;
- review the comparison against `main`;
- run the checks relevant to the iteration;
- confirm that the working tree contains no accidental changes;
- confirm that no credentials or local environment files are included.

The Pull Request title should match the intended squash commit subject.

The description must contain:

- `Summary` with the purpose of the change;
- `Verification` with the checks that were actually performed.

Checks that were not performed must not be reported as successful.

If the branch changes after review, the updated diff and relevant checks must be reviewed again.

## Merge and cleanup

Pull Requests are merged using `Squash and merge`.

Squashing preserves one logical commit per accepted iteration even when review corrections required additional branch commits.

After the merge:

1. delete the remote iteration branch;
2. switch the local repository to `main`;
3. update `main` using fast-forward only;
4. prune removed remote branches;
5. delete the local iteration branch;
6. verify that the working tree is clean.

```bash
git switch main
git pull --ff-only origin main
git fetch --prune origin
git branch -D iteration/<number>-<short-description>
git status --short
git branch -vv
```

The next iteration starts only after the previous Pull Request has been merged and repository cleanup has been verified.
