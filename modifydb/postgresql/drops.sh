#!/bin/sh

echo Start drops at `date`

DBURL="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
export DBURL

echo drops statements...
modifydb -query node_drop.sql -urlref DBURL -driver postgres
modifydb -query edge_drop.sql -urlref DBURL -driver postgres

echo End drops at `date`
