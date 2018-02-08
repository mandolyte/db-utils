package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mandolyte/db-utils"

	_ "github.com/cznic/sqlite"
)

// global variables - beware!
var query = flag.String("query", "", "SQL statement filename")
var output = flag.String("output", "", "Output CSV filename")
var urlref = flag.String("urlref", "", "Environment variable with DB URL")
var driverName = flag.String("driverName", "", "Driver name; required")
var help = flag.Bool("help", false, "Show help message")
var debug = flag.Bool("debug", true, "Show debug messages")

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

	if *output == "" {
		usage("ERROR: Output CSV filename is missing\n")
	}

	// get sql file
	sqlS := fileToString(*query)

	// open database
	db, dberr := sql.Open(*driverName, urlvar)
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


	x := &dbq.Dbq{Db: db, SQL: sqlS, RowReader: theReader}
	err := x.Query()
	if err != nil {
		log.Fatalf("Error:%v\n", err)
	}

	stop := time.Since(now)
	dbg(fmt.Sprintf("Total Rows: %v\n", rows))
	dbg(fmt.Sprintf("Elapsed Time: %v\n", stop))

}

// the CSV Writer used by the RowReader
var w *csv.Writer

// Number of rows read by the RowReader
var rows uint64

func theReader(aRow []string) error {
	// Write out the rows (including the first one, the header row)
	err := w.Write(aRow)
	if err != nil {
		log.Fatalf("Error on csv.Write(aRow):\n%v\n", err)
	}

	rows++
	return nil
}

func usage(msg string) {
	fmt.Println(msg)
	flag.PrintDefaults()
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
