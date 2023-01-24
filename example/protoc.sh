#!/usr/bin/env bash

shopt -s globstar # on MacOS need to bash installed via brew @see https://gist.github.com/reggi/475793ea1846affbcfe8#updating-bash

GATEWAY_PATH=$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.15.0/

protoc -I ./api/src -I $GATEWAY_PATH \
    --go_out=./api --go_opt=paths=source_relative --go-grpc_out=./api --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out ./api --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=false \
    api/src/**/*.proto

rm -rf api/google