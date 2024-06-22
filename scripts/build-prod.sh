#!/bin/bash

# Build binary for linux platform
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/pokeapi ./cmd/api

# Build the Docker image
docker build -t jschwarzhaupt/pokeapi -f ./remote/Dockerfile .

# Push the Docker image
docker push jschwarzhaupt/pokeapi