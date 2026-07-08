---
name: go-repo-onboarding
description: >-
  Confirm the minimum repository defaults (repo, branch, root path) before
  writing or committing Go study materials. Use when the user asks to save,
  sync, export, or update study materials and the target repo/branch/path is
  missing, stale, or explicitly changed — or before the first write of a
  session. Verifies the repo and branch are reachable, records compact defaults,
  then resumes the original request. Never invents repo names or writes before
  access is verified.
---

# go-repo-onboarding

Collect the minimum durable defaults required to write progress, exercises,
notes, and draft code into the study repository, then get out of the way.

## When to Use

- The user asks to save, sync, export, or update study materials.
- Repository defaults are missing, stale, or explicitly changed by the user.
- You need to confirm branch or root path before writing files.

## Workflow

1. Check existing durable defaults before asking the user anything (repo root,
   `go-study-defaults.yaml`, git remote/branch).
2. If defaults are missing, ask only for the missing values:
   - repository full name, e.g. `owner/repo`
   - branch, default `main`
   - optional root path, default repository root
3. Verify the repository and branch are reachable (e.g. `git remote -v`,
   `git branch --show-current`).
4. Store compact defaults for future sessions.
5. Resume the user's original request.

## Output

Keep the confirmation short. Do not turn onboarding into a lesson.

## Contract

See `references/onboarding-contract.yaml` for required/optional defaults,
validation rules, and what must never be stored.

## Guardrails

- Do not invent repository names, owners, branches, or paths.
- Do not write or commit until repository access has been verified.
- Do not store full transcripts, access tokens, or long-form answers in durable state.

## Reference

- `references/onboarding-contract.yaml` — defaults, validation, and memory policy.
