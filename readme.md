# To-Do API

A small Go REST API for managing to-do items. The service supports creating, updating, fetching, listing, and deleting todos through HTTP endpoints.

## Tech Stack

- Go `1.25.0`
- `gorilla/mux` for routing
- `google/uuid` for todo identifiers
- In-memory storage using Go maps

## Project Structure

```text
cmd/todo/main.go                 # application entrypoint
internal/api/app                 # application wiring
internal/api/server              # HTTP server setup
internal/api/routes              # route registration
internal/api/handler             # request validation and HTTP responses
internal/api/service             # business logic
internal/api/repository          # in-memory data storage
internal/api/models              # request/response/domain models
internal/api/middlewares         # request logging middleware
```

## Features

- Create a new todo with a unique title
- Update a todo title and completion status
- Fetch a single todo by ID
- List all todos
- Delete a todo by ID
- Log incoming HTTP requests

## How It Works

The application follows a simple layered structure:

1. `handler` validates request data and returns JSON responses.
2. `service` applies business rules.
3. `repository` stores todos in memory.

Because storage is in memory, all data is lost when the server stops.

## Running the Project

### Prerequisites

- Go `1.25.0` or newer available locally

### Start the API

```bash
go run ./cmd/todo
```

The server starts on `http://localhost:8080`.

## API Endpoints

### Create Todo

`POST /todos`

Request body:

```json
{
  "title": "Buy groceries"
}
```

Success response: `201 Created`

```json
{
  "id": "3f1f4e0a-1a78-4b18-b8cf-4774f0d88d89",
  "title": "Buy groceries",
  "is_completed": false
}
```

### Get All Todos

`GET /todos`

Success response: `200 OK`

```json
[
  {
    "id": "3f1f4e0a-1a78-4b18-b8cf-4774f0d88d89",
    "title": "Buy groceries",
    "is_completed": false
  }
]
```

### Get Todo By ID

`GET /todos/{id}`

Success response: `200 OK`

```json
{
  "id": "3f1f4e0a-1a78-4b18-b8cf-4774f0d88d89",
  "title": "Buy groceries",
  "is_completed": false
}
```

### Update Todo

`PUT /todos/{id}`

Request body:

```json
{
  "previous_title": "Buy groceries",
  "updated_title": "Buy groceries and fruits",
  "is_completed": true
}
```

Success response: `200 OK`

```json
{
  "id": "3f1f4e0a-1a78-4b18-b8cf-4774f0d88d89",
  "title": "Buy groceries and fruits",
  "is_completed": true
}
```

Note: `previous_title` is optional in the request. If it is omitted, the handler loads the current todo and fills it automatically before calling the service layer.

### Delete Todo

`DELETE /todos/{id}`

Success response: `200 OK`

```json
{
  "message": "todo deleted successfully"
}
```

## Example cURL Commands

Create a todo:

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries"}'
```

List todos:

```bash
curl http://localhost:8080/todos
```

Get a todo by ID:

```bash
curl http://localhost:8080/todos/<todo-id>
```

Update a todo:

```bash
curl -X PUT http://localhost:8080/todos/<todo-id> \
  -H "Content-Type: application/json" \
  -d '{"updated_title":"Buy groceries and fruits","is_completed":true}'
```

Delete a todo:

```bash
curl -X DELETE http://localhost:8080/todos/<todo-id>
```

## Validation and Error Behavior

- `title` is required when creating a todo
- `updated_title` is required when updating a todo
- `is_completed` is required when updating a todo
- Todo titles must be unique
- Invalid UUIDs return `400 Bad Request`
- Missing todos return `404 Not Found` for fetch and delete operations

Example error response:

```json
{
  "error": "todo not found"
}
```

## Current Limitations

- Data is not persisted to a database
- There is no pagination or filtering for `GET /todos`
- Todos are not safe across restarts
- The project does not currently include automated tests

## Future Improvements

- Add persistent storage such as PostgreSQL or SQLite
- Add unit and integration tests
- Support pagination, search, and filtering
- Add graceful shutdown and configuration via environment variables
