#!/bin/bash -e

protoc -I contract/ --go_out=plugins=grpc:contract contract/hello.proto
