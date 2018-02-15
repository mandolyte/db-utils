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

In the subfolders find complete tests using cockroachdb and postgresql.