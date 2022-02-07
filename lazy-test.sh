#!/usr/bin/env bash

go env
go mod tidy
go test -v --cover ./...

full_path=$(realpath $0)
# echo $full_path
dir_path=$(dirname $full_path)
# echo $dir_path
mkdir -p "${dir_path}/bin"

GOOS=js GOARCH=wasm go build -o "${dir_path}/bin/test.wasm"
