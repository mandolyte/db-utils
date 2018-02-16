#!/bin/sh

echo Start indexing at `date`

DB="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
DN="postgres"
export DB DN

modifydb -query node_index.sql -urlref DB -driver $DN
modifydb -query edge_index.sql -urlref DB -driver $DN
modifydb -query edge_from_index.sql -urlref DB -driver $DN
modifydb -query edge_to_index.sql -urlref DB -driver $DN

echo End indexing at `date`

