#!/bin/sh

echo Start drops at `date`

DB="postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable"
DN="postgres"
export DB DN

echo drops statements...
modifydb -query node_drop.sql -urlref DB -driver $DN
modifydb -query edge_drop.sql -urlref DB -driver $DN

echo End drops at `date`
