---
name: sync-module-state
description: >-
  Read or update compact, resumable Go module progress state so the learning
  flow survives across sessions without storing full conversations. Use when the
  user completes a lesson/exercise, asks to resume "where were we", or asks to
  save/sync progress, and when weak spots or study preferences should persist.
  Keeps progress/progress.md (human-readable) plus optional
  go-study-defaults.yaml (machine-readable). Save progress state before longer
  notes when both are requested.
---

# sync-module-state

Keep the learning flow resumable across sessions without storing full
conversations.

## When to Use

- The user completes a lesson or exercise.
- You need to resume the current module.
- The user asks to save or sync progress.
- Study preferences or weak spots should persist for future sessions.

## Workflow

1. Read existing progress state first.
2. Update only durable state:
   - current module and lesson
   - completed modules or exercises
   - last task category
   - known weak spots
   - repository defaults when relevant
3. Keep state compact and easy to inspect.
4. Save progress state before saving longer notes when both are requested.

## State Shape

Prefer markdown for user-visible progress and YAML for compact machine-readable
defaults. See `references/state-files.md` for the exact file shapes.

## Guardrails

- Do not store full chat logs.
- Do not store secrets.
- Do not mark a module complete until the current exercise is reviewed and closed.
- Do not jump ahead in module order.

## Reference

- `references/state-files.md` — file names, fields, and update rules.
