#!/usr/bin/env bash

git pull origin master

sudo docker rm -f docker-shoplist:latest

sudo docker build --build-arg app_env=prod --build-arg app_host_port=3000 --no-cache --tag docker-shoplist:latest -f build/package/Dockerfile .
