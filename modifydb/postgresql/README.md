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
Start drops at Fri, Feb 16, 2018 8:18:50 AM
drops statements...
2018/02/16 08:18:50 SQL is:
drop table if exists public.node
2018/02/16 08:18:50 Total Affected Rows: 0
2018/02/16 08:18:50 Elapsed Time: 73.0483ms
2018/02/16 08:18:50 SQL is:
drop table if exists public.edge
2018/02/16 08:18:50 Total Affected Rows: 0
2018/02/16 08:18:50 Elapsed Time: 64.0408ms
End drops at Fri, Feb 16, 2018 8:18:50 AM
$
```

Second, create the tables:
```
$ sh creates.sh
Start creates at Fri, Feb 16, 2018 8:18:53 AM
creates statements...
2018/02/16 08:18:53 SQL is:
CREATE TABLE  if not exists public.node (
    id text primary key
)
2018/02/16 08:18:53 Total Affected Rows: 0
2018/02/16 08:18:53 Elapsed Time: 84.0545ms
2018/02/16 08:18:54 SQL is:
CREATE TABLE if not exists  public.edge (
    id text primary key,
    from_id text not null,
    to_id text not null
)
2018/02/16 08:18:54 Total Affected Rows: 0
2018/02/16 08:18:54 Elapsed Time: 73.0495ms
End creates at Fri, Feb 16, 2018 8:18:54 AM
$ 
```

Load, the tables using input CSV:
```
$ sh load.sh
Start create and import at Fri, Feb 16, 2018 8:19:16 AM
inserting into node table at Fri, Feb 16, 2018 8:19:16 AM
2018/02/16 08:19:16 SQL is:
insert into node (id) values ($1)
2018/02/16 08:19:30 Total Affected Rows: 96898
2018/02/16 08:19:30 Elapsed Time: 14.1319206s
inserting into edge table at Fri, Feb 16, 2018 8:19:30 AM
2018/02/16 08:19:30 SQL is:
insert into public.edge (id,from_id, to_id)
values ($1, $2, $3)
2018/02/16 08:20:36 Total Affected Rows: 399959
2018/02/16 08:20:36 Elapsed Time: 1m6.1740116s
End create and import at Fri, Feb 16, 2018 8:20:36 AM
$
```

Finally, index the tables:
```
$ sh index.sh
Start indexing at Fri, Feb 16, 2018 8:27:34 AM
2018/02/16 08:27:34 Error ioutil.ReadFile() on node_indes.sql:
open node_indes.sql: The system cannot find the file specified.
2018/02/16 08:27:34 SQL is:
create unique index edgeidx on edge(id)
2018/02/16 08:27:38 Total Affected Rows: 0
2018/02/16 08:27:38 Elapsed Time: 3.598039s
2018/02/16 08:27:38 SQL is:
create index fromidx on edge(from_id)
2018/02/16 08:27:41 Total Affected Rows: 0
2018/02/16 08:27:41 Elapsed Time: 2.6677706s
2018/02/16 08:27:41 SQL is:
create index toidx on edge(to_id)
2018/02/16 08:27:43 Total Affected Rows: 0
2018/02/16 08:27:43 Elapsed Time: 2.4883815s
End indexing at Fri, Feb 16, 2018 8:27:43 AM
$
```