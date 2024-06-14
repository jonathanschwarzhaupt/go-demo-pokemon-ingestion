#!/bin/bash

if [ -f .env ]; then
    source .env
fi


cd sql/schema || exit
goose turso "$DATABASE_CONN" up
