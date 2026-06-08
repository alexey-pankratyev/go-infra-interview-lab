# save-study-progress

Use this skill when the user asks to save or sync Go interview study progress to GitHub.

## Purpose

Persist compact, human-readable study state without turning the repository into a transcript dump.

## When to Use

- The user asks to save current progress.
- The mentor has reviewed an answer and needs to record the exercise, answer, and review.
- The user asks to sync notes, code drafts, or module status.

## Workflow

1. Confirm repository defaults are available.
2. Read existing target files before writing.
3. Normalize and validate every target path before any read, write, or delete.
4. Update stable files instead of creating many tiny one-off files.
5. Check the material for sensitive data before saving it.
6. Save only useful durable material:
   - current module and lesson state
   - exercise prompt
   - user's answer summary or quoted answer when useful
   - mentor review and corrected example
   - short topic notes or cheat sheets
7. Keep commit messages specific and boring.

## Default Paths

- `README.md` for study plan overview
- `progress/progress.md` for checklist state
- `notes/` for topic summaries
- `exercises/` for tasks, answers, and reviews
- `code/` for Go snippets and corrected examples

## Path Boundary

The default repository is `alexey-pankratyev/go-infra-interview-lab` on branch `main`.

All repository paths must be repository-relative and must stay inside the repository root. Reject or ask for clarification when a path is:

- absolute, for example `/tmp/file` or `/home/user/file`
- parent-traversing, for example `../notes.md`
- ambiguous about the target repository
- outside the configured study root path

Do not read, write, move, or delete anything above or outside `go-infra-interview-lab` unless the user explicitly changes the repository defaults and that new repository is verified.

## Sensitive Data Handling

Before saving or updating repository files, inspect the proposed content and any relevant existing content for sensitive data indicators, including:

- API keys, access tokens, session cookies, and bearer tokens
- private keys, SSH keys, certificates, and kubeconfigs
- cloud credentials for AWS, GCP, Azure, GitHub, Docker, or CI systems
- `.env` style secrets, database URLs, passwords, and connection strings
- personal data that is not needed for study progress

If sensitive data is found or suspected:

1. Do not persist the secret value as-is.
2. Redact the value when a saved note is still useful.
3. Tell the user what category of sensitive data was detected and where it appeared.
4. Ask for explicit confirmation before saving anything that may still expose private data.

## Deletion Policy

Never delete repository files as part of cleanup, sync, overwrite, or reorganization unless the user explicitly asks for deletion.

A valid deletion request must name the target file or clear scope. Before deleting, confirm the path is inside the repository boundary and summarize what will be removed.

## Guardrails

- Do not save full chat transcripts.
- Do not overwrite user-written material without reading it first.
- Do not claim a sync succeeded unless the GitHub write succeeded.
- Do not write outside the verified study repository boundary.
- Do not delete files without an explicit user deletion request.
- Do not persist sensitive data silently; highlight it and notify the user.
