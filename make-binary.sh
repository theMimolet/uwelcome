#!/bin/bash
# Script to build uwelcome binary for Linux on x86_64 and arm64

go build
VERSION=$(./uwelcome version)
rm ./uwelcome

# x86_64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o uwelcome_"${VERSION}"_linux_amd64

# arm64
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o uwelcome_"${VERSION}"_linux_arm64

echo "Built uwelcome v$VERSION for Linux on x86_64 and arm64"
