# Task Manager API

A RESTful API for managing personal tasks. Users can register, authenticate, and manage their own tasks — create, read, update, and delete. Built as a portfolio project demonstrating production-ready Go backend development practices.

## Tech Stack

- **Go** — primary language
- **Gin** — HTTP framework
- **PostgreSQL** — database
- **sqlx** — database interaction
- **JWT** — authentication (access + refresh tokens)
- **Docker** — containerization
- **golang-migrate** — database migrations

## Getting Started

**1. Clone the repository**
```bash
git clone https://github.com/aidostt/task-manager.git
cd task-manager
```

**2. Create environment file**
```bash
cp .env.example .env
```
Fill in the values in `.env` (see `.env.example` for required fields).

**3. Start the database**
```bash
make docker-up
```

**4. Run migrations**
```bash
make migrate-up
```

**5. Start the server**
```bash
make run
```

Server will start at `http://localhost:8080`.

## Environment Variables

```env
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=taskmanager
DB_HOST=localhost
DB_PORT=5432

JWT_SECRET=your_secret_key
ACCESS_TOKEN_TTL=15m
REFRESH_TOKEN_TTL=168h

SERVER_HOST=localhost
SERVER_PORT=8080
```

## API Endpoints

### Auth

| Method | Path | Description |
|--------|------|-------------|
| POST | `/auth/register` | Register a new user |
| POST | `/auth/login` | Login and receive tokens |
| POST | `/auth/refresh` | Refresh access token |

### Tasks

> All task endpoints require `Authorization: Bearer <token>` header.

| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/tasks/` | Create a new task |
| GET | `/api/tasks/` | Get all tasks for the authenticated user |
| GET | `/api/tasks/:id` | Get a task by ID |
| PUT | `/api/tasks/:id` | Update a task by ID |
| DELETE | `/api/tasks/:id` | Delete a task by ID |

### Task fields

| Field | Type | Values |
|-------|------|--------|
| `title` | string | required, max 255 chars |
| `description` | string | optional, max 500 chars |
| `status` | string | `todo`, `in_progress`, `done` |
| `priority` | string | `low`, `medium`, `high` |

## Examples

**Register**
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'
```

**Login**
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'
```

**Save token to variable**
```bash
TOKEN=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}' \
  | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
```

**Create task**
```bash
curl -X POST http://localhost:8080/api/tasks/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy milk", "description": "from store", "status": "todo", "priority": "low"}'
```

**Get all tasks**
```bash
curl -X GET http://localhost:8080/api/tasks/ \
  -H "Authorization: Bearer $TOKEN"
```

**Get task by ID**
```bash
curl -X GET http://localhost:8080/api/tasks/{task_id} \
  -H "Authorization: Bearer $TOKEN"
```

**Update task**
```bash
curl -X PUT http://localhost:8080/api/tasks/{task_id} \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy bread", "description": "from bakery", "status": "in_progress", "priority": "high"}'
```

**Delete task**
```bash
curl -X DELETE http://localhost:8080/api/tasks/{task_id} \
  -H "Authorization: Bearer $TOKEN"
```

**Refresh token**
```bash
curl -X POST http://localhost:8080/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{"refresh_token": "your_refresh_token"}'
```

## Project Structure

```
task-manager/
├── cmd/api/          # entry point
├── internal/
│   ├── config/       # configuration
│   ├── handler/      # HTTP handlers and middleware
│   ├── model/        # data models and input structs
│   ├── repository/   # database layer
│   └── service/      # business logic
├── migrations/       # SQL migrations
├── pkg/
│   ├── db/           # database connection
│   ├── jwt/          # JWT token manager
│   └── server/       # HTTP server
├── .env.example
├── docker-compose.yml
└── Makefile
```
