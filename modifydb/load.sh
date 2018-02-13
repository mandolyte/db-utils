#!/bin/sh

echo Start create and import at `date`

# remove the db and start from scratch
rm -f here.db
DBURL=here.db
export DBURL

echo creates statements...
go run modifydb.go -query node_create.sql -urlref DBURL -driver sqlite
go run modifydb.go -query edge_create.sql -urlref DBURL -driver sqlite

echo inserting into node table at `date`
go run modifydb.go -query node_insert.sql \
    -input $HOME/data/hier/hier_nodes.csv \
    -parameters 1 \
    -urlref DBURL -driver sqlite

echo inserting into edge table at `date`
go run modifydb.go -query edge_insert.sql \
    -input $HOME/data/hier/hier_edges.csv \
    -parameters 1,2,3 \
    -urlref DBURL -driver sqlite

echo End create and import at `date`
