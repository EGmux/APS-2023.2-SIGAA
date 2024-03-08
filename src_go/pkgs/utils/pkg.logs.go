package utils

import (
	"fmt"
	"os"
)

type log struct{}

var logs log

// fn: name of function that called log
func (l *log) DropTable(fn string, tablename string, err error) {
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
func (l *log) InsertRows(fn string, structname string, tablename string, err error) {
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
func (l *log) ImportSQL(fn string, PATH string, err error) {
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
func (l *log) TableExists(fn string, tablename string, err error) {
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
func (l *log) ConnectDB(fn string, err error) {
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
func (l *log) CommitTransaction(fn string, err error) {
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
func (l *log) InitTransaction(fn string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from", fn, ":",
		"Failed initiating transaction in DB", os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn: name of function that called log
func (l *log) CreateTable(fn string, tablename string, err error) {
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
func (l *log) CloseConnectionDB(fn string, err error) {
	fmt.Fprintln(
		os.Stderr,
		"Call from", fn, ":",
		"Failed close connection to",
		os.Getenv("PGDATABASE"), "ERROR",
		err,
	)
}

// fn:name of function that called log
func (l *log) CreateCredentialRepo(fn string, err error) {
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
