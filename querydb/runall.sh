#!/bin/sh
LIST="q1 q2 q3 q4"

## constants ##
DN=sqlite
DB=/home/cecil/data/hier/hier.db
export DN DB

for i in $LIST
do
	go run querydb.go \
	    -query $1.sql \
	    -output $1.txt \
	    -driverName $DN \
	    -urlref DB
done
