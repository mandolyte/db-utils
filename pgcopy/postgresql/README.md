# PGCOPY
Copy/import a CSV into a Postgresql database. Use `-help` to show:
```
$ pgcopy -help
Start Time: 2017-05-15 12:07:42.9777408 -0400 EDT
Help:
  -help
        Show help message
  -input string
        Input CSV to COPY from
  -schema string
        Schema with table to COPY to
  -table string
        Table name to COPY to
  -urlvar string
        Environment variable with DB URL
The input CSV file must have headers that match the table names
$
```

## A complete test using Postgresql 10
The below was done on a Windows 10 laptop with Postresql 10.x installed.

First use modifydb to create the tables:
```
$ sh creates.sh
Start creates at Wed, Feb 14, 2018 9:47:30 AM
creates statements...
2018/02/14 09:47:30 SQL is:
CREATE TABLE public.node (
    id text primary key
)
2018/02/14 09:47:30 Total Affected Rows: 0
2018/02/14 09:47:30 Elapsed Time: 85.1851ms
2018/02/14 09:47:30 SQL is:
CREATE TABLE public.edge (
    id text primary key,
    from_id text not null,
    to_id text not null
)
2018/02/14 09:47:31 Total Affected Rows: 0
2018/02/14 09:47:31 Elapsed Time: 75.0512ms
End creates at Wed, Feb 14, 2018 9:47:31 AM
$
```

Next run load script. This uses the PG Copy API.
```
$ sh load.sh
Start create and import at Wed, Feb 14, 2018 9:55:49 AM
inserting into node table at Wed, Feb 14, 2018 9:55:49 AM
Start Time: 2018-02-14 09:55:50.3517138 -0500 EST m=+0.003000201
Changing all headers to lowercase!
Stop Time: 2018-02-14 09:55:51.6089346 -0500 EST m=+1.260221001
Total run time: 1.2572208s
Inserted 96898 rows
inserting into edge table at Wed, Feb 14, 2018 9:55:51 AM
Start Time: 2018-02-14 09:55:52.5705976 -0500 EST m=+0.002999801
Changing all headers to lowercase!
Stop Time: 2018-02-14 09:56:02.3947349 -0500 EST m=+9.827137101
Total run time: 9.8241373s
Inserted 399959 rows
End create and import at Wed, Feb 14, 2018 9:56:02 AM
$
```

Finally, run the indexing script, using modifydb:
```
$ sh index.sh
Start indexing at Wed, Feb 14, 2018 10:05:05 AM
2018/02/14 10:05:05 Error ioutil.ReadFile() on node_indes.sql:
open node_indes.sql: The system cannot find the file specified.
2018/02/14 10:05:05 SQL is:
create unique index edgeidx on edge(id)
2018/02/14 10:05:09 Total Affected Rows: 0
2018/02/14 10:05:09 Elapsed Time: 3.7941596s
2018/02/14 10:05:09 SQL is:
create index fromidx on edge(from_id)
2018/02/14 10:05:12 Total Affected Rows: 0
2018/02/14 10:05:12 Elapsed Time: 2.7969959s
2018/02/14 10:05:12 SQL is:
create index toidx on edge(to_id)
2018/02/14 10:05:14 Total Affected Rows: 0
2018/02/14 10:05:14 Elapsed Time: 2.5111155s
End indexing at Wed, Feb 14, 2018 10:05:15 AM
$
```

## Original Benchmark
Sample executions:
```
go run pgcopy.go \
  -schema mysch \
  -table mytab \
  -urlvar pgcred \
  -input ./my.csv

Start Time: 2017-01-17 16:38:29.4029527 -0500 EST
Stop Time: 2017-01-17 16:39:11.060118 -0500 EST
Total run time: 41.6571653s
Inserted 7969122 rows
$
```
*Note 1.* The above was an actual run and shows a rate of almost 200K rows/second.

*Note 2.* I recommend just using the TEXT datatype for all columns and correcting to actual datatypes after load into database. This is because of this issue with handling null values for dates (and perhaps other types): https://github.com/lib/pq/issues/591

