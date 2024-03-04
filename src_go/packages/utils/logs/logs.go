package logs

import (
	"fmt"
	"os"
)

func DropTable(tablename string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Failed to drop table",
		tablename,
		"in DB",
		os.Getenv("PGDATABASE"),
		err,
	)
}

func InsertRows(structname string, tablename string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Failed Inserting rows from",
		tablename,
		"in DB",
		os.Getenv("PGDATABASE"),
		"into",
		structname,
		err,
	)
}

func ImportSQL(PATH string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Failed importing SQL file",
		PATH,
		"in DB",
		os.Getenv("PGDATABASE"),
		err,
	)
}

func TableExists(tablename string, err error) {
	fmt.Fprintln(os.Stderr, "Table", tablename, "does not exist in:"+os.Getenv("PGDATABASE"), err)
}

func ConnectDB(err error) {
	fmt.Fprintln(os.Stderr, "Error connecting to DB:"+os.Getenv("PGDATABASE"), err)
}

func CommitTransaction(err error) {
	fmt.Fprintln(os.Stderr, "Failed transaction in DB", os.Getenv("PGDATABASE"), err)
}

func InitTransaction(err error) {
	fmt.Fprintln(
		os.Stderr,
		"Failed initiating transaction in DB", os.Getenv("PGDATABASE"),
		err,
	)
}

func CreateTable(tablename string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Can't create table",
		tablename,
		"in DB",
		os.Getenv("PGDATABASE"),
		err,
	)
}
