# Go Dating App

This is a RESTful Dating App built with Go, Gin, and PostgreSQL. It provides user registration and login functionality with JWT authentication.

## Prerequisites

To run this project you need to have the following installed:

- [Go] Go 1.23 or higher
- [GNU Make](https://www.gnu.org/software/make/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Project Structure

```
auth-service/
├── cmd/
│   └── main.go
├── internal/
│   ├── handler/
│   │   └── auth.go
│   ├── middleware/
│   │   └── auth.go
│   ├── model/
│   │   └── user.go
│   └── repository/
│       └── user.go
├── database.sql
├── .env
├── go.mod
└── README.md
```

## Setup Instructions

1. Clone the repository:
```bash
git clone https://github.com/mafazans/coding-test-11-01-2025-dating-app.git
cd coding-test-11-01-2025-dating-app.git
```

2. Install dependencies:
```bash
go mod tidy
```

## API Endpoints

### Register User
```
POST /api/register
Content-Type: application/json

{
    "username": "user@example.com",
    "password": "securepassword"
}
```

### Login
```
POST /api/login
Content-Type: application/json

{
    "username": "user@example.com",
    "password": "securepassword"
}
```

### Protected Route Example
```
GET /api/profile
Authorization: Bearer <your-jwt-token>
```

## Testing the API

You can use curl or Postman to test the endpoints:

```bash
# Register a new user
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"user@example.com","password":"securepassword"}'

# Login
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"user@example.com","password":"securepassword"}'

# Access protected route
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer <your-jwt-token>"
```
