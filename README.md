# GOFIBER STARTERKIT MYSQL
ready setup

## Project Structure

```bash
├── cmd
│   └── main.go
├── config
│   ├── app.go
│   ├── environtment.go
│   ├── fiber.go
│   ├── logrus.go
│   ├── mysql.go
│   └── validator.go
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   └── example
│   │       ├── dto
│   │       ├── entity
│   │       ├── handler
│   │       │   └── example_handler.go
│   │       ├── repository
│   │       │   └── resource
│   │       └── service
│   └── infrastructure
│       └── rest
│           ├── middleware
│           └── routes
│               └── route.go
├── pkg
│   └── response_status_exception.go
└── test
```

## ENV 
build file `.env`

```bash
APP_NAME="your app"

DB_MYSQL_URL="username:password@tcp(localhost:3306)/yourdb?charset=utf8mb4&parseTime=True&loc=Local"
DB_POOL_MAX_IDLE_CONNS=5
DB_POOL_MAX_OPEN_CONNS=15
DB_POOL_MAX_IDLE_TIME=10 # minute
DB_POOL_MAX_LIFETIME=30 # minute
```