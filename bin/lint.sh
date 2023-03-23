#!/bin/bash

if ! command -v golangci-lint &> /dev/null
then
    go get  github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
    go install  github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
fi

check()
{
  golangci-lint run
}

fix()
{
  golangci-lint run --fix
}

"$@"