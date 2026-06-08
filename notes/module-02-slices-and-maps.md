# Module 02 — Slices and Maps

## Goal

Build a precise mental model of slices and maps so interview questions about hidden mutations, reallocations, memory behavior, and concurrency do not catch you off guard.

This module is foundational for later topics like escape analysis, GC pressure, race conditions, and production debugging.

## Part 1 — Arrays vs slices

### Arrays

An array in Go has fixed size and the size is part of its type.

Examples:

- `[3]int`
- `[10]byte`

Important properties:

- copying an array copies all elements
- arrays are value types
- arrays are rarely used directly in application code, but they matter for understanding slices

### Slices

A slice is a lightweight descriptor over an underlying array.

Conceptually, a slice header contains:

- pointer to backing array
- length
- capacity

Important properties:

- copying a slice copies only the header, not the underlying data
- two slices can share the same backing array
- mutation through one slice can affect another slice

## Part 2 — len and cap

For a slice:

- `len` = number of visible elements
- `cap` = how much of the backing array is available starting from the slice pointer

Example idea:

- if a slice has spare capacity, `append` may reuse the same backing array
- if capacity is exhausted, `append` may allocate a new backing array

Interview implication:
You must always ask whether two slices still share backing storage after append.

## Part 3 — append behavior

`append` is one of the most common interview traps.

### Case 1 — append fits into existing capacity
Then Go usually reuses the same backing array.

Consequence:
- the original slice and the appended slice may still alias the same memory

### Case 2 — append exceeds capacity
Then Go allocates a new backing array and copies elements.

Consequence:
- the new slice no longer aliases the old backing array
- later mutations may diverge

### Practical rule
If correctness depends on isolation, make an explicit copy instead of relying on append behavior.

Example copy patterns:

- `dst := append([]T(nil), src...)`
- `dst := make([]T, len(src)); copy(dst, src)`

## Part 4 — slice aliasing

Aliasing means multiple slices point into the same underlying array.

This causes bugs like:

- unexpected mutation in helper functions
- corrupted buffers
- data races when shared between goroutines
- memory retention when a small subslice keeps a large backing array alive

### Classic bug shape

A function receives a slice, takes a subslice, appends to it, and unexpectedly mutates the caller-visible data.

### Interview takeaway

When reading slice code, always check:

1. where the backing array comes from
2. whether slices overlap
3. whether append can reuse capacity
4. whether a copy is needed for ownership isolation

## Part 5 — slicing and memory retention

A small subslice can keep a large array alive.

Example scenario:

- load a 100 MB file into memory
- take `buf[:100]`
- store that small slice somewhere long-term

Problem:
the tiny slice still references the large backing array, so GC cannot reclaim it.

### Fix

Copy the small piece into a new slice if the original large buffer is no longer needed.

This matters in:

- log processing
- network parsers
- file scanners
- cache layers
- Kubernetes/operator controllers processing large payloads

## Part 6 — maps under the hood

At a high level, a Go map is a hash table with buckets.

You do not need runtime source-level detail in most interviews, but you do need the right mental model:

- keys are hashed
- entries are distributed across buckets
- lookup and insert are approximately O(1) on average
- growth and rehash-related work happens as the map changes

### Important practical implications

- iteration order is not stable
- maps can grow
- map operations are not free in hot loops
- maps are reference-like structures from the programmer's perspective

## Part 7 — why maps are not thread-safe

Go maps are not safe for concurrent read/write access without synchronization.

### Safe-ish case
- concurrent reads only, if no writes happen

### Unsafe case
- any write combined with another write or read without synchronization

This can lead to:

- runtime panic in some cases
- data race
- memory corruption risk avoided by runtime checks in some scenarios, but still incorrect code

### Interview answer pattern

If multiple goroutines access a map