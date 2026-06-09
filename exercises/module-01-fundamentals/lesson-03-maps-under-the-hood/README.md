# Lesson 03 - Maps Under the Hood

## Goal

Understand how Go maps are implemented internally, why they are not thread-safe, what happens during concurrent access, and how to handle maps safely in production code.

## Files

- `theory.md` — practical foundation for the lesson
- `find-bug/task.md` — predict-output and bug-hunting prompt
- `find-bug/main.go` — runnable starter code for the bug-hunting exercise
- `mini-script/task.md` — implementation prompt
- `mini-script/main.go` — runnable starter code with TODOs

## Suggested Order

1. Read `theory.md`.
2. Open `find-bug/task.md` and answer without running the code first.
3. Run `find-bug/main.go` to verify your prediction.
4. Implement `mini-script/main.go`.
5. Compare your reasoning with the mentor review after submitting your answer.
