# go-simple-http-server

A simple http server for managing students, built with Go and `gorilla/mux`.

## Requirements

  * Go (1.18+ recommended)

## How to Run

1.  **Clone the repository** (or just have your project code ready)

2.  **Initialize Go Modules** (if you haven't):

    ```bash
    go mod init go-simple-http-server
    ```

3.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

4.  **Run the server:**

    ```bash
    go run main.go
    ```

    The server will start on `http://localhost:8080`.

## How to Test

To run the unit and integration tests for the API, navigate to the project's root directory and run:

```bash
go test -v ./...
```

## Available Endpoints

### GET /students

Returns a JSON list of all students.

### GET /students/{id}

Returns a single student by their ID.