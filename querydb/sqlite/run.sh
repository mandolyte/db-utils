#!/bin/sh

if [ "$1x" = "x" ]; then
    echo query arg is mising - do not include extension
    echo for example: q1
    exit 
fi
echo running query $1

## constants ##
DN=sqlite
DB=../../modifydb/sqlite/here.db

export DN DB

querydb \
    -query $1.sql \
    -output $1.txt \
    -driver $DN \
    -urlref DB
