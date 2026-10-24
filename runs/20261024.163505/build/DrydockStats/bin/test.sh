#!/bin/sh
# test.sh — runs the project's own Go test suite from the application root.
set -eu

cd "$(dirname "$0")/.."
go test ./...
