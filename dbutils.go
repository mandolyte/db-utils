// Package dbutils runs SQL with the args provided,
// invoking the function provided by the client to process each row
package dbutils

import (
	"database/sql"
	"log"
)

// Dbu is the struct for running the statement and managing the results
type Dbu struct {
	// Db is the opened database
	Db *sql.DB
	// SQL is the SQL query to execute
	SQL string
	// Args is a slice of values to pass into SQL statement
	Args []interface{}
	// RowReader is a user provided function to handle each row
	// returned row from query. Note that the first row returned
	// is the column headers. If this is not needed, then
	// simply discard the results from first use of this function.
	RowReader func([]string) error
	// ColumnReader is a user provided function to handle the column
	// headers of the result set. It will be called prior to any
	// rows given to the RowReader.
	ColumnHeaders []string
	isDataRow bool
}

// Query is the method to execute the query and manage results
func (dbutils *Dbu) Query() error {
	rows, err := dbutils.Db.Query(dbutils.SQL, dbutils.Args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// send the headers on first time in
	if !dbutils.isDataRow {
		dbutils.isDataRow = true
		if dbutils.ColumnHeaders == nil {
			err = dbutils.RowReader(columns)
			if err != nil {
				return err
			}
		} else {
			dbutils.ColumnHeaders = append(dbutils.ColumnHeaders,columns...)
			err = dbutils.RowReader(dbutils.ColumnHeaders)
			if err != nil {
				return err
			}
		}	
	}

	for rows.Next() {
		rcols := make([]interface{}, len(columns))
		wcols := make([]sql.NullString, len(columns))
		for n := range rcols {
			rcols[n] = &wcols[n]
		}
		if err = rows.Scan(rcols...); err != nil {
			log.Fatalf("Error rows.Scan():\n%v\n", err)
		}
		var vals []string
		for _, v := range wcols {
			vals = append(vals, v.String)
		}
		err = dbutils.RowReader(vals)
		if err != nil {
			return err
		}
	}
	return nil
}

// Exec is the method to execute the query and manage results
func (dbutils *Dbu) Exec() (int64, error) {

	result, err := dbutils.Db.Exec(dbutils.SQL, dbutils.Args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
