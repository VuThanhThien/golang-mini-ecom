include app.env
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOLINT=golangci-lint
DATABASE_URL := "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?timezone=${DB_TIMEZONE}"

# Main package path
MAIN_PATH=./cmd/server

# Binary name
BINARY_NAME=ecombase

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
	docker network create beego_network

# create image postgres
postgres:
	docker run --name ${POSTGRES_DB} --network beego_network -p ${DB_PORT}:${DB_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:14-alpine

# start db
startdb:
	docker container start ${POSTGRES_DB}

# run dev
dev:
	go run $(MAIN_PATH)/main.go

# build swagger
docs:
	swag init -g $(MAIN_PATH)/main.go --parseDependency --parseInternal

.PHONY: migrate dev docs postgres network startdb build run clean test test-coverage deps update-deps build-all lint
