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
var output = flag.String("output", "", "Output CSV filename")
var urlref = flag.String("urlref", "", "Environment variable with DB URL")
var driver = flag.String("driver", "", "Driver name; required")
var help = flag.Bool("help", false, "Show help message")
var debug = flag.Bool("debug", true, "Show debug messages")

var input = flag.String("input", "", "Optional input CSV supplying parameters to SQL query")
var parameters = flag.String("parameters", "",
	"Comma delimited list of olumn numbers of input CSV in order needed;\nfirst is 1, not zero")

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

	if *driver == "" {
		usage("ERROR: driver is missing\n")
	}

	if *query == "" {
		usage("ERROR: SQL filename for query is missing\n")
	}

	if *output == "" {
		usage("ERROR: Output CSV filename is missing\n")
	}

	// get sql file
	sqlS := fileToString(*query)

	// open database
	db, dberr := sql.Open(*driver, urlvar)
	if dberr != nil {
		log.Fatalf("ERROR: sql.Open() connection failed: %v", dberr)
	}

	// open output file
	fo, foerr := os.Create(*output)
	if foerr != nil {
		log.Fatal("os.Create() Error:" + foerr.Error())
	}
	defer fo.Close()
	w = csv.NewWriter(fo)
	defer w.Flush()

	// setup the basic dbutils struct
	x := &dbutils.Dbu{Db: db, SQL: sqlS, RowReader: theRowReader}

	if *input == "" {
		x.RowReader = theRowReader
		singleQuery(x)
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
		parms := strings.Split(*parameters, ",") // as strings
		// now convert...
		parmcolumns := getParameterColumns(parms)

		// get the column headers from the input CSV (they are required!)
		// by reading the first row from input CSV
		x.ColumnHeaders = getHeaders(r)

		// process input
		multiQuery(x, r, parmcolumns)
	}

	stop := time.Since(now)
	// Row count includes header row
	dbg(fmt.Sprintf("Total Rows: %v\n", rows))
	dbg(fmt.Sprintf("Elapsed Time: %v\n", stop))

}

func getHeaders(r *csv.Reader) []string {
	cells, rerr := r.Read()
	if rerr != nil {
		log.Fatal("r.Read() Error:" + rerr.Error())
	}
	return cells
}

func getParameterColumns(p []string) []int {
	// create ints from parm list
	parmindex := make([]int, len(p))
	for n := range p {
		i, err := strconv.Atoi(p[n])
		if err != nil {
			log.Fatalf("Parameter is not number: %v\n", p[n])
		}
		// account for offset being one-based instead of zero
		parmindex[n] = i - 1
	}
	return parmindex

}

func singleQuery(x *dbutils.Dbu) {
	err := x.Query()
	if err != nil {
		log.Fatalf("Error:%v\n", err)
	}

}

func multiQuery(x *dbutils.Dbu, r *csv.Reader, parms []int) {
	// read loop for CSV
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

		parmvals := make([]interface{}, len(parms))
		for n, v := range parms {
			parmvals[n] = cells[v]
		}
		x.Args = parmvals
		err := x.Query()
		if err != nil {
			log.Fatalf("Error:%v\n", err)
		}
	}
}

// the CSV Writer used by the RowReader
var w *csv.Writer
var cells []string

// Number of rows read by the RowReader
var rows uint64

func theRowReader(aRow []string) error {
	/*
		Note! the first time this is called it will be to write
		out the headers. So the first call must be treated
		differently.
	*/
	if rows == 0 {
		// first time called... this is the header columns only
		// the values in aRow will have all the columns whether
		// for a single SQL or one driven from an input CSV
		err := w.Write(aRow)
		if err != nil {
			log.Fatalf("Error on csv.Write(aRow):\n%v\n", err)
		}
		rows++
		return nil
	}
	// Write out the rows
	var cols []string
	if len(cells) > 0 {
		cols = append(cols, cells...)
	}

	cols = append(cols, aRow...)

	err := w.Write(cols)
	if err != nil {
		log.Fatalf("Error on csv.Write(aRow):\n%v\n", err)
	}

	rows++
	return nil
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
