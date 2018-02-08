## PGCOPY
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
