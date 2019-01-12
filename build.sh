#!/bin/bash

rm -rf build
mkdir build

env GOOS=linux GOARCH=arm64 go build -o ./build/onlinebc main.go


cp -rf configs ./build/
cp -rf migrations ./build/
cp -rf templates ./build/
cp docker-compose.yml ./build/
cp .gitignore ./build/
cp README-FRONT.md ./build/README.md

cp -rf build/* ../onlinebc
cp build/.gitignore ../onlinebc