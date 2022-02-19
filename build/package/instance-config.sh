#!/usr/bin/env bash

# Updates the system
sudo yum update

# Install Git
sudo yum install git

# Install Docker
sudo yum install docker -y

# Start docker service
sudo service docker start

# Add ec2-user to docker user group to allow running commands without sudo
sudo usermod -a -G docker ec2-user

# Creates a directory for the app
mkdir app

# Clone the repo
git clone https://github.com/lgarcia93/shoplist.git

# Navigate to the source code folder
cd shoplist

# Checkout master
git checkout master

# Remove existing images
sudo docker rm -f docker-shoplist:latest

# Build the docker image
sudo docker build  --no-cache --tag docker-shoplist:latest -f build/package/Dockerfile .

# Run container mapping host port 3000 to container's port 5000
docker run -p 3000:5000 docker-shoplist --tail


