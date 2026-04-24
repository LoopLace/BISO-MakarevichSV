#!/usr/bin/env bash
set -euo pipefail

mkdir -p artifacts

echo "[1/4] gofmt"
gofmt -w ./cmd ./internal

echo "[2/4] tests + race + coverage"
go test -v -race -coverprofile=artifacts/coverage.out ./...
go tool cover -func=artifacts/coverage.out | tee artifacts/coverage.txt

echo "[3/4] build"
go build -o artifacts/app ./cmd/app

echo "[4/4] optional linters"
if command -v golangci-lint >/dev/null 2>&1; then
  golangci-lint run ./...
else
  echo "golangci-lint not found: skipped"
fi

if command -v gosec >/dev/null 2>&1; then
  gosec ./...
else
  echo "gosec not found: skipped"
fi

echo "Local CI completed successfully."
