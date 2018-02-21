#!/bin/sh

echo running query q5.sql

## constants ##
DN=sqlite
DB=../../modifydb/sqlite/here.db

export DN DB

querydb \
    -input input1.csv \
    -parameters 1,2 \
    -query q5.sql \
    -output q5.txt \
    -driver $DN \
    -urlref DB
