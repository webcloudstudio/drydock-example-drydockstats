#!/bin/sh
# CommandCenter Operation
# Name: Run Tests
# Type: batch

# Vet and test every package in the module. Exit status is the verdict.
set -eu
PROJECT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_DIR"

go vet ./...
go test ./...
