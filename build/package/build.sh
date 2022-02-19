#!/usr/bin/env bash

docker rm -f docker-shoplist:latest
docker build  --no-cache --tag docker-shoplist:latest -f build/package/Dockerfile .

#docker run docker-shoplist:latest
