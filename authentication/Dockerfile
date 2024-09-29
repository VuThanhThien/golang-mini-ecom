# Build stage
FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o authentication-service .

# Final stage
FROM alpine:latest

# Add CA certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/authentication-service .

# Copy the app.env file
COPY --from=builder /app/app.env .

# Expose the port the app runs on
EXPOSE 8000

# Command to run the application
CMD ["./authentication-service"]