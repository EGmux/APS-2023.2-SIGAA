package utils

import (
	"fmt"
	"os"
)

// fn: name of function that called log
func DropTable(fn string, tablename string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from", fn, ":",
		"Failed to drop table",
		tablename,
		"in DB",
		os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn: name of function that called log
func InsertRows(fn string, structname string, tablename string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from", fn, ":",
		"Failed Inserting rows from",
		tablename,
		"in DB",
		os.Getenv("PGDATABASE"),
		"into",
		structname, "ERROR",
		err,
	)
}

// fn: name of function that called log
func ImportSQL(fn string, PATH string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from", fn, ":",
		"Failed importing SQL file",
		PATH,
		"in DB",
		os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn: name of function that called log
func TableExists(fn string, tablename string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from",
		fn,
		":",
		"Table",
		tablename,
		"does not exist in:"+os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn: name of function that called log
func ConnectDB(fn string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from",
		fn,
		":",
		"Error connecting to DB:",
		os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn: name of function that called log
func CommitTransaction(fn string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from",
		fn,
		":",
		"Failed transaction in DB",
		os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn: name of function that called log
func InitTransaction(fn string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from", fn, ":",
		"Failed initiating transaction in DB", os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn: name of function that called log
func CreateTable(fn string, tablename string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from", fn, ":",
		"Can't create table",
		tablename,
		"in DB",
		os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn: name of function that called log
func CloseConnectionDB(fn string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from", fn, ":",
		"Failed close connection to",
		os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn:name of function that called log
func CreateCredentialRepo(fn string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from",
		fn,
		":",
		"Failed to retrieve Credential repo in",
		os.Getenv("PGDATABASE"),
		"ERROR",
		err,
	)
}
