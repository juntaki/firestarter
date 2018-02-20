#!/bin/bash

protoc \
    --proto_path=$GOPATH/src:. \
    --twirp_out=. \
    --go_out=. \
    proto/*.proto
