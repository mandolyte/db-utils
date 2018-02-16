#!/bin/sh

echo Start create and import at `date`

DB="postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable"
DN="postgres"
export DB DN

echo inserting into node table at `date`
modifydb -query node_insert.sql \
    -input $HOME/data/hier/hier_nodes.csv \
    -parameters 1 \
    -urlref DB -driver $DN

echo inserting into edge table at `date`
modifydb -query edge_insert.sql \
    -input $HOME/data/hier/hier_edges.csv \
    -parameters 1,2,3 \
    -urlref DB -driver $DN

echo End create and import at `date`
