# Module 01 — Go Fundamentals Refresh

## Goal

Quickly restore the minimum Go base required for deeper interview topics.

This module is not about beginner syntax drilling. It is about rebuilding fluency so later topics like slices, concurrency, memory, and profiling feel natural.

## What to refresh

- functions and multiple return values
- named returns
- error handling
- structs and methods
- pointer vs value receivers
- interfaces
- packages and visibility
- strings, bytes, runes
- basic slice and map usage as preparation for deeper internals

## High-level reminders

### 1. Go favors explicitness
Go code is intentionally direct:

- explicit error returns
- explicit interfaces
- explicit ownership of values and pointers

In interviews, over-engineering simple code is usually a negative signal.

### 2. Interfaces are satisfied implicitly
A type implements an interface by having the required methods.

Important interview implication:
- think in terms of behavior contracts
- keep interfaces small
- avoid designing giant “god interfaces”

### 3. Value vs pointer semantics matter early
Even before deep memory analysis, you should already notice:
- copying structs can be cheap or expensive depending on size
- pointer receivers affect mutation semantics
- method sets matter for interface satisfaction

### 4. Strings are bytes, not characters
A Go string is a read-only byte slice.

Important consequences:
- `len(s)` returns bytes
- indexing gives bytes
- Unicode-safe iteration usually means `for range`
- `rune` matters when interviewers test Unicode awareness

## Common interview traps in this module

### Trap 1 — assuming `len(string)` means characters
It means bytes.

### Trap 2 — misunderstanding receiver choice
Use:
- value receivers for small immutable-like structs
- pointer receivers when mutating or avoiding copies