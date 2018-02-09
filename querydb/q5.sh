#!/bin/sh

echo running query q5.sql

## constants ##
DN=sqlite
DB=/home/cecil/data/hier/hier.db

export DN DB

go run querydb.go \
    -input input1.csv \
    -parameters 1,2 \
    -query q5.sql \
    -output q5.txt \
    -driverName $DN \
    -urlref DB
