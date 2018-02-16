#!/bin/sh

echo Start indexing at `date`

DB="postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable"
DN="postgres"
export DB

modifydb -query node_indes.sql -urlref DB -driver $DN
modifydb -query edge_index.sql -urlref DB -driver $DN
modifydb -query edge_from_index.sql -urlref DB -driver $DN
modifydb -query edge_to_index.sql -urlref DB -driver $DN

echo End indexing at `date`

