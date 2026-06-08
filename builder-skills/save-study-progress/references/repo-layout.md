# Study Repository Layout

Use stable, predictable paths so future sessions can continue without searching through scattered files.

## Top-Level Files

- `README.md`: curriculum overview, module order, and how to use the repository.
- `progress/progress.md`: current module, completed lessons, open tasks, and checklist state.

## Notes

Use `notes/` for concise topic summaries and interview cheat sheets.

Suggested files:

- `notes/module-01-go-fundamentals.md`
- `notes/module-02-slices-maps.md`
- `notes/module-03-memory-gc.md`
- `notes/module-04-concurrency-runtime.md`
- `notes/module-05-networking-systems.md`
- `notes/module-06-testing-profiling.md`

## Exercises

Use `exercises/` for lesson-scoped practice.

Preferred structure:

```text
exercises/
└── module-XX-topic/
    ├── README.md
    └── lesson-YY-topic-name/
        ├── README.md
        ├── theory.md
        ├── find-bug/
        │   ├── task.md
        │   └── main.go
        └── mini-script/
            ├── task.md
            └── main.go
```

Each lesson folder should contain:

- `theory.md` for the short practical foundation
- `find-bug/` for predict-output or bug-hunting practice
- `mini-script/` for a small runnable implementation task

Go starter files should be runnable from their directory with:

```bash
go run main.go
```

## Code

Use `code/` only for corrected examples, larger snippets, or reusable drafts that no longer belong to a single exercise folder.

Suggested naming:

- `code/module-02-slice-aliasing.go`
- `code/module-04-worker-pool.go`

## Save Rules

- Update existing module and lesson files when material belongs to the same module.
- Add a new lesson folder when a new topic starts.
- Keep saved answers concise unless the exact wording is useful for review.
- Do not keep duplicate flat exercise files when the material has moved into a lesson folder.
