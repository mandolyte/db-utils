# Notes 

## Setup 
Before running the queries below:
- the data must be loaded into the target database (see data_prep folder)
- the correct driver must be included in the import

- Note that my own testing uses the cznic/sqlite package. And the "url"
for sqlite is simply a filename.

## Queries
Run q1:
```
$ sh run.sh q1
running query q1
2018/02/08 09:07:29 SQL is:
select count(*) from node
union all
select count(*) from edge
2018/02/08 09:07:29 Total Rows: 3
2018/02/08 09:07:29 Elapsed Time: 37.069216ms
$ 
```

Run q2:
```
$ sh run.sh q2
running query q2
2018/02/08 09:08:00 SQL is:
select n.id, count(e.id)
from node n
inner join edge e
on n.id = e.from_id
group by n.id
order by count(e.from_id) desc
limit 5
2018/02/08 09:08:11 Total Rows: 6
2018/02/08 09:08:11 Elapsed Time: 11.347837484s
$ 
```

Run q3:
```
$ sh run.sh q3
running query q3
2018/02/08 09:08:32 SQL is:
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
2018/02/08 09:08:32 Total Rows: 295
2018/02/08 09:08:32 Elapsed Time: 94.070851ms
$ 
```

Run q4:
```
$ sh run.sh q4
running query q4
2018/02/08 09:08:57 SQL is:
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
2018/02/08 09:08:58 Total Rows: 2
2018/02/08 09:08:58 Elapsed Time: 1.328337719s
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
2018/02/09 14:20:03 SQL is:
select id
from node
where id like (? || '%')
and id like ('%' || ?)
2018/02/09 14:20:04 Total Rows: 277
2018/02/09 14:20:04 Elapsed Time: 762.919339ms
$
```