#!/usr/bin/env bash

set -e

GOOS=linux GOARCH=amd64 go build -o dkr main.go

GOOS=linux GOARCH=amd64 go build -o dkrd cmd/boot/main.go
