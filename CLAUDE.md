# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Repository Is

A structured interview-prep laboratory for Go/Senior DevOps roles. It is not a deployable software project — there is no `go.mod`, no build system, and no test suite. Individual Go snippets in `exercises/` are run with `go run main.go` directly.

## Running Exercises

```bash
# Run any exercise snippet directly
go run exercises/module-01-fundamentals/lesson-01-pointers-receivers-mutation/find-bug/main.go
go run exercises/module-01-fundamentals/lesson-01-pointers-receivers-mutation/mini-script/main.go
```

## Repository Structure

```
agent/               # ChatGPT Agent configuration (mentor persona, prompts)
builder-skills/      # Reusable skill packages for ChatGPT Agent Builder
  go-github-onboarding/
  save-study-progress/
  review-go-answer/
  create-next-exercise/
  sync-module-state/
exercises/           # Practical tasks organized by module and lesson
  module-NN-topic/
    lesson-NN-topic/
      theory.md        # Short conceptual foundation
      find-bug/        # Predict output or identify the bug
      mini-script/     # Small implementation task
notes/               # Cheat sheets per module
progress/            # Study tracking
  progress.md          # Current status and checklist
  study-plan.md        # Full 7-module curriculum
code/                # Draft Go snippets and corrected examples
```

## Exercise Conventions

- Each lesson directory contains `theory.md`, a `find-bug/` subdirectory, and a `mini-script/` subdirectory.
- `find-bug/` exercises: predict output or locate the bug — answer goes in the same directory as `answer.md` (or similar).
- `mini-script/` exercises: implement a small function — solution in `main.go`.
- Mentor reviews live alongside exercises in `review.md` files.

## Curriculum (7 Modules)

1. Go fundamentals: pointers, receivers, mutation semantics
2. Arrays, slices, and maps under the hood
3. Memory management and allocations
4. Concurrency and Go runtime internals (goroutines, GMP scheduler, channels, sync)
5. Systems programming and networking (os/exec, signals, net/http, gRPC)
6. Testing, race detection, benchmarking, profiling
7. Final interview drill

## Agent / Mentoring Framework

`agent/system-instructions.md` defines the ChatGPT mentor persona — task formats, GitHub save workflow, and the required topic coverage checklist. When modifying skill files in `builder-skills/`, keep the interfaces consistent with what `system-instructions.md` describes (skill names, input/output contracts).

## Progress Tracking

Current module/lesson status is in `progress/progress.md`. Update this file after completing exercises.
