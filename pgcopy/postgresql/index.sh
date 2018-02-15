#!/bin/sh

echo Start indexing at `date`

DBURL="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
#DBURL="postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable"
export DBURL

modifydb -query node_indes.sql -urlref DBURL -driver postgres
modifydb -query edge_index.sql -urlref DBURL -driver postgres
modifydb -query edge_from_index.sql -urlref DBURL -driver postgres
modifydb -query edge_to_index.sql -urlref DBURL -driver postgres

echo End indexing at `date`

