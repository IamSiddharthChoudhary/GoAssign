# Go User Management API

A production-ready RESTful API built in **Go** to manage users with **name** and **date of birth (DOB)**, returning **age calculated dynamically** at request time.

This project demonstrates clean backend architecture using **Fiber**, **SQLC**, **PostgreSQL**, structured logging with **Zap**, and Docker-based deployment.

---

## Features

- Create, Read, Update, Delete users
- Store DOB in database, calculate age dynamically in Go
- Type-safe database access using SQLC
- Clean layered architecture (handler / service / repository)
- Input validation using `go-playground/validator`
- Structured logging with Uber Zap
- Request ID + request duration middleware
- Docker & Docker Compose support

---

## Tech Stack

- **Go 1.22+**
- **Fiber** – HTTP framework
- **PostgreSQL** – Database
- **SQLC** – Type-safe DB layer
- **pgx** – PostgreSQL driver
- **Zap** – Structured logging
- **Validator** – Input validation
- **Docker / Docker Compose**

---

## API Endpoints

### Create User

**POST** `/users`

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### Get User by ID

**GET** `/users/:id`

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 34
}
```

---

### Update User

**PUT** `/users/:id`

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

---

### Delete User

**DELETE** `/users/:id`

Returns **204 No Content**

---

### List Users

**GET** `/users`

```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 34
  }
]
```

---

## Age Calculation

Age is **not stored** in the database.
It is calculated dynamically in Go using the `time` package to ensure correctness and testability.

```go
func CalculateAge(dob time.Time) int
```

---

## Local Setup (Without Docker)

### Prerequisites

- Go 1.22+
- PostgreSQL
- sqlc

### Environment Variable

Set the database connection string **before running the app**.

#### Local PostgreSQL

````bash
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
```bash
export DATABASE_URL="postgres://user:password@localhost:5432/assignment?sslmode=disable"
````

### Create Database Tables (Required)

The application does **not** auto-create tables. You must run the migration once per database.

````bash
psql "$DATABASE_URL" -c "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT NOT NULL, dob DATE NOT NULL, created_at TIMESTAMPTZ DEFAULT now(), updated_at TIMESTAMPTZ DEFAULT now());"
```bash
psql $DATABASE_URL -f db/migrations/001_users.sql
````

### Generate SQLC Code

```bash
sqlc generate
```

### Run Server

```bash
go run cmd/server/main.go
```

---

## Docker Setup (Recommended)

### Environment variable (Local Postgres)

```bash
export DATABASE_URL="postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
```

### Start containers

```bash
docker compose up --build
```

### Create tables (first run only)

```bash
docker exec -it assignment_db psql -U postgres -d postgres \
  -c "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT NOT NULL, dob DATE NOT NULL, created_at TIMESTAMPTZ DEFAULT now(), updated_at TIMESTAMPTZ DEFAULT now());"
```

API will be available at:

```
http://localhost:8000
```

---

## Logging & Middleware

Each request logs:

- HTTP method
- Path
- Status code
- Latency
- Request ID

Example:

```
{"method":"GET","path":"/users/1","status":200,"latency":"2ms","request_id":"..."}
```

---

## Supabase Setup (Optional)

Supabase can be used instead of a local PostgreSQL database.

### Steps

1. Create a Supabase project
2. Go to **Settings → Database → Connection pooling**
3. Select **Session pooler** (IPv4 compatible)
4. Copy the connection URI

It will look like:

```text
postgresql://postgres:<PASSWORD>@aws-0-<region>.pooler.supabase.com:6543/postgres
```

### Set environment variable

```bash
export DATABASE_URL="postgresql://postgres:<PASSWORD>@aws-0-<region>.pooler.supabase.com:6543/postgres?sslmode=require"
```

### Create table (one-time)

Run this in **Supabase SQL Editor**:

```sql
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  dob DATE NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now()
);
```

After this, start the API (locally or via Docker) and test normally.

---

## Future Improvements

- Pagination for `/users`
- Unit tests for service layer
- JWT authentication
- Swagger / OpenAPI docs

---
