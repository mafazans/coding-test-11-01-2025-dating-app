# Go Dating App

This is a RESTful Dating App built with Go, Gin, and PostgreSQL. It provides user registration and login functionality with JWT authentication.

## Prerequisites

To run this project you need to have the following installed:

- [Go] Go 1.23 or higher
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
│   │   └── handler_test.go
│   │   └── profile.go
│   │   └── server.go
│   │   └── subscription.go
│   ├── middleware/
│   │   └── auth.go
│   ├── model/
│   │   └── profile.go
│   │   └── subscription.go
│   │   └── swipe.go
│   │   └── user.go
│   └── repository/
│       └── implementation.go
│       └── interfaces.go
│       └── interfaces.mock.gen.go
│       └── profile.go
│       └── repository.go
├── tests/
│   └── coverage.out
├── database.sql
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md

```

## Setup Instructions

1. Clone the repository:
```bash
git clone https://github.com/mafazans/coding-test-11-01-2025-dating-app.git
cd coding-test-11-01-2025-dating-app.git
```

2. Install dependencies (For local development):
```bash
go mod tidy
```

3. Run app in the docker
```bash
docker compose up -d --build
```

4. Run test & coverage
```bash
	go clean -testcache
	go test -coverprofile tests/coverage.out -short -v ./...
```

## API Endpoints

You can use curl or Postman to test the endpoints:

```bash
# Register a new user
curl --location 'http://localhost:8080/api/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test1@gmail.com",
    "password": "password"
}'

# Login
curl --location 'http://localhost:8080/api/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test1@gmail.com",
    "password": "password"
}'

# Get Profiles
curl --location 'http://localhost:8080/api/profiles?limit=30' \
--header 'Authorization: Bearer {{token}}'

# Verify Profile (premium)
curl --location 'http://localhost:8080/admin/verify-profile' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 21
}'

#Swipe
curl --location 'http://localhost:8080/api/swipe' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY3NjMzMzQsInVzZXJfaWQiOjIxfQ.Zz1N26IHTy8iss_vo2zGf8bIFsXtiYyrxVyYC3uYUUE' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 20,
    "is_like": true
}'
```

# Postman Collection
```
{
	"info": {
		"_postman_id": "b938ef4a-b2e5-447a-96ed-de763f3dbb43",
		"name": "dealls",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Register",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Authorization",
						"value": "",
						"type": "default",
						"disabled": true
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"username\": \"test1@gmail.com\",\n    \"password\": \"passwordsuperusefull\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://localhost:8080/api/register",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"register"
					]
				}
			},
			"response": []
		},
		{
			"name": "Login",
			"event": [
				{
					"listen": "prerequest",
					"script": {
						"exec": [
							""
						],
						"type": "text/javascript"
					}
				},
				{
					"listen": "test",
					"script": {
						"exec": [
							"// Parse the JSON response",
							"var jsonData = pm.response.json();",
							"",
							"// Check if the token exists in the response",
							"if (jsonData.token) {",
							"    // Set the token as an environment variable",
							"    pm.environment.set(\"token\", jsonData.token);",
							"} else {",
							"    console.error(\"Token not found in the response\");",
							"}"
						],
						"type": "text/javascript"
					}
				}
			],
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"username\": \"test1@gmail.com\",\n    \"password\": \"passwordsuperusefull\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://localhost:8080/api/login",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"login"
					]
				}
			},
			"response": []
		},
		{
			"name": "Get Profile",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/api/login",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"login"
					]
				}
			},
			"response": []
		},
		{
			"name": "Verify Profile",
			"request": {
				"method": "POST",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/admin/generate-profiles",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"admin",
						"generate-profiles"
					]
				}
			},
			"response": []
		},
		{
			"name": "Swipe",
			"request": {
				"method": "POST",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/admin/generate-profiles",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"admin",
						"generate-profiles"
					]
				}
			},
			"response": []
		}
	]
}
```