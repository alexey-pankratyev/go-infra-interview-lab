# State Files

Use these files to keep progress durable and easy to resume.

## Durable Defaults

`go-study-defaults.yaml`

Suggested fields:

```yaml
repository_full_name: alexey-pankratyev/go-infra-interview-lab
branch: main
root_path: .
current_module: 1
current_lesson: fundamentals-refresh
last_task_category: predict_or_find_bug
```

Store this only in a durable location explicitly chosen by the user (repo root or a config path).

## Repository Progress

`progress/progress.md`

Suggested sections:

- Current position
- Module checklist
- Completed exercises
- Open weak spots
- Next recommended task

## Module Exercise Files

`exercises/module-XX-topic/lesson-YY-topic/`

Suggested contents per exercise:

- Exercise prompt (`task.md`)
- User answer
- Review (`review.md`)
- Corrected solution (`main.go`)
- Follow-up notes

## Topic Notes

`notes/topic.md`

Keep notes short and interview-focused. Prefer production implications over textbook definitions.

## Update Rules

- Update progress after every reviewed exercise.
- Update weak spots only when they are stable and useful.
- Keep current module state synchronized before creating the next exercise.
