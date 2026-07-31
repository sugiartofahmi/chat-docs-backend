# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

- Install deps: `go mod download`
- Copy `.env-example` to `.env` before running anything (requires local PostgreSQL and Redis; OpenSearch is optional — skipped with a warning if `OPENSEARCH_HOST`/`OPENSEARCH_PASSWORD` are unset)
- Run migrations: `go run . --migration=true --exec=up` (also `--exec=down`, `--exec=fresh`, `--exec=create --fileName=<name>`)
- Auto-migrate from entity structs (dev shortcut, no SQL files): `go run . --automigrate=true`
- Seed data: `go run . --seed=true` (all seeders) or `go run . --seed=true --target=RoleSeeder,UserSeeder` (specific)
- Run server: `go run .` (listens on `APP_PORT`, default `8080`)
- Build: `go build -o go-service .`
- Test all: `go test ./...`
- Test one: `go test ./infrastructure/middlewares/ -run TestParseBindingErrors_InvalidJSON` (standard `go test -run`; only one test file exists in the repo today, so there's no established per-package testing convention to follow — use stdlib `testing`, no testify)

No linter or formatter is configured (no `golangci-lint`, no CI pipeline) — fall back to `go vet`/`gofmt`.

## Architecture

**Entry point** `main.go` runs one fixed pipeline: `extractArgs → initializeSingleton → runnerMigration → runnerAutoMigrate → runnerSeeder → initializeRouter → initializeRepositories → initializeServices → initializeControllers → initializeHttpServer`. Dependency wiring is entirely manual (no DI framework) — new repositories/services/controllers must be constructed and wired by hand in `main.go`. CLI flags (`--migration`, `--exec`, `--automigrate`, `--seed`, `--target`) short-circuit into one-off migration/seed runs and exit before the server starts.

**Directory layout**
- `domain/<feature>/` (`auth`, `role`, `user`) — `constants/`, `dtos/`, `interfaces/`, `repositories/`, `services/` per feature. Repositories are split into `*_query_repository.go` (reads) and `*_store_repository.go` (writes).
- `entities/` — GORM model structs shared across domains.
- `infrastructure/` — cross-cutting code: `config/` (env-var-backed constants, one file per concern), `databases/` (Postgres/OpenSearch connection constructors), `redis/`, `singleton/` (process-wide clients: `PostgresSingleton()`, `RedisSingleton()`, `OpensearchSingleton()`), `middlewares/`, `exceptions/`, `validators/`, `utils/`.
- `presentation/api/` — Gin controllers only, grouped `health/`, `v1/auth/`, `v1/role/`, `v1/user/`. Each controller owns its own `router.Group(...)` and route table — there is no central `routes.go`.
- `migration/` — `golang-migrate` SQL runner (`migration.go`, `files/*.sql`) plus a GORM `AutoMigrate` path (`automigrate.go`).
- `seeder/` — JSON-driven seeders (`seeder.go`, `files/*.json`).

**New feature modules** follow the `domain/auth` / `domain/role` / `domain/user` pattern: create `domain/<module>/{constants,dtos,interfaces,repositories,services}` (query/store split), add entities under `entities/`, add a controller under `presentation/api/v1/<module>/controller/`, then wire repository → service → controller by hand in `main.go`'s `initializeRepositories`/`initializeServices`/`initializeControllers`.

**Request flow** (traced via `POST /api/v1/users`):
1. Global middleware, attached in `initializeRouter`: `gin.Recovery()` → `LoggingMiddleware()` → CORS → `ExceptionMiddleware()`. `ExceptionMiddleware` installs a deferred `recover()` that turns panics anywhere downstream into structured JSON error responses — this is the primary error-handling mechanism in the app.
2. Per-route-group middleware (declared in the controller, not global): `AuthorizationMiddleware()` validates the `Authorization: Bearer <jwt>` header and stores claims in the Gin context; `RoleMiddleware([]string{...})` checks the caller's role and panics `ForbiddenException` if disallowed.
3. `middlewares.ValidateRequestJson[T]()` — a generic middleware that binds+validates the JSON body into `T` (via `go-playground/validator` struct tags) and stores the parsed DTO in context. This replaced older per-controller `Assign*RequestDto(c)` helper functions (see git log) — **query-string DTOs still use the old manual `Assign*QueryRequestDto(httpContext)` pattern**, and path params (`:id`) are validated manually via `validators.ValidateUUID(...)`; this split is intentional, not leftover inconsistency to "fix".
4. Controller pulls the DTO via `httpContext.MustGet(middlewares.RequestBodyJsonKey).(*dtos.XxxRequestDto)`, calls the service, wraps the result via a `dtos.XxxResponseDtoFromEntity` mapper, writes `utils.SuccessResponse`.
5. Service applies business rules and calls into the query/store repositories.
6. Repository does GORM calls against the entity, panicking `exceptions.ServerErrorException` on DB errors.

**Error handling is panic-based throughout services/repositories/middleware** (custom exception types like `ForbiddenException`, `ServerErrorException`), not Go's usual returned-`error` convention — `ExceptionMiddleware`'s recover is what turns these into responses.

**Storage**: PostgreSQL via GORM is the primary DB. Redis backs a cache-service abstraction (`infrastructure/redis/services/redis_cache_service.go`). An OpenSearch client/connection exists (`infrastructure/databases/opensearch.go`, `infrastructure/config/opensearch_config.go`) but as of now nothing in `domain/` consumes it — it's connection scaffolding for future search features, and connecting to it is non-fatal if unconfigured (unlike Postgres/Redis, which panic on connection failure).

**Config**: custom loader, no viper. `infrastructure/config/config.go`'s `LoadConfig()` calls `godotenv.Load()` once; `Get(key, default)` / `GetRequired(key)` read via `os.Getenv`. Note `GetRequired` does **not** panic on a missing var — it returns `"<key>-not-set"` as a sentinel string, so a missing required env var fails silently/late rather than at startup.
