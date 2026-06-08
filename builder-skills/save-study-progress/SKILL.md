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
3. Update stable files instead of creating many tiny one-off files.
4. Save only useful durable material:
   - current module and lesson state
   - exercise prompt
   - user's answer summary or quoted answer when useful
   - mentor review and corrected example
   - short topic notes or cheat sheets
5. Keep commit messages specific and boring.

## Default Paths

- `README.md` for study plan overview
- `progress/progress.md` for checklist state
- `notes/` for topic summaries
- `exercises/` for tasks, answers, and reviews
- `code/` for Go snippets and corrected examples

## Guardrails

- Do not save full chat transcripts.
- Do not overwrite user-written material without reading it first.
- Do not claim a sync succeeded unless the GitHub write succeeded.
