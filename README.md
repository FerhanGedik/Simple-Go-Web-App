# Simple Go Web Application - Gorm & Gorilla Mux & SQLite

This is a simple Go project using Gorilla Mux and Gorm to create a RESTful API for managing tasks.

## Getting Started

### Prerequisites

Make sure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).

### Installing Dependencies

```bash
go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

## Running the Application

1. Clone the repository:

```bash
git clone https://github.com/ferhangedik/simple-go-web-app.git
```
2.Change into the project directory:
```bash
cd simple-go-web-app
```
3.Run the application:
```bash
go run main.go
```
The server will be running at http://localhost:8080.


# API Endpoints

## Get all tasks:

- **Endpoint:** `/tasks`
- **Method:** GET

## Get a task by ID:

- **Endpoint:** `/tasks/{id}`
- **Method:** GET

## Create a new task:

- **Endpoint:** `/tasks`
- **Method:** POST

## Update a task by ID:

- **Endpoint:** `/tasks/{id}`
- **Method:** PUT

## Delete a task by ID:

- **Endpoint:** `/tasks/{id}`
- **Method:** DELETE

# Serving Static Files

Static files, such as documentation, can be served from the `/static/` endpoint.

# Contributing

Feel free to contribute to this project by opening issues or pull requests. If you encounter any problems or have suggestions for improvements, please let us know!


