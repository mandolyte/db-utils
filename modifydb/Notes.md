# Notes 

## Setup 
Run the load script:
```
$ sh load.sh 
Start create and import at Sun Feb 11 20:49:04 EST 2018
creates statements...
2018/02/11 20:49:05 SQL is:
CREATE TABLE node (
    id text key
)
2018/02/11 20:49:05 Total Affected Rows: 0
2018/02/11 20:49:05 Elapsed Time: 183.321195ms
2018/02/11 20:49:06 SQL is:
CREATE TABLE edge (
    id text key,
    from_id text not null,
    to_id text not null
)
2018/02/11 20:49:06 Total Affected Rows: 0
2018/02/11 20:49:06 Elapsed Time: 157.148619ms
inserting into node table at Sun Feb 11 20:49:06 EST 2018
2018/02/11 20:49:07 SQL is:
insert into node (id)
values (?)
2018/02/11 20:49:29 Total Affected Rows: 96897
2018/02/11 20:49:29 Elapsed Time: 21.903801574s
inserting into edge table at Sun Feb 11 20:49:29 EST 2018
2018/02/11 20:49:29 SQL is:
insert into edge (id,from_id, to_id)
values (?, ?, ?)
2018/02/11 20:51:21 Total Affected Rows: 399959
2018/02/11 20:51:21 Elapsed Time: 1m51.276099892s
End create and import at Sun Feb 11 20:51:21 EST 2018
$ 
```