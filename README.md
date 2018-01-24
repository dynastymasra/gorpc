# Gorpc

[![License: GPL-3.0](https://img.shields.io/badge/license-GPL--3.0-yellow.svg)](https://opensource.org/licenses/GPL-3.0)

Exercise build client server with Golang using GRPC.

## How To Run

Firstly, Install [Golang](https://golang.org/) and [gRPC](https://grpc.io/), follow the instructions in there.
and then run `./gen.sh` for generate `gRPC` files.

In this folder have three folder
- `contract` - Folder for place protobuf files and protobuf generated files.
- `server` - Golang folder for handle server `gRPC`
- `client` - Client folder for request to server.

Go to server folder and run Golang code in server folder with command `go run main.go`.
If server code is running and then go to client folder and run client code with command like run server code `go run main.go`