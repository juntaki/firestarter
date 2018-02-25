#!/bin/bash

# go get github.com/twitchtv/twirp/protoc-gen-twirp
# go get github.com/thechriswalker/protoc-gen-twirp_js

protoc \
    -I proto \
    --proto_path=$GOPATH/src:. \
    --twirp_out=proto \
    --go_out=proto \
    --twirp_js_out=./admin/proto \
    --js_out=import_style=commonjs,binary:./admin/proto \
    proto/*.proto
