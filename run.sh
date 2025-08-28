#!/bin/bash
GITHUB_USERNAME=gfedacs
GITHUB_EMAIL=gfedacs@hotmail.com

# SERVICE_NAME=order
SERVICE_NAME=shipping
RELEASE_VERSION=v1.2.3

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
export PATH="$PATH:$(go env GOPATH)/bin"
# source ~/.zshrc

echo "Generating Go source code"
mkdir -p golang
protoc --go_out=./golang \
  --go_opt=paths=source_relative \
  --go-grpc_out=./golang \
  --go-grpc_opt=paths=source_relative \
 ./${SERVICE_NAME}/*.proto

echo "Generated Go source code files"
ls -al ./golang/${SERVICE_NAME}

cd golang/${SERVICE_NAME}
go mod init \
  github.com/${GITHUB_USERNAME}/microservices-proto/golang/${SERVICE_NAME} || true
go mod tidy || true