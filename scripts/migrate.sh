#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd migrations/schema || exit 1

case $1 in
  up)
    echo "running 'up'"
    goose turso "$DB_URL" up

    cd ../..
    sqlc generate
    ;;
  down)
    echo "running 'down'"
    goose turso "$DB_URL" down
    ;;
  *)
    echo "goose command not specified or invalid use (use 'up' or 'down')"
    ;;
esac