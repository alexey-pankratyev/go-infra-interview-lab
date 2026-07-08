---
name: save-study-progress
description: >-
  Persist compact, human-readable Go study state into this repo (commit to git)
  without dumping transcripts. Use when the user asks to save/sync progress,
  record a reviewed exercise/answer/review, or save notes and corrected code
  snippets. Reads target files before writing, updates stable files instead of
  creating many tiny ones, validates every path stays inside the repo, and
  screens content for secrets before committing.
---

# save-study-progress

Persist compact, human-readable study state without turning the repository into
a transcript dump.

## When to Use

- The user asks to save current progress.
- An answer was reviewed and the exercise, answer, and review need recording.
- The user asks to save notes, code drafts, or module status.

## Workflow

1. Read existing target files before writing.
2. Normalize and validate every target path before any read, write, or delete.
3. Update stable files instead of creating many tiny one-off files.
4. Screen the material for sensitive data before saving it.
5. Save only useful durable material:
   - current module and lesson state
   - exercise prompt
   - user's answer summary or quoted answer when useful
   - mentor review and corrected example
   - short topic notes or cheat sheets
6. Keep commit messages specific and boring. Only commit/push when the user
   asks; branch first if on `main` and the user wants a PR.

## Default Paths

See `references/repo-layout.md` for the full layout. Key targets:

- `progress/progress.md` — checklist state
- `notes/module-XX-topic.md` — module-level topic summaries
- `exercises/module-XX-topic/lesson-YY-topic/` — lesson-scoped practice
- `code/` — corrected examples or larger reusable snippets

Each `main.go` starter should be runnable with `go run main.go`.

## Path Boundary

All paths must be repository-relative and stay inside the repository root.
Reject or ask for clarification when a path is absolute, parent-traversing
(`../`), or outside the study root. Do not read, write, move, or delete anything
outside `go-infra-interview-lab`.

## Sensitive Data Handling

Before saving, inspect proposed and relevant existing content for: API keys,
tokens, session cookies, bearer tokens, private/SSH keys, certificates,
kubeconfigs, cloud credentials (AWS/GCP/Azure/GitHub/Docker/CI), `.env` secrets,
database URLs, passwords, connection strings, and unneeded personal data.

If found or suspected: do not persist the secret as-is; redact when the note is
still useful; tell the user what category was detected and where; ask for
explicit confirmation before saving anything that may still expose private data.

## Deletion Policy

Never delete repository files as part of cleanup, sync, overwrite, or
reorganization unless the user explicitly asks. A valid deletion request must
name the target or a clear scope; confirm the path is inside the repo boundary
and summarize what will be removed before deleting.

## Guardrails

- Do not save full chat transcripts.
- Do not overwrite user-written material without reading it first.
- Do not claim a save/commit succeeded unless the write actually succeeded.
- Do not write outside the repository boundary.
- Do not delete files without an explicit user deletion request.
- Do not persist sensitive data silently; highlight it and notify the user.

## Reference

- `references/repo-layout.md` — stable repository layout and save rules.
