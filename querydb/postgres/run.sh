#!/bin/sh

if [ "$1x" = "x" ]; then
    echo query arg is mising - do not include extension
    echo for example: q1
    exit 
fi
echo running query $1

## constants ##
DN=postgres
DB="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"

export DN DB

querydb \
    -query $1.sql \
    -output $1.txt \
    -driver $DN \
    -urlref DB
