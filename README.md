# CRUD API with Go, Gin, GORM, and MySQL

This project implements a simple CRUD (Create, Read, Update, Delete) API for managing books, built with Go, the Gin web framework, GORM (an ORM for Go), and a MySQL database.

## Features

- **Create Book**: Add new book records to the database.
- **Get All Books**: Retrieve a list of all books.
- **Get Book by ID**: Retrieve a single book record by its unique ID.
- **Delete Book**: Remove a book record from the database.
- **MySQL Integration**: Uses GORM to interact with a MySQL database.

## Technologies Used

- **Go**: The programming language.
- **Gin Web Framework**: A high-performance HTTP web framework.
- **GORM**: An ORM library for Go, simplifying database interactions.
- **MySQL**: The relational database used to store book information.

## Project Structure

```
crudapiwithsql/
├── cmd/
│   └── main/
│       └── main.go             # Main application entry point
├── pkg/
│   ├── config/
│   │   └── app.go              # Database connection and configuration
│   ├── controllers/
│   │   └── controller.go       # Handles API logic and interacts with models
│   ├── models/
│   │   └── book.go             # Defines the Book model and database operations
│   ├── routes/
│   │   └── bookstore-routes.go # Defines API endpoints
│   └── utils/
│       └── utils.go            # Utility functions for JSON marshalling/unmarshalling
└── go.mod                      # Go module definition
└── go.sum                      # Go module checksums
└── README.md                   # Project README
```

## Setup and Installation

### Prerequisites

- Go (version 1.16 or higher)
- MySQL Server
- Git

### Steps

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/vaddepally-shiva-shankar/go-crudapi-sql.git
    cd go-crudapi-sql/crudapiwithsql
    ```

2.  **Configure MySQL Database:**
    - Create a MySQL database (e.g., `gocrudapisql`).
    - Update the database connection string in `pkg/config/app.go` with your MySQL credentials:
        ```go
        // pkg/config/app.go
        func Connect() {
            d, err := gorm.Open("mysql", "shiva:yourpassword@/gocrudapisql?charset=utf8&parseTime=True&loc=Local")
            if err != nil {
                panic(err)
            }
            db = d
            fmt.Println("Connected to database successfully")
            log.Println("Connected to database successfully")
        }
        ```
        Replace `shiva:yourpassword` with your MySQL username and password.

3.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```

4.  **Run the application:**
    ```bash
    go run cmd/main/main.go
    ```
    The API will start on `http://localhost:8080`.

## API Endpoints

The API exposes the following endpoints:

### Books

-   **`GET /book`**
    -   Retrieves all books.
    -   **Response:** `200 OK` with an array of book objects.

-   **`GET /book/{id}`**
    -   Retrieves a single book by its ID.
    -   **Parameters:** `id` (path parameter, integer)
    -   **Response:** `200 OK` with a book object, or `404 Not Found` if the book does not exist.

-   **`POST /book`**
    -   Creates a new book.
    -   **Request Body (JSON):**
        ```json
        {
            "name": "The Hitchhiker's Guide to the Galaxy",
            "author": "Douglas Adams",
            "publication": "Pan Books"
        }
        ```
    -   **Response:** `200 OK` with the created book object.

-   **`DELETE /book/{id}`**
    -   Deletes a book by its ID.
    -   **Parameters:** `id` (path parameter, integer)
    -   **Response:** `200 OK` with the deleted book object, or `404 Not Found` if the book does not exist.

## Example Usage (using `curl`)

### Create a Book

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "The Go Programming Language",
    "author": "Alan A. A. Donovan, Brian W. Kernighan",
    "publication": "Addison-Wesley Professional"
}' http://localhost:8080/book
```

### Get All Books

```bash
curl http://localhost:8080/book
```

### Get Book by ID

```bash
curl http://localhost:8080/book/1
```

### Delete Book

```bash
curl -X DELETE http://localhost:8080/book/1
