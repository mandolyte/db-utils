# Notes 

*Observation*: the `lib/pq` using ordinal markers instead of question 
marks. For example, instead of the typical SQL:
```
insert into public.edge (id,from_id, to_id)
values (?,?,?) 
```

The `lib/pq` package uses `$n` style:
```
insert into public.edge (id,from_id, to_id)
values ($1, $2, $3) 
```


First, drop the tables if needed:
```
$ sh drops.sh
Start drops at Fri, Feb 16, 2018 11:14:30 AM
drops statements...
2018/02/16 11:14:30 SQL is:
drop table if exists public.node
2018/02/16 11:14:32 Total Affected Rows: 0
2018/02/16 11:14:32 Elapsed Time: 1.0860185s
2018/02/16 11:14:32 SQL is:
drop table if exists public.edge
2018/02/16 11:14:33 Total Affected Rows: 0
2018/02/16 11:14:33 Elapsed Time: 1.0615896s
End drops at Fri, Feb 16, 2018 11:14:33 AM
$
```

Second, create the tables. Note the "database" already existed
from other testing. For CDB, it is the equivalent of a schema.
```
$ sh creates.sh
Start creates at Fri, Feb 16, 2018 11:14:58 AM
creates statements...
2018/02/16 11:14:58 SQL is:
create database public
2018/02/16 11:14:59 dbutils.Exec() error: Exec() and Rollback() Errors:
SQL:create database public;Args:[];Message:pq: database "public" already exists
and
<nil>
2018/02/16 11:14:59 SQL is:
CREATE TABLE public.node (
    id text primary key
)
2018/02/16 11:15:00 Total Affected Rows: 0
2018/02/16 11:15:00 Elapsed Time: 1.0220585s
2018/02/16 11:15:00 SQL is:
CREATE TABLE public.edge (
    id text primary key,
    from_id text not null,
    to_id text not null
)
2018/02/16 11:15:01 Total Affected Rows: 0
2018/02/16 11:15:01 Elapsed Time: 1.0271762s
End creates at Fri, Feb 16, 2018 11:15:01 AM
$
```

Load, the tables using input CSV. Note that I gave up after 16m.
```
$ sh load.sh
Start create and import at Fri, Feb 16, 2018 11:18:03 AM
inserting into node table at Fri, Feb 16, 2018 11:18:03 AM
2018/02/16 11:18:03 SQL is:
insert into public.node (id) values ($1)

$ echo `date`
Fri, Feb 16, 2018 11:34:08 AM
$
```

Finally, index the tables:
```
```