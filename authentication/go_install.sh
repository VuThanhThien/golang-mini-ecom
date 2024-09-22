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


echo "========= Finish CLI tools installation ========="