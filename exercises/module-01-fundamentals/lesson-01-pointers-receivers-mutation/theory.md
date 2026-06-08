# Theory - Pointers, Receivers, and Mutation Semantics

## Core Model

Go passes arguments by value. Every function call receives a copy of each argument.

For plain values, the function receives an independent local copy:

```go
package main

import "fmt"

func setRetries(retries int) {
    retries = 10
}

func main() {
    retries := 3

    setRetries(retries)

    fmt.Println(retries) // 3
}
```

`setRetries(retries)` copies the value `3` into the function parameter `retries`. Inside the function, `retries = 10` changes only that local copy. The caller's variable is still `3`.

Mental model:

```text
main has:        retries = 3
function gets:   retries = copy of 3
function sets:   local retries = 10
main still has:  retries = 3
```

For pointers, the pointer value is still copied, but the copied pointer points to the same object as the original pointer:

```go
package main

import "fmt"

func setRetries(retries *int) {
    *retries = 10
}

func main() {
    retries := 3

    setRetries(&retries)

    fmt.Println(retries) // 10
}
```

`setRetries(&retries)` copies the address of `retries`. Inside the function, `*retries = 10` follows that copied address and mutates the caller's original variable.

Mental model:

```text
main has:          retries = 3
function gets:     copy of address pointing to retries
function writes:   value at that address = 10
main now has:      retries = 10
```

The key distinction: assigning to a value parameter changes the local copy; dereferencing a pointer parameter can change the caller's object.

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
