#!/bin/bash

rm -rf build

go build -o ./build/onlinebc

cp -rf configs ./build/
cp -rf migrations ./build/
cp -rf templates ./build/
cp docker-compose.yml ./build/
