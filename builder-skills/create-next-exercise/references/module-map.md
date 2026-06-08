# Module Map

Move through modules strictly in order. Each module should close its current exercise before the next module begins.

## 1. Quick Go Fundamentals Refresh

Short lessons only. Focus on interview-relevant fundamentals rather than beginner control-flow drills.

Coverage:

- values, pointers, interfaces, errors
- defer, panic/recover, package boundaries
- struct methods and receiver tradeoffs

## 2. Arrays, Slices, and Maps Under the Hood

Coverage:

- arrays vs slices
- len/cap and append behavior
- reallocation and aliasing
- map bucket model
- why maps are not thread-safe

## 3. Memory Management and Allocations

Coverage:

- pointers
- stack vs heap
- escape analysis
- GC basics and stop-the-world effects
- allocation reduction

## 4. Concurrency and Go Runtime Internals

Coverage:

- goroutines vs OS threads
- GMP scheduler and work stealing
- channels, close semantics, reads from closed channels
- sync primitives: Mutex, RWMutex, WaitGroup, Once, Map
- select, worker pools, fan-in/fan-out
- context cancellation and timeouts

## 5. Systems Programming and Networking in Go

Coverage:

- os/exec
- signals and graceful shutdown
- net/http transports and connection reuse
- HTTP timeouts
- JSON, Protobuf, and gRPC tradeoffs

## 6. Testing, Race Detection, Benchmarking, and Profiling

Coverage:

- unit tests
- benchmarks
- race detector
- pprof
- memory leaks and goroutine leaks

## 7. Final Interview Drill

Coverage:

- mixed practical scenarios
- runtime explanation under pressure
- production debugging tradeoffs
- concise senior-level communication
