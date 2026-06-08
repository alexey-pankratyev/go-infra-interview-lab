# go-github-onboarding

Use this skill when the Go interview mentor needs GitHub repository defaults before saving study materials.

## Purpose

Collect the minimum durable defaults required to write progress, exercises, notes, and draft code into the user's study repository.

## When to Use

- The user asks to save, sync, export, or update study materials in GitHub.
- Repository defaults are missing, stale, or explicitly changed by the user.
- The mentor needs to confirm branch or root path before writing files.

## Workflow

1. Check existing durable defaults before asking the user anything.
2. If defaults are missing, ask only for the missing values:
   - repository full name, for example `owner/repo`
   - branch, default `main`
   - optional root path, default repository root
3. Verify that the repository and branch are reachable.
4. Store compact defaults for future sessions.
5. Resume the user's original request.

## Output

Keep the confirmation short. Do not turn onboarding into a lesson.

## Guardrails

- Do not invent repository names, owners, branches, or paths.
- Do not write to GitHub until repository access has been verified.
- Do not store full transcripts or long explanations in durable state.
