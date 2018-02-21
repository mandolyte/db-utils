# Notes 

## Setup 
Before running the queries below:
- the data must be loaded into the target database (see data_prep folder)
- the correct driver must be included in the import

- Note that my own testing uses the cznic/sqlite package. And the "url"
for sqlite is simply a filename.

- These tests were done using the Ubuntu subsystem in Windows 10.

## Queries
Run q1:
```
$ sh run.sh q1
running query q1
2018/02/19 10:03:52 SQL is:
select count(*) from node
union all
select count(*) from edge
2018/02/19 10:03:52 Total Rows: 3
2018/02/19 10:03:52 Elapsed Time: 105.8287ms
$
```

Run q2:
```
$ sh run.sh q2
running query q2
2018/02/19 10:04:21 SQL is:
select n.id, count(e.id)
from node n
inner join edge e
on n.id = e.from_id
group by n.id
order by count(e.from_id) desc
limit 5
2018/02/19 10:04:29 Total Rows: 6
2018/02/19 10:04:29 Elapsed Time: 8.2069265s
$
```

Run q3:
```
$ sh run.sh q3
running query q3
2018/02/19 10:04:55 SQL is:
with recursive expand(level, from_id, to_id) AS (
    select 1 as level, from_id, to_id
    from edge as e
    where from_id = '/usr/lib/grub'

    union

    select x.level+1, e.from_id, e.to_id
    from expand x
    inner join edge e
    on e.from_id = x.to_id
)
select level, from_id as parent, to_id as child from expand
2018/02/19 10:04:55 Total Rows: 295
2018/02/19 10:04:55 Elapsed Time: 73.5758ms
$
```

Run q4:
```
$ sh run.sh q4
running query q4
2018/02/19 10:05:24 SQL is:
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
2018/02/19 10:05:25 Total Rows: 2
2018/02/19 10:05:25 Elapsed Time: 1.0964558s
$
```

The next query demonstrates the use of an input CSV to drive 
substitution parameters in the SQL. The SQL query is executed
once per row with values from the input.
This execution set two extra arguments: `-input` and `-parameters`.
See the script `q5.sh` for details.
```
$ sh q5.sh
running query q5.sql
2018/02/19 10:06:44 SQL is:
select id
from node
where id like (? || '%')
and id like ('%' || ?)
2018/02/19 10:06:44 Total Rows: 278
2018/02/19 10:06:44 Elapsed Time: 546.1548ms
$
```

Note that the postgres query returns 261 rows whereas the sqlite query
returns 278. The sqlite query is case insensitive and returns results
with capital letters. The matches the behavior of sqlite3 CLI:
```
$ sqlite3 test.db
SQLite version 3.11.0 2016-02-15 17:29:24
Enter ".help" for usage hints.
sqlite> create table test (
   ...>  cola text
   ...> );
sqlite> .schema
CREATE TABLE test (
 cola text
);
sqlite> insert into test values ('aWonderfulLife');
sqlite> insert into test values ('toBeOrNotToBe');
sqlite> insert into test values ('AWonderfulLife');
sqlite> insert into test values ('ToBeOrNotToBe');
sqlite> select * from test;
aWonderfulLife
toBeOrNotToBe
AWonderfulLife
ToBeOrNotToBe
sqlite> select * from test where cola like 'a%';
aWonderfulLife
AWonderfulLife
sqlite> .quit
$
```