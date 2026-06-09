# Theory - Maps Under the Hood

## Internal Representation

A Go map is a pointer to a runtime `hmap` struct. When you assign or pass a map, you copy the pointer — not the data.

```go
a := map[string]int{"x": 1}
b := a         // b and a point to the same hmap
b["x"] = 99    // mutates a too
```

This is the same aliasing hazard as with slices, but without a visible header struct. The pointer is always hidden.

## Bucket Model

Internally, a map is an array of buckets. Each bucket holds up to 8 key-value pairs. When the load factor exceeds ~6.5 keys per bucket, the map grows: a new bucket array is allocated and keys are incrementally migrated.

Key points for interviews:

- iteration order is randomized on purpose to prevent reliance on ordering
- map growth and key migration happen incrementally, not all at once
- the zero value of a map is `nil` — reads return zero values, writes panic

```go
var m map[string]int
fmt.Println(m["x"]) // 0 — safe read from nil map
m["x"] = 1         // panic: assignment to entry in nil map
```

## Why Maps Are Not Thread-Safe

Go's map implementation does not use internal locking. Concurrent reads are safe. Concurrent writes, or a concurrent read and write, are not — the runtime detects this and panics with:

```text
fatal error: concurrent map read and map write
```

This is a deliberate design choice: adding a lock to every map operation would slow down the majority of use cases that are single-goroutine.

```go
// This is a data race — do not do this
m := map[string]int{}
go func() { m["a"] = 1 }()
go func() { m["b"] = 2 }()
```

## Safe Concurrent Access Patterns

**sync.Mutex — simple and explicit:**

```go
type SafeMap struct {
    mu sync.Mutex
    m  map[string]int
}

func (s *SafeMap) Set(k string, v int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m[k] = v
}

func (s *SafeMap) Get(k string) (int, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    v, ok := s.m[k]
    return v, ok
}
```

**sync.RWMutex — when reads dominate:**

```go
func (s *SafeMap) Get(k string) (int, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    v, ok := s.m[k]
    return v, ok
}
```

Multiple goroutines can hold `RLock` simultaneously. A writer waits for all readers to release before acquiring the write lock.

**sync.Map — for specific access patterns:**

`sync.Map` is optimized for two cases:
- key is written once and read many times
- many goroutines read/write disjoint sets of keys

It is not a general-purpose replacement for `map + Mutex`. For general concurrent maps, prefer `Mutex` — it is easier to reason about.

## The Comma-Ok Idiom

Always use comma-ok when you need to distinguish a missing key from a zero value:

```go
v, ok := m["key"]
if !ok {
    // key does not exist
}
```

Without comma-ok, a missing `int` key returns `0`, which is indistinguishable from a key that was explicitly set to `0`.

## Interview Signals

Be ready to explain:

- why map assignment copies a pointer, not data
- what happens on concurrent map writes and why the runtime panics
- the difference between Mutex, RWMutex, and sync.Map and when to use each
- why map iteration order is random
- nil map read vs write behavior
