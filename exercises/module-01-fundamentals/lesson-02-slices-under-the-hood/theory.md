# Theory - Slices Under the Hood

## Internal Representation

A slice is a three-field struct:

```text
type SliceHeader struct {
    Data uintptr  // pointer to backing array
    Len  int      // number of accessible elements
    Cap  int      // total capacity from Data to end of backing array
}
```

When you pass a slice to a function or assign it to another variable, Go copies this header — not the underlying array.

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]
// b.Data points into a's backing array at index 1
// b.Len = 2, b.Cap = 4  (capacity runs to end of a's array)
```

Mental model:

```text
backing array:  [1][2][3][4][5]
a:              ptr=0, len=5, cap=5
b:              ptr=1, len=2, cap=4  ← same array, offset by 1
```

## Aliasing

Because `b` shares the backing array with `a`, writing through `b` mutates `a`:

```go
b[0] = 99
// backing array is now: [1][99][3][4][5]
// a[1] is now 99
```

This is the aliasing hazard. Two slice variables, one mutation, two callers surprised.

## Append and Reallocation

`append` adds elements to a slice. The behavior depends on remaining capacity.

**Capacity available** — append writes into the existing backing array:

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]          // len=2, cap=4
b = append(b, 100)   // writes 100 at index 3 of backing array
// backing array: [1][99][3][100][5]  ← a[3] is now 100
// b: len=3, cap=4
```

`b` still shares the array with `a`. The append silently overwrote `a[3]`.

**Capacity exhausted** — append allocates a new backing array, copies elements, and returns a new slice header:

```go
b = append(b, 200, 300, 400)  // exceeds cap, new array allocated
// b now points to a different array
// a is unaffected from this point forward
```

After reallocation, `a` and `b` are independent.

## The Reallocation Rule

- `len < cap`: append writes in place, shares backing array, can mutate other slices
- `len == cap`: append allocates, copies, returns independent slice

This is why `append` returning a new value matters:

```go
b = append(b, x)  // always reassign — b may be a new slice
```

## Common Traps

**Subslice append clobbers the parent:**

```go
a := make([]int, 3, 5)  // len=3, cap=5
b := a[:2]              // len=2, cap=5 — shares array
b = append(b, 99)       // writes at index 2, mutating a[2]
```

**Function that appends does not affect caller's slice header:**

```go
func addOne(s []int) {
    s = append(s, 1)  // local header updated, caller's len unchanged
}
```

The caller's `len` stays the same even if the backing array was mutated — because the slice header was copied.

**Defensive copy before independent mutation:**

```go
b := make([]int, len(a))
copy(b, a)
// now b is independent
```

## Interview Signals

Be ready to explain:

- what three fields a slice header contains
- when append allocates a new array vs writes in place
- how a subslice can silently mutate the parent
- why a function that appends cannot update the caller's length
- how to make a fully independent copy of a slice
