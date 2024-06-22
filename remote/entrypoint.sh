#!/bin/bash

# Default values of arguments
DB_CONN=${DB_URL:-""}
ENVIRONMENT=${ENV:-"staging"}

# Execute the binary with dynamic parameters
exec /usr/bin/pokeapi -dbUrl="$DB_CONN" -env="$ENVIRONMENT"
