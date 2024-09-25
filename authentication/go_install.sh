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
go install -tags 'grpc-gateway' github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
check_tool protoc-gen-grpc-gateway

echo "Install protoc-gen-openapiv2..."
go install -tags 'protoc-gen-openapiv2' github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
check_tool protoc-gen-openapiv2

echo "Install protoc-gen-go-grpc..."
go install -tags 'protoc-gen-go-grpc' google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
check_tool protoc-gen-go-grpc

echo "Install protoc-gen-go..."
go install -tags 'protoc-gen-go' google.golang.org/protobuf/cmd/protoc-gen-go@latest
check_tool protoc-gen-go


echo "========= Finish CLI tools installation ========="