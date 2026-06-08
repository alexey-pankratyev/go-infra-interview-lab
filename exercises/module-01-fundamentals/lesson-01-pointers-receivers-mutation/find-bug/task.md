# Find the Bug - Config Mutation

## Task Type

Predict the output / find the bug.

## Instructions

Read `main.go` first. Do not run it before answering.

Answer these questions:

1. What does the program print?
2. Why does `Retries` behave differently from `Labels`?
3. What is the production risk in this API design?
4. What would you change if this were real configuration code?

After writing your answer, run:

```bash
go run main.go
```

## Expected Reasoning Areas

- value receiver semantics
- map reference behavior inside copied structs
- mutation surprises in APIs
- production config design and concurrency risk

## Review Placeholder

Add the user's answer and mentor review here after the exercise is completed.
