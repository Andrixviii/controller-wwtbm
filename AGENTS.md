# AGENTS.md

## Purpose
This repository is a simple Go CRUD web app for managing quiz questions in PostgreSQL.
AI agents should make minimal, safe, and consistent changes with the current project structure.

---

## Tech Stack
- Go (`module: go-web`)
- HTTP: `net/http`
- Templates: `html/template`
- DB: PostgreSQL (`github.com/lib/pq`)
- Env: `github.com/joho/godotenv`

---

## Project Structure
- `main.go` — app entrypoint and route registration
- `config/db_connection.go` — database connection setup
- `controllers/homecontroller/homecontroller.go` — home page handler
- `controllers/categorycontroller/categorycontroller.go` — actions handlers (list/create/edit/delete)
- `models/question_model/question_model.go` — DB query layer for questions
- `entitas/questions.go` — entity definitions
- `view/` — HTML templates

---

## Run Instructions
### Required env
- `DATABASE_URL` (PostgreSQL DSN)

### Local run
```bash
go run main.go
```

- Server runs at `http://localhost:8080`.

# Routes
`GET / — home`
`GET /actions — list questions + pagination`
`GET /actions/create — create form`
`POST /actions/create — insert question`
`GET /actions/edit — edit` (not fully implemented yet)
`GET /actions/delete — delete` (not fully implemented yet)

## Database Conventions (IMPORTANT)
- Use PostgreSQL placeholders: `$1`, `$2`, …
- Do not use `?` placeholders.
- For insert + id retrieval, use `RETURNING id` with `QueryRow().Scan(&id)`.
- After iterating `rows.Next()`, always check `rows.Err()`.

## Controller + Template Conventions
- Avoid panic for request-level errors.
- Return proper HTTP errors or re-render page with message.
- Keep template data keys consistent and case-sensitive.
- `view/actions/index.html` expects:
  - `questions`
  - `CurrentPage`
  - `TotalPages`
  - `PrevPage`
  - `NextPage`
- For form submission, ensure endpoint matches handler.

## Create Form Conventions
- Create form must post to `/actions/create`.
- In HTML forms, set `name` on `<input>/<select>`, not on `<option>`.

## Coding Guidelines for Agents
- Keep changes focused and minimal.
- Preserve existing package layout and naming style.
- Don't refactor unrelated code.
- Don't introduce new dependencies unless required.
- Don't modify routes or DB schema unless task explicitly asks for it.
- If behavior changes, update corresponding template/controller/model together.

## Validation Checklist (before finishing a task)
- `go build ./...` passes
- `/actions` renders without panic
- Pagination works (`?page=`)
- Create form submits and inserts row into `questions`
- No silent failure on insert errors

## Known Gaps / TODO
- `Edit` and `Delete` handlers are placeholders and need full 
  implementation.
- Error handling can be standardized across controllers.
- Optional: split pagination logic into reusable function service.