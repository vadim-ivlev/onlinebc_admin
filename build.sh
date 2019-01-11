#!/bin/bash

rm -rf build
mkdir build

go build -o ./build/onlinebc


cp -rf configs ./build/
cp -rf migrations ./build/
cp -rf templates ./build/
cp docker-compose.yml ./build/
cp .gitignore ./build/
cp README-FRONT.md ./build/README.md

cp -rf build/* ../onlinebc
cp build/.gitignore ../onlinebc