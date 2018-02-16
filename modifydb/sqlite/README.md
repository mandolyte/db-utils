# Notes 

## Windows 10 Ubuntu subsystem
Works!

Run the create step:
```
$ sh creates.sh
Start create at Fri Feb 16 16:05:10 STD 2018
creates statements...
2018/02/16 16:05:10 SQL is:
CREATE TABLE node (
    id text key
)
2018/02/16 16:05:10 Total Affected Rows: 0
2018/02/16 16:05:10 Elapsed Time: 39.0029ms
2018/02/16 16:05:10 SQL is:
CREATE TABLE edge (
    id text key,
    from_id text not null,
    to_id text not null
)
2018/02/16 16:05:10 Total Affected Rows: 0
2018/02/16 16:05:10 Elapsed Time: 40.6141ms
End create at Fri Feb 16 16:05:10 STD 2018
```

Run the load step:
```
$ sh load.sh
Start load at Fri Feb 16 16:07:23 STD 2018
inserting into node table at Fri Feb 16 16:07:23 STD 2018
2018/02/16 16:07:23 SQL is:
insert into node (id)
values (?)
2018/02/16 16:07:36 Total Affected Rows: 96898
2018/02/16 16:07:36 Elapsed Time: 13.0487115s
inserting into edge table at Fri Feb 16 16:07:36 STD 2018
2018/02/16 16:07:36 SQL is:
insert into edge (id,from_id, to_id)
values (?, ?, ?)
2018/02/16 16:08:46 Total Affected Rows: 399959
2018/02/16 16:08:46 Elapsed Time: 1m9.8318125s
End load at Fri Feb 16 16:08:46 STD 2018
$
```

Finally, the index step:
```
$  sh index.sh
Start indexing at Fri Feb 16 16:15:23 STD 2018
2018/02/16 16:15:23 SQL is:
create unique index nodeidx on node(id)
2018/02/16 16:15:24 Total Affected Rows: 0
2018/02/16 16:15:24 Elapsed Time: 1.3334092s
2018/02/16 16:15:24 SQL is:
create unique index edgeidx on edge(id)
2018/02/16 16:15:32 Total Affected Rows: 0
2018/02/16 16:15:32 Elapsed Time: 8.2696662s
2018/02/16 16:15:32 SQL is:
create index fromidx on edge(from_id)
2018/02/16 16:15:39 Total Affected Rows: 0
2018/02/16 16:15:39 Elapsed Time: 6.4002363s
2018/02/16 16:15:39 SQL is:
create index toidx on edge(to_id)
2018/02/16 16:15:45 Total Affected Rows: 0
2018/02/16 16:15:45 Elapsed Time: 6.0539035s
End indexing at Fri Feb 16 16:15:45 STD 2018
$
```

## Windows 10 experience
Note that the cznic/sqlite package
isn't working on Windows yet and fails as shown.
```
$ sh creates.sh
Start create at Fri, Feb 16, 2018 11:40:12 AM
creates statements...
2018/02/16 11:40:12 SQL is:
CREATE TABLE node (
    id text key
)
7(0x7): failed to allocate 56 bytes of memory
7(0x7): OOM at line 140093 of [424a0d3803]
2018/02/16 11:40:12 db.Begin() error: malloc(8) failed
2018/02/16 11:40:12 SQL is:
CREATE TABLE edge (
    id text key,
    from_id text not null,
    to_id text not null
)
7(0x7): failed to allocate 56 bytes of memory
7(0x7): OOM at line 140093 of [424a0d3803]
2018/02/16 11:40:12 db.Begin() error: malloc(8) failed
End create at Fri, Feb 16, 2018 11:40:12 AM
$
```

## Linux (Ubuntu) experience
First, create the tables into new db file:
```
$ sh creates.sh
Start create at Fri Feb 16 11:44:46 EST 2018
creates statements...
2018/02/16 11:44:47 SQL is:
CREATE TABLE node (
    id text key
)
2018/02/16 11:44:47 Total Affected Rows: 0
2018/02/16 11:44:47 Elapsed Time: 502.383851ms
2018/02/16 11:44:47 SQL is:
CREATE TABLE edge (
    id text key,
    from_id text not null,
    to_id text not null
)
2018/02/16 11:44:47 Total Affected Rows: 0
2018/02/16 11:44:47 Elapsed Time: 157.329233ms
End create at Fri Feb 16 11:44:47 EST 2018
$
```

Next load the data:
```
$ sh load.sh
Start load at Fri Feb 16 11:46:50 EST 2018
inserting into node table at Fri Feb 16 11:46:50 EST 2018
2018/02/16 11:46:50 SQL is:
insert into node (id)
values (?)
2018/02/16 11:47:11 Total Affected Rows: 96898
2018/02/16 11:47:11 Elapsed Time: 21.541168257s
inserting into edge table at Fri Feb 16 11:47:11 EST 2018
2018/02/16 11:47:11 SQL is:
insert into edge (id,from_id, to_id)
values (?, ?, ?)
2018/02/16 11:48:56 Total Affected Rows: 399959
2018/02/16 11:48:56 Elapsed Time: 1m44.638819107s
End load at Fri Feb 16 11:48:56 EST 2018
$
```

Finally, index it:
```
$ sh index.sh
Start indexing at Fri Feb 16 11:51:55 EST 2018
2018/02/16 11:51:55 SQL is:
create unique index nodeidx on node(id)
2018/02/16 11:51:57 Total Affected Rows: 0
2018/02/16 11:51:57 Elapsed Time: 1.922811148s
2018/02/16 11:51:57 SQL is:
create unique index edgeidx on edge(id)
2018/02/16 11:52:06 Total Affected Rows: 0
2018/02/16 11:52:06 Elapsed Time: 8.977729896s
2018/02/16 11:52:06 SQL is:
create index fromidx on edge(from_id)
2018/02/16 11:52:13 Total Affected Rows: 0
2018/02/16 11:52:13 Elapsed Time: 7.185557827s
2018/02/16 11:52:13 SQL is:
create index toidx on edge(to_id)
2018/02/16 11:52:19 Total Affected Rows: 0
2018/02/16 11:52:19 Elapsed Time: 6.488586253s
End indexing at Fri Feb 16 11:52:19 EST 2018
$
```