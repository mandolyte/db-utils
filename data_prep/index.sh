#!/bin/sh

echo Start indexing at `date`

# remove the db and start from scratch
rm hier.db

sqlite3 hier.db <<-EoF
create unique index nodeidx on node(id)
;

create unique index edgeidx on edge(id)
;

create index fromidx on edge(from_id)
;

create index toidx on edge(to_id)
;

.quit
EoF

echo End indexing at `date`