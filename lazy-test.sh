#!/usr/bin/env bash

full_path=$(realpath $0)
# echo $full_path
dir_path=$(dirname $full_path)
# echo $dir_path
cd "${dir_path}"

mkdir -p "${dir_path}/bin"

echo "***** go env *****"
go env
echo "***** go mod tidy *****"
go mod tidy
echo "***** go test -v --cover *****"
go test -v --cover ./...
# go test --cover ./...
# go test ./...

echo "***** GOOS=js GOARCH=wasm go build -o ${dir_path}/bin/test.wasm *****"
GOOS=js GOARCH=wasm go build -o "${dir_path}/bin/test.wasm"
