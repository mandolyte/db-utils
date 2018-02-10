#!/bin/sh
echo start at `date`
go run hier_data_prep.go \
    -i $HOME/data/hier/hier_base.csv \
    -node $HOME/data/hier/hier_nodes.csv \
    -edge $HOME/data/hier/hier_edges.csv
echo end at `date`
