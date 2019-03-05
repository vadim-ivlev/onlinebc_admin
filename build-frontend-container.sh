#!/bin/bash


# build a docker image 
docker build -t rgru/onlinebc_admin:latest -f Dockerfile-frontend . 

# push the docker image 
# docker login
# docker push rgru/onlinebc_admin:latest


# копируем docker-compose-frontend.yml и 
cp docker-compose-frontend.yml ../onlinebc/docker-compose.yml
cp readme-frontend.md ../onlinebc/readme.md
