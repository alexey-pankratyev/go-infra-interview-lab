# Mini-Script - Thread-Safe Counter Map

## Task Type

Write a small runnable function.

## Goal

Practice safe concurrent map access using sync.RWMutex.

## Instructions

Open `main.go` and implement the TODOs.

The script should:

1. Implement `Inc(key string)` — increment the counter for the given key safely.
2. Implement `Get(key string) int` — return the current count for the given key safely.
3. Implement `Top() (string, int)` — return the key with the highest count. Read-only operation.
4. Run 1000 goroutines, each incrementing one of 5 keys, then print the top key and total sum.

Run it with:

```bash
go run main.go
```

And verify no race with:

```bash
go run -race main.go
```

## Expected Output

```text
top key: worker-N  count: ~200
total: 1000
```

(exact top key and count will vary by run)

## Review Focus

When reviewing the answer, focus on:

- correct lock/unlock pairing (defer is preferred)
- whether RLock is used for reads and Lock for writes
- whether the zero-value missing key case is handled in Get
- whether Top holds the lock for the full read operation
