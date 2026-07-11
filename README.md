# Task2Schedule

Task2Schedule turns task lists into actionable schedules. Product focus:

- Goal setting
- Task and todo management
- Schedule generation
- Google Calendar integration

## Repository

```text
AGENT.md       Root development guide
backend/       Go API, PostgreSQL schema, Swagger
docs/          Project reference notes
frontend/      React client
```

Read [backend/AGENT.md](backend/AGENT.md) before backend work. Read
[frontend/README.md](frontend/README.md) before frontend work.

## Backend quick start

Prerequisites: Docker, Go, Air.

```bash
cd backend
cp .env.example .env # first setup only
make env-up
make migrate-up
air
```

API docs: <http://localhost:8080/swagger/index.html>

Useful backend commands:

```bash
make go-fmt
make test
make build
make swagger-gen
```
