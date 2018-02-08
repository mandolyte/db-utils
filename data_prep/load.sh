#!/bin/sh

echo Start create and import at `date`

# remove the db and start from scratch
rm hier.db

sqlite3 hier.db <<-EoF
CREATE TABLE node (
    id text key
)
;
CREATE TABLE edge (
    id text key,
    from_id text not null,
    to_id text not null
)
;
.mode csv
.import hier_nodes.csv node
.import hier_edges.csv edge

.quit
EoF

echo End create and import at `date`