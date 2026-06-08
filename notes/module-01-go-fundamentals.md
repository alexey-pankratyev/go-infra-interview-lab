# Module 01 - Go Fundamentals Refresh

## Goal

Quickly restore the minimum Go base required for deeper interview topics.

This module is not about beginner syntax drilling. It is about rebuilding fluency so later topics like slices, concurrency, memory, and profiling feel natural.

## What to Refresh

- functions and multiple return values
- named returns
- error handling
- structs and methods
- pointer vs value receivers
- interfaces
- packages and visibility
- strings, bytes, runes
- basic slice and map usage as preparation for deeper internals

## High-Level Reminders

### 1. Go Favors Explicitness

Go code is intentionally direct:

- explicit error returns
- explicit interfaces
- explicit ownership of values and pointers

In interviews, over-engineering simple code is usually a negative signal.

### 2. Interfaces Are Satisfied Implicitly

A type implements an interface by having the required methods.

Important interview implications:

- think in terms of behavior contracts
- keep interfaces small
- avoid designing giant "god interfaces"
- remember that interface values contain both a dynamic type and a dynamic value

### 3. Value vs Pointer Semantics Matter Early

Even before deep memory analysis, you should already notice:

- Go passes arguments by value
- copying structs can be cheap or expensive depending on size
- pointer receivers affect mutation semantics
- method sets matter for interface satisfaction
- passing a pointer copies the pointer value, but both copies still refer to the same object

### 4. Strings Are Bytes, Not Characters

A Go string is a read-only byte sequence.

Important consequences:

- `len(s)` returns bytes
- indexing gives bytes
- Unicode-safe iteration usually means `for range`
- `rune` matters when interviewers test Unicode awareness

## Lesson 1: Pointers, Receivers, and Mutation Semantics

Key ideas:

- A value receiver gets a copy of the receiver.
- A pointer receiver can mutate the original object and avoids copying large structs.
- Receiver choice affects API consistency, mutation semantics, method sets, interface satisfaction, and sometimes allocation behavior.

Interview signals:

- Explain whether a function mutates the caller's object or only a local copy.
- Notice when a large struct is copied unnecessarily.
- Understand when pointer receivers are required for interface implementation.
- Avoid using pointers reflexively when values are small, immutable, or clearer by value.

Production concerns:

- Mutable shared state becomes dangerous under concurrency.
- Pointer-heavy designs can make ownership and lifetime harder to reason about.
- Extra allocations may appear when values escape to the heap.

## Lesson 2: Interfaces, Nil, and Errors

Key ideas:

- An interface value stores both a dynamic type and a dynamic value.
- An interface is nil only when both the dynamic type and dynamic value are nil.
- Error handling should preserve context without hiding the original failure.

Interview signals:

- Recognize typed-nil interface bugs.
- Use `errors.Is` and `errors.As` for wrapped errors.
- Avoid returning opaque errors from infrastructure code where debugging context matters.

Production concerns:

- Poor error context slows incident response.
- Typed nils can create misleading `err != nil` or interface checks.
- Logging and returning the same error at every layer can create noisy traces.

## Lesson 3: Defer, Cleanup, and Boundaries

Key ideas:

- `defer` runs at function return, in last-in-first-out order.
- Deferred functions capture variables according to normal closure rules.
- `panic` and `recover` are for exceptional boundaries, not normal control flow.

Interview signals:

- Know when deferred cleanup is correct and when it extends resource lifetime too long.
- Understand defer ordering around locks, files, network connections, and spans.
- Keep panic recovery at process, worker, or request boundaries when appropriate.

Production concerns:

- Deferring cleanup inside long loops can leak resources until the function returns.
- Recovering without logging or surfacing failure can hide broken states.
- Cleanup must account for partial initialization and failed startup paths.

## Common Interview Traps in This Module

### Trap 1: Assuming `len(string)` Means Characters

It means bytes.

### Trap 2: Misunderstanding Receiver Choice

Use value receivers for small immutable-like structs. Use pointer receivers when mutating or avoiding expensive copies.

### Trap 3: Missing Typed-Nil Interface Behavior

An interface with a non-nil dynamic type and a nil dynamic value is not nil.

### Trap 4: Deferring Cleanup in Long-Running Loops

Deferred cleanup runs when the surrounding function returns, not at the end of each loop iteration.

## Module 1 Completion Criteria

The module is complete when the user can:

- Predict pointer/value mutation behavior.
- Explain receiver choice tradeoffs.
- Identify typed-nil interface bugs.
- Preserve useful error context.
- Reason about defer and cleanup lifetime.
