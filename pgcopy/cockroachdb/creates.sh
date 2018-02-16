#!/bin/sh

echo Start creates at `date`

DBURL="postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable"
export DBURL

echo creates statements...
modifydb -query createdb.sql -urlref DBURL -driver postgres

modifydb -query node_create.sql -urlref DBURL -driver postgres
modifydb -query edge_create.sql -urlref DBURL -driver postgres

echo End creates at `date`
