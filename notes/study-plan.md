# Study Plan

## Goal

Fast, interview-focused Golang refresh for Senior DevOps / Platform Engineer roles.

Primary emphasis:

- systems programming
- concurrency and runtime internals
- memory behavior and allocations
- networking and HTTP
- debugging, profiling, and production tradeoffs

## Learning rules

- one module at a time
- one practical task at a time
- answer first, review after
- focus on interview-relevant reasoning
- no long theory dumps unless explicitly requested

## Module 1 — Quick Go fundamentals refresh

### Goal
Refresh the minimum base needed for deeper topics without wasting time on beginner syntax drills.

### Topics
- functions and multiple return values
- methods and receivers
- structs and interfaces
- errors and wrapping
- packages and visibility
- strings, runes, bytes
- basic slice and map usage refresh

### Outcome
Be comfortable reading and writing small Go snippets quickly in interview conditions.

## Module 2 — Arrays, slices, and maps under the hood

### Goal
Build a precise mental model of Go collections and their runtime behavior.

### Topics
- arrays vs slices
- slice header: pointer, len, cap
- append behavior
- reallocation and backing array changes
- aliasing and unexpected mutation
- map internals at a high level
- buckets and growth
- why maps are not thread-safe

### Outcome
Confidently explain tricky code involving slices and maps and spot hidden bugs.

## Module 3 — Memory management and allocations

### Goal
Understand how Go places values in memory and how that affects performance.

### Topics
- pointers
- value vs pointer semantics
- stack vs heap
- escape analysis
- common causes of heap allocations
- garbage collector basics
- tri-color marking
- stop-the-world phases
- reducing allocations in hot paths

### Outcome
Be able to reason about memory pressure, GC overhead, and optimization opportunities.

## Module 4 — Concurrency and Go runtime internals

### Goal
Master the concurrency model expected in senior platform interviews.

### Topics
- goroutines vs OS threads
- GMP scheduler
- work stealing
- channels: buffered vs unbuffered
- channel closing semantics
- reading from closed channels
- select
- sync.Mutex
- sync.RWMutex
- sync.WaitGroup
- sync.Once
- sync.Map
- worker pool
- fan-in / fan-out
- context cancellation and timeouts
- goroutine leaks
- race conditions
- deadlocks

### Outcome
Be able to design and debug production-grade concurrent Go systems.

## Module 5 — Systems programming and networking

### Goal
Focus on platform-oriented Go use cases.

### Topics
- os/exec
- process management
- signals and graceful shutdown
- net/http server basics
- HTTP client pitfalls
- custom Transport
- timeouts
- connection reuse
- JSON serialization
- Protobuf basics
- gRPC basics

### Outcome
Be ready for practical platform and backend interview scenarios.

## Module 6 — Testing, benchmarking, and profiling

### Goal
Be able to verify correctness and debug performance issues in production-style code.

### Topics
- unit tests
- table-driven tests
- benchmarks
- race detector
- pprof
- CPU profiling
- memory profiling
- leak investigation
- goroutine leak patterns

### Outcome
Be able to explain how to find and fix performance and correctness issues.

## Module 7 — Final interview drill

### Goal
Simulate realistic interview pressure.

### Format
- bug hunts
- output prediction tasks
- micro-implementation tasks
- architecture/debugging scenarios
- code review style feedback

### Outcome
Convert refreshed knowledge into interview performance.

## Suggested repository usage

- `progress/progress.md` — module completion and checkpoints
- `notes/` — short topic summaries
- `exercises/` — tasks and answers
- `code/` — snippets and corrected examples

## First milestone

Finish:

- Module 1
- one slice bug hunt
- one map/concurrency bug hunt
- one worker pool exercise
- one graceful shutdown exercise
- one profiling/debugging scenario