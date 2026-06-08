# Lesson 01 - Pointers, Receivers, and Mutation Semantics

## Goal

Understand how Go value passing, pointer receivers, and reference-containing fields affect mutation behavior.

This lesson is intentionally small. The goal is to build the mental model before moving into slices, maps, escape analysis, and concurrency.

## Files

- `theory.md` - practical foundation for the lesson
- `find-bug/task.md` - predict-output and bug-hunting prompt
- `find-bug/main.go` - runnable starter code for the bug-hunting exercise
- `mini-script/task.md` - implementation prompt
- `mini-script/main.go` - runnable starter code with TODOs

## Suggested Order

1. Read `theory.md`.
2. Open `find-bug/task.md` and answer without running the code first.
3. Run `find-bug/main.go` to verify your prediction.
4. Implement `mini-script/main.go`.
5. Compare your reasoning with the mentor review after submitting your answer.
