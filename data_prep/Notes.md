# Notes 

## Hierarchical Dataset Preparation
First, use some of the main Linux folders to get a "file report"
using https://github.com/mandolyte/FileReporter.
```
$ FileReporter /bin hier_bin.csv
$ FileReporter /dev hier_dev.csv
$ FileReporter /lib64 hier_lib64.csv
$ FileReporter /boot hier_boot.csv
$ FileReporter /etc hier_etc.csv
$ FileReporter /opt hier_opt.csv
$ FileReporter /usr hier_usr.csv
$ FileReporter /lib hier_lib.csv
$ FileReporter /sbin hier_sbin.csv
$ FileReporter /sys hier_sys.csv
$ FileReporter /var hier_var.csv
$ ls -al hier*.csv
-rw-r--r-- 1 cecil cecil     5272 Feb  4 17:14 hier_bin.csv
-rw-r--r-- 1 cecil cecil    23930 Feb  4 17:14 hier_boot.csv
-rw-r--r-- 1 cecil cecil    21381 Feb  4 17:14 hier_dev.csv
-rw-r--r-- 1 cecil cecil   184277 Feb  4 17:15 hier_etc.csv
-rw-r--r-- 1 cecil cecil       99 Feb  4 17:14 hier_lib64.csv
-rw-r--r-- 1 cecil cecil  5661623 Feb  4 17:17 hier_lib.csv
-rw-r--r-- 1 cecil cecil    22709 Feb  4 17:15 hier_opt.csv
-rw-r--r-- 1 cecil cecil     8411 Feb  4 17:17 hier_sbin.csv
-rw-r--r-- 1 cecil cecil  3231132 Feb  4 17:18 hier_sys.csv
-rw-r--r-- 1 cecil cecil 33897287 Feb  4 17:18 hier_usr.csv
-rw-r--r-- 1 cecil cecil  1110104 Feb  4 17:18 hier_var.csv
$ catcsv -o hier_dataset.csv hier_*.csv
2018/02/04 17:19:25 Individual file row counts include header row
2018/02/04 17:19:25 Total row count does not include header rows
2018/02/04 17:19:25 File hier_bin.csv had 163 rows
2018/02/04 17:19:25 File hier_boot.csv had 326 rows
2018/02/04 17:19:25 File hier_dev.csv had 533 rows
2018/02/04 17:19:25 File hier_etc.csv had 2482 rows
2018/02/04 17:19:25 File hier_lib64.csv had 2 rows
2018/02/04 17:19:25 File hier_lib.csv had 38465 rows
2018/02/04 17:19:25 File hier_opt.csv had 205 rows
2018/02/04 17:19:25 File hier_sbin.csv had 224 rows
2018/02/04 17:19:25 File hier_sys.csv had 26001 rows
2018/02/04 17:19:27 File hier_usr.csv had 256494 rows
2018/02/04 17:19:27 File hier_var.csv had 9695 rows
2018/02/04 17:19:27 Total rows in output hier_dataset.csv has 334579 rows
$ 
```

Next, split out just the last column which is the full path.
```
$ splitcsv -c 5 < hier_dataset.csv | sort -u > hier_base.csv
$ wc -l hier_base.csv 
334580 hier_base.csv
$ # remove the "Fullpath" header!
```

NOTE: the above can be run with the `frun.sh` script.

For next step, I'll create two output files, one for a node table and one 
for an edge table.

The node table will have a single column, being the name of a file or directory.

This approach will cause same-named files (not dirs) to be treated as the same node.
This will, hopefully, result in some fair number of items having 
multiple parents. Thus a nicer graph dataset.

The edge table will look like this:
- an ID, being a concatenation of the below. 
This probably will not be used much in the current testing.
- a from_id being a directory node (fully specified)
- a to_id being a subdirectory node (fully specified) or a filename.

This may result in multiple edges between two nodes... 
that's a good thing for a graph dataset.

Run it:
```
$ time go run hier_data_prep.go \
    -i $HOME/data/hier/hier_base.csv \
    -node $HOME/data/hier/hier_nodes.csv \
    -edge $HOME/data/hier/hier_edges.csv

Processed 334579 full paths.
Number of nodes: 96898.
Number of edges: 399959.

real	0m4.879s
user	0m4.138s
sys	0m1.028s

```

## Entitlement Benchmarks
Assuming that noting can be faster than SQLite3 itself!

Run the load script which creates tables, imports data, and creates indexes.
```
$ sh load.sh 
Start at Tue Feb 6 07:49:05 EST 2018
End at Tue Feb 6 07:49:13 EST 2018
$ ls -al hier.db
-rw-r--r-- 1 cecil cecil 132853760 Feb  6 07:49 hier.db
$ 
```

Run some queries:
```
-- some counts
sqlite> select count(*) from node;
96898
sqlite> select count(*) from edge;
399959
sqlite> 
-- joins grouping
select n.id, count(e.id)
from node n
inner join edge e
on n.id = e.from_id
group by n.id
order by count(e.from_id) desc
limit 5
;
/usr/share/doc|1713
/usr/src/linux-headers-4.13.0-25-generic/include/config|1663
/usr/src/linux-headers-4.13.0-31-generic/include/config|1663
/usr/src/linux-headers-4.13.0-32-generic/include/config|1663
/usr/src/linux-headers-4.13.0-16-generic/include/config|1662
sqlite> 
sqlite> 
```

```
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
;

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
;

.headers on
.mode column

max(level)  count(*)  
----------  ----------
11          8739      
sqlite> 

```

## Queries
Run q1.sh:
```
$ time sh q1.sh 

real	0m0.718s
user	0m0.000s
sys	0m0.125s
```

Run q2.sh:
```
$ time sh q2.sh 

real	0m12.404s
user	0m0.419s
sys	0m1.000s
$ 
```

Run q3.sh
```
$ time sh q3.sh

real	0m0.012s
user	0m0.009s
sys	0m0.003s
```

Run q4.sh
```
$ time sh q4.sh

real	0m0.067s
user	0m0.058s
sys	0m0.007s
```
