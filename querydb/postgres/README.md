# Notes 

## Setup 
Before running the queries below:
- the data must be loaded into the target database (see data_prep folder)
- the correct driver must be included in the import

## Queries
Run q1:
```
$ sh run.sh q1
running query q1
2018/02/19 09:45:01 SQL is:
select count(*) from node
union all
select count(*) from edge
2018/02/19 09:45:01 Total Rows: 3
2018/02/19 09:45:01 Elapsed Time: 282.6353ms
$
```

Run q2:
```
$ sh run.sh q2
running query q2
2018/02/19 09:46:02 SQL is:
select n.id, count(e.id)
from node n
inner join edge e
on n.id = e.from_id
group by n.id
order by count(e.from_id) desc
limit 5
2018/02/19 09:46:03 Total Rows: 6
2018/02/19 09:46:03 Elapsed Time: 414.2078ms
$
```

Run q3:
```
$ sh run.sh q3
running query q3
2018/02/19 09:46:34 SQL is:
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
2018/02/19 09:46:34 Total Rows: 295
2018/02/19 09:46:34 Elapsed Time: 73.1467ms
$
```

Run q4:
```
$ sh run.sh q4
running query q4
2018/02/19 09:47:00 SQL is:
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
2018/02/19 09:47:00 Total Rows: 2
2018/02/19 09:47:00 Elapsed Time: 160.2465ms
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
2018/02/19 09:52:03 SQL is:
select id
from node
where id like ($1 || '%')
and id like ('%' || $2)
2018/02/19 09:52:04 Total Rows: 261
2018/02/19 09:52:04 Elapsed Time: 106.9699ms
$
```

Note that the postgres query returns 261 rows whereas the sqlite query
returns 278. The sqlite query is case insensitive and returns results
with capital letters. The matches the behavior of sqlite3 CLI.
See the sqlite folder for details.
