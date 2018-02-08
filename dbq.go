// Package dbq runs SQL with the args provided,
// invoking the function provided by the client to process each row
package dbq

import (
	"database/sql"
	"log"
)

// Dbq is the struct for running the query and managing the results
type Dbq struct {
	// Db is the opened database 
	Db        *sql.DB
	// SQL is the SQL query to execute
	SQL       string
	// Args is a slice of values to pass into SQL statement
	Args      []interface{}
	// RowReader is a user provided function to handle each row
	// returned row from query. Note that the first row returned
	// is the column headers. If this is not needed, then
	// simply discard the results from first use of this function.
	RowReader func([]string) error
}

// Query is the method to execute the query and manage results
func (dbq *Dbq) Query() error {

	rows, err := dbq.Db.Query(dbq.SQL, dbq.Args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// send the headers
	err = dbq.RowReader(columns)
	if err != nil {
		return err
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
		err = dbq.RowReader(vals)
		if err != nil {
			return err
		}
	}
	return nil
}
