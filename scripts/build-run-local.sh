#!/bin/bash

if [ -f .env ]; then
    source .env
fi

# build the go binary
go build -o pokeapi ./cmd/api

# Run the binary with flags
./pokeapi -dbConn="$DB_URL" -env="development"