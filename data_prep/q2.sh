#!/bin/sh
DB=$HOME/data/hier/hier.db
sqlite3 $DB <<-EoF

.mode csv
.headers on
.output q2.txt

select n.id, count(e.id)
from node n
inner join edge e
on n.id = e.from_id
group by n.id
order by count(e.from_id) desc
limit 5
;

.quit
EoF