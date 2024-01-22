# Golang Monorepo

This repository is an example monorepo using golang.

Notes:

- Due to the limitations of the go module system, we use a single `go.mod` file at the root level.
- We use `lib` instead of `pkg` because the modules under `lib` are meant for direction consumption in the services.

## Stack

This repository uses [go-chi](https://github.com/go-chi/chi) as a router, it uses [sqlc](https://sqlc.dev/) for the DAL and PGX as the DB driver
for Postgres.

## Setup

### pre-commit

We use pre-commit to orchestrate linting.

1. install [pre-commit](https://pre-commit.com/) on your machine.
2. install the pre-commit hooks by executing `pre-commit install` in the repository root.

#### pre-commit commands

- `pre-commit autoupdate` to update the hooks
- `pre-commit run --all-files` to execute against all files

### Docker

There is a `dockerfile` in the repository root which uses `distroless` as the production image. Use docker compose for local development:

- `docker compose up --build`

### SQLC

1. [sqlc](https://sqlc.dev/).
2. run `sqlc generate` or any of the other commands in the root.

## Migrations

This project uses [goose](https://github.com/pressly/goose) for database migrations.

- Use the `taskfile` commands to run the database migrations.
- Make sure to have `DATABASE_URL` in your \*.env file.

### TaskFile

There is a `Taskfile.yaml` file, which is a task runner. This project uses [go-task](https://taskfile.dev/) for running task, and it's pretty simple to use.

- For installation of `taskfile` use [these instructions](https://taskfile.dev/installation/).
- For running the tasks run: `task <task-name>`, ex: `task greet`
