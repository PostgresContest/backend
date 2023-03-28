#!/bin/bash

if ! command -v golangci-lint &> /dev/null
then
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
fi

check()
{
  golangci-lint run
}

fix()
{
  gci write --skip-generated .
  goimports -w .
  golangci-lint run --fix
}

"$@"