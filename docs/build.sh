#!/bin/bash

# regenerate files in build/ directory

rm -rf build
mkdir build

env GOOS=linux GOARCH=amd64 go build -v -o ./build/onlinebc main.go

cp -rf configs ./build/
cp -rf migrations ./build/
cp -rf templates ./build/
cp docker-compose.yml ./build/
cp docker-compose-frontend.yml ./build/
cp .gitignore ./build/
cp readme-production.md ./build/
cp readme-front.md ./build/

cp build/docker-compose-frontend.yml ../onlinebc/docker-compose.yml
cp build/readme-front.md ../onlinebc
cp build/.gitignore ../onlinebc


# build a docker image 
docker build -t vadimivlev/onlinebc:latest -f Dockerfile-frontend ./build 

# push the docker image 
docker login
docker push vadimivlev/onlinebc:latest
