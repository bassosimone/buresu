#!/bin/bash
set -euo pipefail
go test -race -count 1 -coverprofile=coverage.out "$@"
go tool cover -html=coverage.out
