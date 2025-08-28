# Makefile for stellar Sandbox

APP_NAME=stellar-x
DOCKER_COMPOSE=docker compose

# Auto migrate models to create database table
migrate:
	go build -o stellar_x_db ./cmd/migrate/main.go

# Build sandbox
build:
	go build -o $(APP_NAME) ./cmd/server/main.go
	@echo "✅ Binary $(APP_NAME) built successfully."


# Run in dev mode
dev:
	docker compose up

# Run in prod mode
prod:
	docker compose up --build

# Clean postgresql volume
cleandb:
	stop
	docker volume rm stellar_x_db

# Stop container
stop:
	$(DOCKER_COMPOSE) down


# Clean container and Images
clean:
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans
	@echo "🧹 Cleaned containers, images, and volumes."

# Restart
restart: stop prod

# Log
logs:
	$(DOCKER_COMPOSE) logs -f

# lint
lint:
	~/go/bin/golangci-lint run ./...

# Format Go code
fmt:
	go fmt ./...
	go mod tidy

# Run the sandbox on local host
run:
	go run ./cmd/server