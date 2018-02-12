#!/bin/sh

echo Start indexing at `date`

# remove the db and start from scratch
rm -f here.db
DBRUL=here.db
export DBURL

go run modifydb.go -query node_index.sql -urlref DBURL -driverName sqlite
go run modifydb.go -query edge_index.sql -urlref DBURL -driverName sqlite
go run modifydb.go -query edge_from_index.sql -urlref DBURL -driverName sqlite
go run modifydb.go -query edge_to_index.sql -urlref DBURL -driverName sqlite

echo End indexing at `date`
