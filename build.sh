#!/bin/bash

rm -rf build
mkdir build

env GOOS=linux GOARCH=amd64 go build -v -o ./build/onlinebc main.go
env GOOS=windows GOARCH=386 go build -v -o ./build/onlinebc.exe main.go


cp -rf configs ./build/
cp -rf migrations ./build/
cp -rf templates ./build/
cp docker-compose.yml ./build/
cp .gitignore ./build/
cp readme-front.md ./build/README.md

cp -rf build/* ../onlinebc
cp build/.gitignore ../onlinebc