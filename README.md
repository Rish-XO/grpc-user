# gRPC User Service in Go

This repository contains a gRPC service written in Go for managing user details with search functionality. The service includes endpoints for fetching user details based on a user ID, retrieving a list of user details based on a list of user IDs, and searching user details based on specific criteria.

## Features

- Fetch user details by user ID
- Retrieve a list of user details by a list of user IDs
- Search user details based on criteria such as city, phone number, and marital status
- Unit tests to verify the correctness of the service

## Prerequisites

- Go 1.18 or later
- Protocol Buffers compiler (`protoc`)
- Protobuf plugins for Go (`protoc-gen-go`, `protoc-gen-go-grpc`)

## Installation

### Install `protoc` on Windows

1. Download `protoc-<version>-win64.zip` from the [Protocol Buffers Releases](https://github.com/protocolbuffers/protobuf/releases) page.
2. Extract the contents to a directory, e.g., `C:\protoc`.
3. Add `C:\protoc\bin` to your system's PATH.

### Install Go Plugins

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


Generate Go Code from Protobuf
protoc --go_out=. --go-grpc_out=. user.proto

Run the Service Locally
go run server.go

Run the unit tests using the following command:
go test ./...
