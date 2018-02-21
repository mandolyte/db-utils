#!/bin/sh

echo running query q5.sql

## constants ##
DN=postgres
DB="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"

export DN DB

querydb \
    -input input1.csv \
    -parameters 1,2 \
    -query q5.sql \
    -output q5.txt \
    -driver $DN \
    -urlref DB
