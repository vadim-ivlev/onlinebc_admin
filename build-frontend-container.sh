#!/bin/bash


# build a docker image 
docker build -t rgru/onlinebc_admin:latest -f Dockerfile-frontend . 

# push the docker image 
docker login
docker push rgru/onlinebc_admin:latest
