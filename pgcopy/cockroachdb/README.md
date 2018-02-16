# PGCOPY

## A complete test using Cockroach DB
The below was done on a Windows 10 laptop with this install:
```
$ ./cockroach.exe start --insecure --host=localhost
*
* WARNING: RUNNING IN INSECURE MODE!
*
* - Your cluster is open for any client that can access localhost.
* - Any user, even root, can log in without providing a password.
* - Any user, connecting as root, can read or write any data in your cluster.
* - There is no network encryption nor authentication, and thus no confidentiality.
*
* Check out how to secure your cluster: https://www.cockroachlabs.com/docs/stable/secure-a-cluster.html
*
CockroachDB node starting at 2018-02-14 15:19:16.5518211 +0000 UTC (took 0.6s)
build:      CCL v1.1.5 @ 2018/02/05 17:44:17 (go1.8.3)
admin:      http://localhost:8080
sql:        postgresql://root@localhost:26257?application_name=cockroach&sslmode=disable
logs:       C:\cockroach-v1.1.5.windows-6.2-amd64\cockroach-data\logs
store[0]:   path=C:\cockroach-v1.1.5.windows-6.2-amd64\cockroach-data
status:     initialized new cluster
clusterID:  bf190cb3-5da2-44ad-b413-524848502958
nodeID:     1
```

First use modifydb to create the tables:
```
$ sh creates.sh
Start creates at Wed, Feb 14, 2018 10:20:18 AM
creates statements...
2018/02/14 10:20:19 SQL is:
create database public
2018/02/14 10:20:20 Total Affected Rows: 0
2018/02/14 10:20:20 Elapsed Time: 1.0317246s
2018/02/14 10:20:20 SQL is:
CREATE TABLE public.node (
    id text primary key
)
2018/02/14 10:20:21 Total Affected Rows: 0
2018/02/14 10:20:21 Elapsed Time: 1.037449s
2018/02/14 10:20:21 SQL is:
CREATE TABLE public.edge (
    id text primary key,
    from_id text not null,
    to_id text not null
)
2018/02/14 10:20:22 Total Affected Rows: 0
2018/02/14 10:20:22 Elapsed Time: 1.0256787s
End creates at Wed, Feb 14, 2018 10:20:22 AM
```

Next run load script. This uses the PG Copy API.
```
$ sh load.sh
Start create and import at Wed, Feb 14, 2018 10:23:39 AM
inserting into node table at Wed, Feb 14, 2018 10:23:39 AM
Start Time: 2018-02-14 10:23:40.046543 -0500 EST m=+0.002999801
Changing all headers to lowercase!
Stop Time: 2018-02-14 10:23:47.1904766 -0500 EST m=+7.146933401
Total run time: 7.1439336s
Inserted 96898 rows
inserting into edge table at Wed, Feb 14, 2018 10:23:47 AM
Start Time: 2018-02-14 10:23:47.3065542 -0500 EST m=+0.003001001
Changing all headers to lowercase!
2018/02/14 10:23:58 Error at row 116377 is:
pq: kv/txn_coord_sender.go:926: transaction is too large to commit: 100349 intents
Args:[/usr/share/code/resources/app/node_modules/jsonfile/CHANGELOG.md /usr/share/code/resources/app/node_modules/jsonfile CHANGELOG.md]
End create and import at Wed, Feb 14, 2018 10:23:58 AM
$
```
*Need to fix above before doing the indexing*

Finally, run the indexing script, using modifydb:
```
$ sh index.sh

```
