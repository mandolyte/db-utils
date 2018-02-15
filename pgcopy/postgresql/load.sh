#!/bin/sh

echo Start create and import at `date`

DBURL="postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
#DBURL="postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable"
export DBURL

echo inserting into node table at `date`
go run pgcopy.go -schema public \
    -table node \
    -input $HOME/data/hier/hier_nodes.csv \
    -urlvar DBURL
    
echo inserting into edge table at `date`
go run pgcopy.go -schema public \
    -table edge \
    -input $HOME/data/hier/hier_edges.csv \
    -urlvar DBURL
    
echo End create and import at `date`
