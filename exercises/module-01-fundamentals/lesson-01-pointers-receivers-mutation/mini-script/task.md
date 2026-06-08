# Mini-Script - Normalize Config

## Task Type

Write a small runnable function.

## Goal

Practice pointer receivers, value semantics, and explicit mutation.

## Instructions

Open `main.go` and implement the TODOs.

The script should:

1. Keep `Config` as the main type.
2. Implement `Normalize` so it mutates the original config.
3. Ensure `Retries` is at least `1`.
4. Ensure `Labels` is initialized when nil.
5. Ensure `Labels["env"]` is set to `"dev"` when missing.
6. Keep the code simple and runnable from the console.

Run it with:

```bash
go run main.go
```

## Expected Output

```text
retries=1 env=dev
retries=3 env=prod
```

## Review Focus

When reviewing the answer, focus on:

- whether mutation affects the caller's config
- whether nil maps are handled safely
- whether receiver choice is intentional
- whether the code avoids unnecessary complexity
