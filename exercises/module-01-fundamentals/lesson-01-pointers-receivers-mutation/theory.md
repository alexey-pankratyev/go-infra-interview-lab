# Theory - Pointers, Receivers, and Mutation Semantics

## Core Model

Go passes arguments by value. That means every function call receives a copy of each argument.

For plain values, the copy is independent:

```go
func set(n int) {
    n = 10
}
```

The caller's integer is unchanged.

For pointers, the pointer value is copied, but both pointer values still point to the same object:

```go
func set(p *int) {
    *p = 10
}
```

The caller's integer is changed through the shared pointed-to object.

## Receiver Choice

A method receiver is just a function argument with special syntax.

A value receiver copies the receiver:

```go
func (c Config) SetRetries(n int) {
    c.Retries = n
}
```

This mutates only the local copy.

A pointer receiver can mutate the original object:

```go
func (c *Config) SetRetries(n int) {
    c.Retries = n
}
```

This mutates the caller's `Config`.

## Reference-Containing Fields

Some values contain references to shared backing data. Maps are the important example for this lesson.

If a struct contains a map and the struct is copied, both struct copies still refer to the same map data:

```go
type Config struct {
    Labels map[string]string
}
```

A value receiver copies the `Config`, but not the underlying map data. Mutating `c.Labels` mutates the shared map.

## Interview Signals

Be ready to explain:

- whether a method mutates the original receiver or a copy
- why a copied struct can still share internal data through maps, slices, pointers, or channels
- when pointer receivers are required for mutation or avoiding expensive copies
- how receiver choice affects interface method sets

## Production Implications

- Avoid surprising APIs where one value receiver method mutates shared internal state and another does not.
- Be careful with mutable maps inside configuration objects.
- Prefer explicit ownership: either make config immutable, clone mutable fields, or use pointer receivers consistently when mutation is intended.
- Under concurrency, shared maps require synchronization or copy-on-write design.
