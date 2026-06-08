# create-next-exercise

Use this skill when the mentor needs to generate the next Go interview preparation task.

## Purpose

Create exactly one practical exercise at a time, aligned to the current module and the required task rotation.

## Workflow

1. Identify the current module and lesson.
2. Check the previous task category.
3. Select the next category in rotation:
   - find the bug / predict the output
   - write a micro-function
   - architectural debugging or production scenario
4. Briefly remind the user of 2-3 key concepts.
5. Provide exactly one task.
6. Wait for the user's answer before giving the solution.

## Exercise Requirements

Exercises should feel like senior platform or infrastructure interviews. Prefer tasks involving:

- slice aliasing or map behavior
- allocation and escape-analysis tradeoffs
- goroutine leaks, races, deadlocks, or worker pools
- context cancellation and timeout propagation
- graceful shutdown and signal handling
- HTTP transport reuse and timeout bugs
- race detector, benchmarks, pprof, or leak debugging

## Guardrails

- Do not skip modules.
- Do not give multiple tasks at once.
- Do not reveal the solution in the prompt.
- Do not use beginner-only drills unless they expose an important Go runtime behavior.
