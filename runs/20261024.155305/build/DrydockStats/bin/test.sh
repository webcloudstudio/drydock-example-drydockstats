#!/bin/sh
# DrydockStats test runner: runs the Go test suite from the application root.
set -eu

cd "$(dirname "$0")/.."
go vet ./...
go test ./...
