
# Go Key-Value Store (go-kv)

A lightweight, in-memory key-value database implemented in Go with support for basic CRUD operations over HTTP. This project demonstrates a thread-safe key-value store that can handle parallel HTTP requests without race conditions.

## Features

- **In-memory store**: The key-value store is implemented using `sync.Map`, ensuring concurrent access with no race conditions.
- **CRUD operations**: Supports `GET`, `PUT`, `DELETE`, and `LIST` operations via HTTP.
- **Parallelism**: The server is designed to handle concurrent requests safely, with tests to verify the absence of race conditions.
- **Unit-tested**: Comprehensive unit tests for all endpoints, including edge cases and parallelism testing.

## Checklist

| Requirement                                  | Status               | Notes |
|----------------------------------------------|----------------------|-------|
| `GET /[key]` returns value or 404            | ✅ Implemented       | Tested with valid/missing key cases. |
| `PUT /[key]` adds/updates key value          | ✅ Implemented       | Tested for adding/updating key values, including invalid inputs. |
| `DELETE /[key]` deletes key or returns 404   | ✅ Implemented       | Tested with valid/missing key cases. |
| `GET /` returns all keys as JSON array       | ✅ Implemented       | Tested for correct key listing. |
| Parallelism (no race conditions)             | ✅ Tested            | Application-level parallelism tests (concurrent access to the store). |
| In-memory store (no persistence)             | ✅ Implemented       | Using `sync.Map` for in-memory storage. |
| Unit tests for each endpoint                 | ✅ Implemented       | Unit tests for all handlers, including edge cases and parallelism. |

## Endpoints

### 1. GET /[key]
Retrieve the value associated with a given key.

- **Response**:
  - `200 OK` with the value if the key exists.
  - `404 Not Found` if the key does not exist.

#### Example:
```
curl http://localhost:8080/key/mykey
```

### 2. PUT /[key]
Store or update the value for a given key.

- **Request Body**: The new value to associate with the key.
- **Response**:
  - `200 OK` on success.
  - `400 Bad Request` if the request body is invalid or missing.

#### Example:
```
curl -X PUT -d 'myvalue' http://localhost:8080/key/mykey
```

### 3. DELETE /[key]
Delete the value associated with a given key.

- **Response**:
  - `200 OK` if the key was deleted.
  - `404 Not Found` if the key does not exist.

#### Example:
```
curl -X DELETE http://localhost:8080/key/mykey
```

### 4. GET /
Retrieve a list of all keys stored in the key-value database.

- **Response**:
  - `200 OK` with a JSON array of keys.
  - `200 OK` with an empty array if no keys exist.

#### Example:
```
curl http://localhost:8080/
```

## Running the Server

To start the server, run the following command:

```
go run main.go
```

The server will start on port `8080`. You can interact with the key-value store using the HTTP methods described above.

## Testing

This project includes unit tests for each endpoint and parallelism testing at the application level.

To run the tests, use:

```
go test ./... -v
```

To run the tests with the race detector (recommended for checking race conditions):

```
go test -race ./... -v
```

## Project Structure

- **application**: Contains the core business logic (key-value store).
- **web**: Contains HTTP handlers to map the HTTP requests to the key-value store logic.
- **main.go**: Entry point to the application, where the server is set up and started.
