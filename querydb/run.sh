#!/bin/sh

if [ "$1x" = "x" ]; then
    echo query arg is mising - do not include extension
    echo for example: q1
    exit 
fi
echo running query $1

## constants ##
DN=sqlite
DB=/home/cecil/data/hier/hier.db

export DN DB

go run querydb.go \
    -query $1.sql \
    -output $1.txt \
    -driverName $DN \
    -urlref DB
