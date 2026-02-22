# Practice 3 (Go + PostgreSQL)

## 1) Prepare DB
- Start Postgres (Postgres.app).
- Create DB in psql:
  - `CREATE DATABASE mydb;`

## 2) Run server
From project root:
- `go mod tidy`
- `go run ./cmd/api`

Server starts on `:8080`.

## 3) Test endpoints (API KEY required)
Health:
- `curl -i http://localhost:8080/health`

Users (with header):
- `curl -i -H "X-API-KEY: my-secret-key" http://localhost:8080/users`

Create:
- `curl -i -H "X-API-KEY: my-secret-key" -H "Content-Type: application/json" -d '{"name":"Alice","email":"alice@mail.com","age":22}' http://localhost:8080/users`

Get by id:
- `curl -i -H "X-API-KEY: my-secret-key" http://localhost:8080/users/1`

Update:
- `curl -i -X PUT -H "X-API-KEY: my-secret-key" -H "Content-Type: application/json" -d '{"name":"Alice Updated","email":"alice@mail.com","age":23}' http://localhost:8080/users/1`

Delete:
- `curl -i -X DELETE -H "X-API-KEY: my-secret-key" http://localhost:8080/users/1`
