include app.env
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOLINT=golangci-lint
DATABASE_URL := "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${SSL_MODE}&timezone=${DB_TIMEZONE}"
GONETWORK=beego_network

# Main package path
MAIN_PATH=./cmd/server
MIGRATIONS_PATH=internal/initializers/migrations
step = 1

# Binary name
BINARY_NAME=ale-ecom

install:
	sh go_install.sh

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)

# Run the project
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)
	./$(BINARY_NAME)

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run linter
lint:
	$(GOLINT) run

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Download dependencies
deps:
	$(GOGET) -v -t -d ./...
	$(GOMOD) tidy

# Update dependencies
update-deps:
	$(GOGET) -u -v -t -d ./...
	$(GOMOD) tidy

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)

# create docker network
network:
	docker network create $(GONETWORK)

# create image postgres
postgres:
	docker run --name ${BINARY_NAME} --network $(GONETWORK) -p ${POSTGRES_PORT}:${POSTGRES_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_DB=${POSTGRES_DB} -d postgres:14-alpine

# start db
startdb:
	docker container start ${BINARY_NAME}

# run dev
dev:
	go run $(MAIN_PATH)/main.go

# build swagger
docs:
	swag init -g $(MAIN_PATH)/main.go --parseDependency --parseInternal

migration-up:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) up

migration-down:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) down

migration-new:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

.PHONY: migrate dev docs postgres network startdb build run clean test test-coverage deps update-deps build-all lint install migration-up migration-down
