#!/bin/sh
LIST="q1 q2 q3 q4"

## constants ##
DN=sqlite
DB=/home/cecil/data/hier/hier.db
export DN DB

for i in $LIST
do
	echo 
	echo ++++++++++++++++++++
	echo Running query $i.sql
	go run querydb.go \
	    -query $i.sql \
	    -output $i.txt \
	    -driverName $DN \
	    -urlref DB
done
