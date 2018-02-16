# Notes 


## Windows 10 experience
First, create the tables. Note that the cznic/sqlite package
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



