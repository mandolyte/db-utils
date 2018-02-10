package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mandolyte/db-utils"

	_ "github.com/cznic/sqlite"
)

// global variables - beware!
var query = flag.String("query", "", "SQL statement filename")
var urlref = flag.String("urlref", "", "Environment variable with DB URL")
var driverName = flag.String("driverName", "", "Driver name; required")
var help = flag.Bool("help", false, "Show help message")
var debug = flag.Bool("debug", true, "Show debug messages")

var input = flag.String("input", "", "Optional input CSV supplying parameters to SQL query")
var parameters = flag.String("parameters", "",
	"Comma delimited list of column numbers of input CSV in order needed;\nfirst is 1, not zero")

func main() {
	now := time.Now().UTC()
	flag.Parse()
	if *help {
		usage("Help:")
	}

	if *urlref == "" {
		usage("ERROR: Environment variable with DB URL is missing\n")
	}
	urlvar := os.Getenv(*urlref)
	if urlvar == "" {
		usage(fmt.Sprintf("ERROR: missing URL from os.Getenv('%v')\n", *urlref))
	}

	if *driverName == "" {
		usage("ERROR: driverName is missing\n")
	}

	if *query == "" {
		usage("ERROR: SQL filename for query is missing\n")
	}

	// get sql file
	sqlS := fileToString(*query)

	// open database
	db, dberr := sql.Open(*driverName, urlvar)
	if dberr != nil {
		log.Fatalf("ERROR: sql.Open() connection failed: %v", dberr)
	}

	// create the Dbq struct...
	x := &dbq.Dbq{Db: db, SQL: sqlS}

	if *input == "" {
		singleExec(db, sqlS, theRowReader, theHeadersReader)
	} else {
		if *parameters == "" {
			usage("Parameters for input CSV are missing")
		}
		// open input file
		fi, fierr := os.Open(*input)
		if fierr != nil {
			log.Fatal("os.Open() Error:" + fierr.Error())
		}
		defer fi.Close()
		// associate to CSV
		r := csv.NewReader(fi)
		// ignore expectations of fields per row
		r.FieldsPerRecord = -1

		// get parameter column numbers
		parms := strings.Split(*parameters, ",")

		// process input
		multiExec(db, sqlS, theRowReader, theHeadersReader, r, parms)
	}

	stop := time.Since(now)
	dbg(fmt.Sprintf("Total Affected Rows: %v\n", rows))
	dbg(fmt.Sprintf("Elapsed Time: %v\n", stop))

}

func singleExec(db *sql.DB, sqlS string, theReader, headers func([]string) error) {
	err := x.Query()
	if err != nil {
		log.Fatalf("Error:%v\n", err)
	}

}

func multiExec(db *sql.DB, sqlS string, theReader, colHeaders func([]string) error, r *csv.Reader, parms []string) {
	// create ints from parm list
	parmindex := make([]int, len(parms))
	for n := range parms {
		i, err := strconv.Atoi(parms[n])
		if err != nil {
			log.Fatalf("Parameter is not number: %v\n", parms[n])
		}
		// account for offset being one-based instead of zero
		parmindex[n] = i - 1
	}

	// read loop for CSV
	firstDataRow := true
	headerRow := true
	var rerr error
	for {
		// read the csv file
		cells, rerr = r.Read()
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			log.Fatal("r.Read() Error:" + rerr.Error())
		}
		if headerRow {
			// don't use the first (header) row as data
			headerRow = false
			saveHeaders = cells
			continue
		}

		parmvals := make([]interface{}, len(parmindex))
		for n, v := range parmindex {
			parmvals[n] = cells[v]
		}

		if firstDataRow {
			firstDataRow = false
			x := &dbq.Dbq{Db: db, SQL: sqlS, RowReader: theReader, ColumnReader: colHeaders, Args: parmvals}
			err := x.Query()
			if err != nil {
				log.Fatalf("Error:%v\n", err)
			}
			continue
		}

		x := &dbq.Dbq{Db: db, SQL: sqlS, RowReader: theReader, ColumnReader: nil, Args: parmvals}
		err := x.Query()
		if err != nil {
			log.Fatalf("Error:%v\n", err)
		}
	}
}

func usage(msg string) {
	fmt.Println(msg)
	fmt.Println("Usage:")
	flag.PrintDefaults()
	fmt.Println(doc)
	fmt.Println()
	os.Exit(0)
}

func dbg(msg string) {
	if *debug {
		log.Print(msg)
	}
}

func fileToString(filename string) string {
	sqlBytes, serr := ioutil.ReadFile(filename)
	if serr != nil {
		log.Fatalf("Error ioutil.ReadFile() on %v:\n%v\n", filename, serr)
	}
	sqlstmt := strings.TrimSpace(string(sqlBytes))
	dbg(fmt.Sprintf("SQL is:\n%v\n", sqlstmt))
	return sqlstmt
}

var doc = `
Notes:
1. If the optional input CSV is supplied, then the rows will be used to supplay
parameter values to the SQL statement. 
2. The input CSV must be accompanied by a list of column numbers in the order
needed to correctly drive the SQL parameter substitution. 
	For example, if the WHERE clause is:
		WHERE x = ? and (y = ? or z = ?)
	and the values in the CSV for x, y, and z are found in columns 
	2, 12, and 3, then the parameters argument will look like this:
		-parameters 2,12,3
3. The SQL statement will be executed one time for each row in the CSV.
4. The list of parameters is one-based like SQL, not zero based. 
So first column is one, not zero
`
