#!/bin/bash

if [ -z "$1" ]
then
  echo "No version supplied"
  exit 1
fi

go mod edit -require github.com/aws/aws-sdk-go@$1
go mod tidy
go get ./...
go install github.com/golang/mock/mockgen