#!/usr/bin/env bash

FMT=$(go fmt ./...)
if [[ -n "$FMT" ]]; then
  exit 1
fi
go vet ./...
golangci-lint run --enable-all ./...