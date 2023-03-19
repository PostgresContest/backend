#!/bin/bash

if ! command -v goose &> /dev/null
then
    go get github.com/pressly/goose/v3/cmd/goose@latest
    go install github.com/pressly/goose/v3/cmd/goose@latest
fi


DSN=$(go run main.go dsn private)

goose -dir migrations postgres "${DSN}" "$@" sql