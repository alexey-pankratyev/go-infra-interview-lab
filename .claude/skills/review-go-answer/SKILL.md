---
name: review-go-answer
description: >-
  Review the user's answer to a Go interview exercise like a senior
  platform/backend interviewer — precise, constructive, focused on production
  behavior. Use after the user submits an answer, prediction, or solution to a
  find-bug or mini-script exercise, or asks "is this right?" / "review my
  answer". Prioritizes hidden allocations, aliasing, goroutine leaks, races,
  scheduler/backpressure, and cancellation/shutdown gaps. Do NOT reveal the
  solution before the user has answered.
---

# review-go-answer

Review the answer like a senior platform/backend interview code review:
precise, constructive, and focused on production behavior.

## Workflow

1. Restate the exercise outcome briefly.
2. Identify what is correct.
3. Call out weak spots and missing edge cases.
4. Explain the key runtime, memory, concurrency, or architecture nuance.
5. Provide a corrected solution or mental model only after reviewing the user's attempt.
6. Decide whether the current question is closed or needs one follow-up.

When useful, save the review alongside the exercise as `review.md` in the
lesson directory, then update `progress/progress.md` (see the
`sync-module-state` skill).

## Review Focus

Prioritize issues that matter in senior infrastructure interviews:

- hidden allocations and escape behavior
- aliasing and mutation surprises
- goroutine leaks and deadlocks
- race conditions and unsafe shared state
- scheduler, blocking, and backpressure behavior
- timeout, cancellation, and graceful shutdown gaps
- testing, benchmarking, and profiling blind spots

Use `references/review-rubric.md` to keep feedback high-signal, and finish with
an interview-signal classification (Strong / Acceptable / Risky / Incorrect).

## Tone

Be concise, technical, and candid. Praise real correctness, but do not soften
important production risks.

## Guardrails

- Do not reveal the solution before the user answers.
- Do not move to a new topic until the current task is properly closed.
- Do not dump unrelated theory.

## Reference

- `references/review-rubric.md` — the six-part review rubric and signal classes.
