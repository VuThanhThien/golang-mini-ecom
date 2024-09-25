#!/bin/sh

echo "========= Install CLI tools ========="

check_tool() {
    if ! command -v $1 &> /dev/null
    then
        echo "$1 not found"
    else
        echo "$1 has been installed at $(command -v $1)"
    fi
}

echo "Install golangci-lint..."
go install -tags 'lint' github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
check_tool golangci-lint

echo "Install golang-migrate..."
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
check_tool migrate

echo "Install protoc-gen-grpc-gateway..."
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
check_tool protoc-gen-grpc-gateway

echo "Install protoc-gen-grpc-gateway..."
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
check_tool protoc-gen-grpc-gateway

echo "Install protoc-gen-go..."
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
check_tool protoc-gen-go

echo "Install protoc-gen-go-grpc..."
go install google.golang.org/protobuf/cmd/protoc-gen-go
check_tool protoc-gen-go-grpc


echo "========= Finish CLI tools installation ========="