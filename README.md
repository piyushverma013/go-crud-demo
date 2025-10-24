# go-crud-demo

A simple **in-memory CRUD API** for users built with **Go** and **Gin**.  
This project demonstrates basic Go concepts like modules, interfaces, packages, and how to structure a small web API. Perfect for beginners learning Go web development.

## Features

- Create, Read, Update, Delete users
- In-memory storage (no database required)
- Uses Gin framework for routing
- Clean project structure:
  - `handler` – handles HTTP requests
  - `service` – business logic
  - `router` – sets up routes

## Installation


1. Clone the repository:
```
git clone =https://github.com/piyushverma013/go-crud-demo/go-crud-demo.git
cd go-crud-demo
```

2. Install dependencies:
```
go mod tidy
```

3. Run the server:
```
go run main.go
```

The API will start on http://localhost:8080.

| Method | Endpoint     | Description       |
| ------ | ------------ | ----------------- |
| POST   | `/users/`    | Create a new user |
| GET    | `/users/`    | Get all users     |
| GET    | `/users/:id` | Get user by ID    |
| PUT    | `/users/:id` | Update user by ID |
| DELETE | `/users/:id` | Delete user by ID |


**Example Requests**
Create User (POST /users/)
```
curl -X POST http://localhost:8080/users/ \
-H "Content-Type: application/json" \
-d '{"name":"John Doe","email":"john@example.com"}'
```

Get All Users (GET /users/)
```
curl http://localhost:8080/users/
```

Get User by ID (GET /users/1)
```
curl http://localhost:8080/users/1
```

Update User (PUT /users/1)
```
curl -X PUT http://localhost:8080/users/1 \
-H "Content-Type: application/json" \
-d '{"name":"Jane Doe","email":"jane@example.com"}'
```

Delete User (DELETE /users/1)
```
curl -X DELETE http://localhost:8080/users/1
```

