#!/bin/sh

echo Start load at `date`

DB=here.db
DN=sqlite
export DB

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

echo End load at `date`
