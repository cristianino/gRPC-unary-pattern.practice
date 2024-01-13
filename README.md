# gRPC Unary Pattern Practice

This repository is a practical guide for software engineers aspiring to master the Unary pattern in gRPC using the Golang programming language.

## Overview

The Unary pattern in gRPC represents a straightforward one-to-one communication model, where a client sends a single request to the server and receives a single response. This repository provides a hands-on implementation of the Unary pattern, allowing developers to grasp the fundamentals effectively.

## Project Structure

The repository consists of the following files:

- `server.go`: Golang implementation of the gRPC server utilizing the Unary pattern.
- `client.go`: Golang implementation of the gRPC client, showcasing how to make a Unary RPC call.
- `hellopb.proto`: Protocol Buffers file defining the service, message types, and RPC methods.

## Run Server

```bash
go run hello_server/serve.go
```

## Run Client

```bash
go run hello_client/client.go
```

## Response
Response Hello: Welcome Sr Cristian!
