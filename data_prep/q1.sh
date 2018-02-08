#!/bin/sh
DB=$HOME/data/hier/hier.db
sqlite3 $DB <<-EoF

.mode csv
.headers on
.output q1.txt

select count(*) from node
union all
select count(*) from edge
;

.quit
EoF