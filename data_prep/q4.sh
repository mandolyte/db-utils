#!/bin/sh
DB=$HOME/data/hier/hier.db
sqlite3 $DB <<-EoF

.mode csv
.headers on
.output q4.txt

with recursive expand(level, from_id, to_id) AS (
    select 1 as level, from_id, to_id 
    from edge as e
    where from_id = '/usr/local/go'
    
    union

    select x.level+1, e.from_id, e.to_id
    from expand x
    inner join edge e
    on e.from_id = x.to_id
), results as (
    select level, from_id as parent, to_id as child from expand
)
select max(level), count(*)
from results
;

.quit
EoF

