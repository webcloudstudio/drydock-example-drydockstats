#!/bin/sh
# ddkstat test suite: go vet, then go test, over the whole module.
set -eu

cd "$(dirname "$0")/.."
go vet ./...
go test ./...
