# Setup

## Goal

Use this repository as a reusable kit for creating your own Go interview mentor agent.

The intended workflow:

1. Fork this repository
2. Create your own agent in ChatGPT Agent Builder
3. Connect GitHub
4. Point the agent to your fork
5. Attach reusable skills from `builder-skills/`
6. Let the agent teach, save progress, review answers, and add new exercises in your fork
7. If you improve the materials, open a PR back to the upstream repository

---

## 1. Fork the repository

Fork this repository into your own GitHub account.

Example:

- upstream: `alexey-pankratyev/go-infra-interview-lab`
- your fork: `your-user/go-infra-interview-lab`

---

## 2. Create your agent

In ChatGPT Agent Builder:

1. Create a new agent
2. Give it a name
3. Paste the base instructions from `agent/system-instructions.md`
4. Add starter prompts if needed
5. Connect GitHub

Recommended default repo for your agent:

- `your-user/go-infra-interview-lab`

---

## 3. Attach reusable skills

Use the skill folders from `builder-skills/`.

Recommended initial skills:

- `go-github-onboarding`
- `save-study-progress`
- `review-go-answer`
- `create-next-exercise`
- `sync-module-state`

Each skill folder should be uploaded as a Builder skill.

---

## 4. Suggested repository structure

The agent is expected to use this structure:

- `README.md`
- `progress/progress.md`
- `notes/`
- `exercises/`
- `code/`
- `builder-skills/`
- `agent/`
- `templates/`

---

## 5. First run

A good first prompt to the agent:

`Show me the high-level module plan and start module 1 after I confirm it.`

A good first GitHub sync prompt:

`Save the starter structure and study plan to my repository.`

---

## 6. Contribution workflow

Work in your own fork.

If you improve:

- exercises
- notes
- templates
- skill packages
- agent instructions

open a PR to the upstream repository so the shared material improves for everyone.

---

## 7. Design principles

- the fork is the learner's personal lab
- the agent teaches against the learner's fork
- saved progress lives in the learner's fork
- improvements flow back through PRs to upstream