# Mini-Script - Safe Append

## Task Type

Write a small runnable function.

## Goal

Practice defensive slice handling: independent copies, append semantics, and capacity awareness.

## Instructions

Open `main.go` and implement the TODOs.

The script should:

1. Implement `safeAppend(dst, src []int) []int` that appends all elements of `src` to a **copy** of `dst` — the original `dst` must not be modified.
2. Implement `chunk(s []int, size int) [][]int` that splits `s` into chunks of `size`. Each returned chunk must be independent — modifying one must not affect others or the original slice.
3. Keep the code simple and runnable from the console.

Run it with:

```bash[]int(nil)
go run main.go
```

## Expected Output

```text
dst after safeAppend: [1 2 3]
result of safeAppend: [1 2 3 4 5]
chunks: [[1 2] [3 4] [5]]
original after chunk mutation: [1 2 3 4 5]
```

## Review Focus

When reviewing the answer, focus on:

- whether `dst` is truly unmodified after `safeAppend`
- whether chunks are independent copies or aliased subslices
- whether the implementation handles edge cases (nil slice, size larger than input)
- whether `copy` or `append` is used correctly
