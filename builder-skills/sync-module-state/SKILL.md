# sync-module-state

Use this skill when the mentor needs to update or read compact module progress state.

## Purpose

Keep the learning flow resumable across sessions without storing full conversations.

## When to Use

- The user completes a lesson or exercise.
- The mentor needs to resume the current module.
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

Prefer markdown for user-visible progress and YAML for compact machine-readable defaults.

## Guardrails

- Do not store full chat logs.
- Do not store secrets.
- Do not mark a module complete until the current exercise is reviewed and closed.
- Do not jump ahead in module order.
