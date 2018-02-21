#!/bin/sh
LIST="q1 q2 q3 q4"

## constants ##
DN=postgres
DB="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
export DN DB

for i in $LIST
do
	echo 
	echo ++++++++++++++++++++
	echo Running query $i.sql
	querydb \
	    -query $i.sql \
	    -output $i.txt \
	    -driver $DN \
	    -urlref DB
done
