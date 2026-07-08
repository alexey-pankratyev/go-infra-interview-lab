# Go Answer Review Rubric

Use this rubric to keep feedback high-signal and interview-focused.

## 1. Correctness

- Does the answer solve the stated task?
- Are outputs, return values, and error cases handled correctly?
- Are there hidden panics, deadlocks, or invalid assumptions?

## 2. Edge Cases

- Empty input, nil input, zero values
- Closed channels, canceled contexts, timeouts
- Reallocation, aliasing, shared mutable state
- Partial failures and cleanup paths

## 3. Runtime and Memory

- Avoidable allocations
- Escape analysis implications
- Stack vs heap behavior
- GC pressure and object lifetime
- Blocking calls and scheduler impact

## 4. Concurrency and Safety

- Race conditions
- Goroutine leaks
- Channel close ownership
- WaitGroup misuse
- Mutex contention or unsafe map access
- Backpressure and cancellation propagation

## 5. Production Fit

- Observability and debuggability
- Clear error handling
- Timeout and retry behavior
- Testability
- Operational failure modes

## 6. Interview Signal

Classify the answer as:

- Strong: correct plus production nuance.
- Acceptable: mostly correct, minor gaps.
- Risky: compiles or sounds plausible but misses a serious runtime or production issue.
- Incorrect: fails the task or relies on a false mental model.
