#!/usr/bin/env bash
#
# go run main.go
go build -tags=jsoniter -o hello_gin .
./hello_gin
