#!/bin/sh
# generate golang for protobuf

protoc -I$GOPATH/src/github.com/kubenext/kubefun/vendor -I$GOPATH/src/github.com/kubenext/kubefun -I. --go_out=plugins=grpc:. dashboard.proto
