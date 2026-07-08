---
name: create-next-exercise
description: >-
  Generate exactly one Go interview-prep exercise aligned to the current module,
  lesson, and task rotation in this repo. Use when the user asks for the next
  task, a new exercise, or "give me something to practice" — or after an
  exercise has been reviewed and closed and it is time for the next one. Reads
  progress/progress.md to find current position and rotates task categories
  (find-bug/predict-output → mini-script → architectural/production scenario).
  Do NOT reveal the solution in the prompt; wait for the user's answer.
---

# create-next-exercise

Create exactly one practical Go exercise at a time, aligned to the current
module and the required task rotation.

## Workflow

1. Read `progress/progress.md` to identify the current module and lesson.
2. Check the previous task category (also in progress notes).
3. Select the next category in rotation:
   - find the bug / predict the output
   - write a micro-function (`mini-script`)
   - architectural debugging or production scenario
4. Briefly remind the user of 2–3 key concepts.
5. Provide exactly one task.
6. Wait for the user's answer before giving the solution.

Scaffold the exercise into the repo layout when the user wants it saved:
`exercises/module-XX-topic/lesson-YY-topic/{find-bug,mini-script}/`, with a
`task.md` and a runnable `main.go` (`go run main.go`).

## Exercise Requirements

Exercises should feel like senior platform / infrastructure interviews. Prefer:

- slice aliasing or map behavior
- allocation and escape-analysis tradeoffs
- goroutine leaks, races, deadlocks, or worker pools
- context cancellation and timeout propagation
- graceful shutdown and signal handling
- HTTP transport reuse and timeout bugs
- race detector, benchmarks, pprof, or leak debugging

## Guardrails

- Do not skip modules; move through them strictly in order (see `references/module-map.md`).
- Do not give multiple tasks at once.
- Do not reveal the solution in the prompt.
- Do not use beginner-only drills unless they expose an important Go runtime behavior.

## Reference

- `references/module-map.md` — the 7-module curriculum and per-module coverage.
