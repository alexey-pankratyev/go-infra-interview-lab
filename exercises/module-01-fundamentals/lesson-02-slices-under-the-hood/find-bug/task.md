# Find the Bug - Slice Aliasing

## Task Type

Predict the output / find the bug.

## Instructions

Read `main.go` first. Do not run it before answering.

Answer these questions:

1. What does the program print?
2. What happened to `a[3]` after the `append` call and why?
3. At what point do `a` and `b` stop sharing the same backing array?
4. What is the production risk in this pattern?

After writing your answer, run:

```bash
go run main.go
```

## Expected Reasoning Areas

- slice header: Data, Len, Cap
- subslice capacity and offset into parent array
- append in-place vs reallocation
- silent mutation of parent slice through subslice append

## Review Placeholder

Add the user's answer and mentor review here after the exercise is completed.
