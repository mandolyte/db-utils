#!/bin/sh

echo Start creates at `date`

DB="postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable"
DN="postgres"
export DB DN

echo creates statements...
modifydb -query createdb.sql -urlref DB -driver $DN

modifydb -query node_create.sql -urlref DB -driver $DN
modifydb -query edge_create.sql -urlref DB -driver $DN

echo End creates at `date`
