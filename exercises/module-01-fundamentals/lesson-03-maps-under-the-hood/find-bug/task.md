# Find the Bug - Concurrent Map Access

## Task Type

Predict the output / find the bug.

## Instructions

Read `main.go` first. Do not run it before answering.

Answer these questions:

1. Is there a bug in this program? If yes, what is it?
2. What will happen when you run it? Will it always behave the same way?
3. How would you detect this bug with Go tooling?
4. What is the minimal fix?

After writing your answer, run:

```bash
go run -race main.go
```

Note the `-race` flag — run it with the race detector enabled.

## Expected Reasoning Areas

- concurrent map write semantics
- runtime panic vs silent data corruption
- race detector usage
- Mutex vs RWMutex choice

## Review Placeholder

Add the user's answer and mentor review here after the exercise is completed.
