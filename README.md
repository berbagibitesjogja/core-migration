# Core System Database Migration
Database migration for BBJ's Core System written in raw PostgreSQL using Goose database migration tool.

## Prerequisites
1. Go 1.16+ - https://go.dev/ (Optional)
2. Goose - https://pressly.github.io/goose/installation
3. PostgreSQL / Other RDBMS (Not Tested)
4. Podman / Docker (Not Tested)

## Getting Started
1. Clone the repository
2. Change directory to `core-migration`
3. Run `make run` on the root directory
4. Wait untill PostgreSQL is loaded and ready to receive connection
5. `Ctrl + C` and run `make migrate` to populate your database

## Contributing
1. Indentation: Use a strict 4-space indentation pattern (configure your tabs).
2. Table Naming: All table names must be plural (e.g., users, reimburses, job_slots).
3. Foreign Keys: Columns referencing external parent rows must be singular followed by _id (e.g., user_id, job_slot_id).
4. ENUM Layouts: Custom database type tokens/enums must be singular (e.g., user_role, sponsor_variant).

## Recommended Development Environment
- OS: Fedora, Debian
- Containers: Podman
- Text Editor: Zed Editor

## To Deploy
- Local port forwarding: `ssh -L`
- Run goose from local machine to vps: `goose -dir migrations postgres "postgres://username:password@127.0.0.1:port/dbname?sslmode=disable" up`
- Close port forwarding: `Ctrl + C`

## To Test
- Run `make run` to initialize PostgreSQL container, wait until psql is initialized
- Run `make migrate` to migrate the files from `migrations/` directory

## To Seed
- Run `go build ./cmd/seed`
- Then run the binary `./seed`

---
By BBJ Team
