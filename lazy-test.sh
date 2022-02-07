#!/usr/bin/env bash

full_path=$(realpath $0)
# echo $full_path
dir_path=$(dirname $full_path)
# echo $dir_path
cd "${dir_path}"

go env
go mod tidy
go test -v --cover ./...

mkdir -p "${dir_path}/bin"

GOOS=js GOARCH=wasm go build -o "${dir_path}/bin/test.wasm"
