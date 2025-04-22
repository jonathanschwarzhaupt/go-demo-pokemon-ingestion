#!/bin/bash

if [ -f .env ]; then
    source .env
fi

# build the go binary
go build -o ./bin/pokeapi ./cmd/api

# Run the binary with runtime flags
./bin/pokeapi -dbUrl="$DB_URL" -env="development"