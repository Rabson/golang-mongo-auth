# Golang MongoDB Authentication API

This project is a backend API built with Golang, MongoDB, and Gin framework. It provides user authentication with JWT tokens and follows OWASP security best practices.

## Features
- User registration with hashed passwords
- User login with JWT authentication
- Middleware for protected routes
- User profile retrieval
- Secure secret management using `.env`

## Prerequisites
- Go 1.18 or later
- MongoDB installed and running
- Environment variables configured in `.env`

## Setup Instructions

### 1. Clone the Repository
```sh
git clone https://github.com/your-repo/golang-mongo-auth.git
cd golang-mongo-auth
```

### 2. Create `.env` File
Create a `.env` file in the root directory and add the following content:
```sh
MONGO_URI=mongodb://localhost:27017
JWT_SECRET=your-secret-key
```

### 3. Install Dependencies
```sh
go mod tidy
```

### 4. Run the Application
```sh
go run main.go
```

## How to Run the Code

1. **Ensure MongoDB is running**
   ```sh
   mongod --dbpath /path/to/your/db
   ```

2. **Export environment variables** (if not using `.env` file)
   ```sh
   export MONGO_URI=mongodb://localhost:27017
   export JWT_SECRET=your-secret-key
   ```

3. **Run the application**
   ```sh
   go run main.go
   ```

4. **Test the API using Postman or curl**
   ```sh
   curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"email": "user@example.com", "password": "securepassword", "name": "John Doe"}'
   ```

### 5. API Endpoints
#### Register User
```sh
POST /register
{
    "email": "user@example.com",
    "password": "securepassword",
    "name": "John Doe"
}
```

#### Login User
```sh
POST /login
{
    "email": "user@example.com",
    "password": "securepassword"
}
```
Response:
```json
{
    "token": "your-jwt-token"
}
```

#### Get User Profile (Protected Route)
```sh
GET /profile
Authorization: Bearer your-jwt-token
```


## Migration 
```
migrate create -ext -mongodb  -dir migration  <NAME>

```

## Installed dependencies
```

go get -u github.com/cosmtrek/air
go get github.com/gin-gonic/gin 
go get github.com/golang-jwt/jwt/v5 
go get go.mongodb.org/mongo-driver 
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv 
go get github.com/golang-migrate/migrate/v4
go get github.com/gin-contrib/cors

```

## Seed
make seed module=<seed_cmd_name>

## Notes
- Ensure MongoDB is running before starting the server.
- Keep your `.env` file secure and never commit it to version control.

## License
MIT License