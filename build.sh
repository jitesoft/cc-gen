#! /bin/bash

mkdir -p bin

echo "Building linux arm64"
GOARCH=arm64 GOOS=linux go build
tar -zc cc-gen -f bin/cc-gen-arm64-linux.tar.gz
rm cc-gen

echo "Building linux amd64"
GOARCH=amd64 GOOS=linux go build
tar -zc cc-gen -f bin/cc-gen-amd64-linux.tar.gz
rm cc-gen

echo "Building linux arm"
GOARCH=arm GOOS=linux go build
tar -zc cc-gen -f bin/cc-gen-arm-linux.tar.gz
rm cc-gen

echo "Building darwin amd64"
GOARCH=amd64 GOOS=darwin go build
tar -zc cc-gen -f bin/cc-gen-amd64-darwin.tar.gz
rm cc-gen

echo "Building darwin arm64"
GOARCH=arm64 GOOS=darwin go build
tar -zc cc-gen -f bin/cc-gen-arm64-darwin.tar.gz
rm cc-gen

echo "Building windows amd64"
GOARCH=amd64 GOOS=windows go build
tar -zc cc-gen.exe -f bin/cc-gen-amd64-windows.tar.gz
rm cc-gen.exe

echo "Building linux riscv64"
GOARCH=riscv64 GOOS=linux go build
tar -zc cc-gen -f bin/cc-gen-riscv64-linux.tar.gz
rm cc-gen

echo "Building linux mips64"
GOARCH=mips64 GOOS=linux go build
tar -zc cc-gen -f bin/cc-gen-mips64-linux.tar.gz
rm cc-gen
