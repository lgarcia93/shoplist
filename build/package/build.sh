#!/usr/bin/env bash

git pull origin master

sudo docker rm -f docker-shoplist:latest

sudo docker build  --no-cache --tag docker-shoplist:latest -f build/package/Dockerfile .
