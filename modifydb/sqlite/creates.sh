#!/bin/sh

echo Start create at `date`

# remove the db and start from scratch
rm -f here.db
DB=here.db
DN=sqlite
export DB

echo creates statements...
modifydb -query node_create.sql -urlref DB -driver $DN
modifydb -query edge_create.sql -urlref DB -driver $DN

echo End create at `date`
