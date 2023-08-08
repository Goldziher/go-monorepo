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

### Migrations
In our case `<migration-dir>` is `migrations` folder
1. run `migrate create -ext sql -dir <migration-dir> -seq <migration-name>` to generate migrations
2. run `migrate -path <migration-dir> -database <DATABASE_URL> -verbose <up/down>` to apply migrations
