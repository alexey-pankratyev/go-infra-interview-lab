# Module 1: Quick Go Fundamentals Refresh

## Lesson 1: Pointers, Receivers, and Mutation Semantics

### Concept Reminder

- Go passes arguments by value, including structs and pointers.
- A pointer receiver can mutate the original object; a value receiver mutates only its local receiver copy.
- Receiver choice matters for API consistency, interface satisfaction, copying cost, and mutation semantics.

## Exercise 1: Predict the Output / Find the Bug

Category: find the bug / predict the output

Read the code and answer the questions below. Do not run it first.

```go
package main

import "fmt"

type Config struct {
    Retries int
    Labels  map[string]string
}

func (c Config) SetRetries(n int) {
    c.Retries = n
}

func (c Config) SetLabel(k, v string) {
    c.Labels[k] = v
}

func update(cfg Config) {
    cfg.SetRetries(5)
    cfg.SetLabel("env", "prod")
}

func main() {
    cfg := Config{
        Retries: 3,
        Labels:  map[string]string{"team": "platform"},
    }

    update(cfg)

    fmt.Println(cfg.Retries)
    fmt.Println(cfg.Labels["env"])
}
```

### Questions

1. What does this program print?
2. Why does one field appear unchanged while the other appears changed?
3. What would you change if this were production configuration code?

### Review Notes Placeholder

The user's answer and review should be added here after the exercise is completed.

### Corrected Solution Placeholder

Do not fill this section before the user answers.
