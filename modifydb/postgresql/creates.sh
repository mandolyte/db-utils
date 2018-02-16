#!/bin/sh

echo Start creates at `date`

DBURL="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
export DBURL

echo creates statements...
modifydb -query node_create.sql -urlref DBURL -driver postgres
modifydb -query edge_create.sql -urlref DBURL -driver postgres

echo End creates at `date`
