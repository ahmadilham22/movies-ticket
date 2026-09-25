# Movies Ticketing

A movie booking REST API built with Go, Gin, and PostgreSQL. Purchases and cancellations use database transactions and row-level locking to protect showtime quotas.

## Features

- JWT authentication, bcrypt password hashing, and admin authorization.
- Movie catalog with multiple genres.
- Admin-managed showtimes, prices, and quotas.
- Ticket purchases, booking history, and owner-only cancellation.

**Stack:** Go 1.26.0, Gin, sqlx, pgx, PostgreSQL 15, Docker Compose.
**Architecture:** Handler -> Service -> Repository -> Database.

## Setup

Requires Git, Go 1.26.0, Docker Compose v2, and a SQL client such as TablePlus. The API runs locally; Docker runs PostgreSQL.

### 1. Clone and Configure

```sh
git clone https://github.com/ahmadilham22/movies-ticket.git
cd movies-ticket
go mod download
```

Create `.env` from `.env.example` without overwriting an existing configuration:

```dotenv
PORT=3000
DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=12345
DB_NAME=ticket_db
SECRET_KEY=REPLACE_WITH_A_LONG_RANDOM_SECRET
```

Generate your own random `SECRET_KEY` before starting. The database password above is for the supplied local Docker setup only. Never commit `.env` or tokens.

### 2. Start the Database

```sh
docker compose up -d postgres-db
docker exec ticket_db_container pg_isready -U postgres -d ticket_db
```

Wait for PostgreSQL to accept connections. Compose creates `ticket_db` on first initialization but does not apply migrations.

### 3. Apply Migrations

In TablePlus, connect to `localhost:5433`, database `ticket_db`, using the credentials above. On an **empty database**, execute each complete file in order:

```text
migrations/0000_initial_schema.sql
migrations/0001_add_unique_booking_code.sql
migrations/0002_create_movie_catalog.sql
migrations/0003_add_ticket_showtimes.sql
migrations/0004_add_user_role.sql
```

Each file includes `BEGIN` and `COMMIT`. Stop on errors and run `ROLLBACK;` in the same connection. There is no migration history tracker: do not replay completed files. Do not use the legacy `table2.sql` for setup.

### 4. Run the API

```sh
go run ./cmd/api
```

Try `GET http://localhost:3000/movies`. A fresh database returns `200 OK` with an empty `data` array.

## Authentication

Register through `POST /users` with `name`, `email`, `password`, and `phone_number`. New accounts receive the `user` role. Log in through `POST /login` with `email` and `password`; the response's `data` field contains the JWT string.

Use `Authorization: Bearer <JWT>` for protected requests. Tokens expire after 24 hours.

For a local admin, register normally, find that account's UUID in TablePlus, and replace `<USER_UUID>` below:

```sql
UPDATE users SET role = 'admin' WHERE id = '<USER_UUID>';
```

Log in again after changing roles. Existing tokens are not automatically revoked.

## Endpoints

| Method | Path | Access | Purpose |
| --- | --- | --- | --- |
| POST | `/users` | Public | Register |
| POST | `/login` | Public | Log in |
| GET | `/genres` | Public | List genres |
| GET | `/movies` | Public | List movies |
| GET | `/movies/:id` | Public | Movie details |
| GET | `/tickets` | Public | List showtimes and quotas |
| POST | `/genres` | Admin | Create genre |
| POST | `/movies` | Admin | Create movie |
| POST | `/tickets/create` | Admin | Create showtime |
| GET | `/users` | Admin | List users |
| POST | `/tickets` | Authenticated | Purchase tickets |
| GET | `/transactions` | Authenticated | Own booking history |
| PATCH | `/transactions/:bookingCode/cancel` | Booking owner | Cancel booking |

A `ticket_id` identifies a showtime, not an individual seat. Purchase with `ticket_id` and a positive `quantity`; retrieve the booking code from `GET /transactions`. Cancellation restores quota once.

## Testing

```sh
go build ./...
go test ./...
```

Automated tests currently cover cancellation service error mapping. Database, HTTP, and concurrency checks were performed manually, including a quota-1 test with two purchase requests: one succeeded, one returned `409`, and the final quota was 0.

## Scope

Portfolio MVP with simulated purchases, no payment gateway or seat selection. Registration validation, email uniqueness, total-price limits, and deployment security still need hardening before production use.
