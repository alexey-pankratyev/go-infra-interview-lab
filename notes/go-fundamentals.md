# Go Fundamentals Refresh

Module 1 is a compact senior-level refresh. The goal is not to relearn syntax, but to sharpen the Go mental models that commonly show up in platform and infrastructure interviews.

## Lesson 1: Pointers, Receivers, and Mutation Semantics

Key ideas:

- Go passes arguments by value. Passing a pointer copies the pointer value, but both copies still refer to the same object.
- A value receiver gets a copy of the receiver. A pointer receiver can mutate the original object and avoids copying large structs.
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
- `panic`/`recover` is for exceptional boundaries, not normal control flow.

Interview signals:

- Know when deferred cleanup is correct and when it extends resource lifetime too long.
- Understand defer ordering around locks, files, network connections, and spans.
- Keep panic recovery at process, worker, or request boundaries when appropriate.

Production concerns:

- Deferring cleanup inside long loops can leak resources until the function returns.
- Recovering without logging or surfacing failure can hide broken states.
- Cleanup must account for partial initialization and failed startup paths.

## Module 1 Completion Criteria

The module is complete when the user can:

- Predict pointer/value mutation behavior.
- Explain receiver choice tradeoffs.
- Identify typed-nil interface bugs.
- Preserve useful error context.
- Reason about defer and cleanup lifetime.
