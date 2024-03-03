package queries

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	conn *pgx.Conn
	rows pgx.Rows
)

// Check if table exists in DB
func TableExists(name string) bool {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error Connection("+os.Getenv("$PGDATABASE")+"):"+name)
	}
	rows, err = conn.Query(
		context.Background(),
		"SELECT * FROM information_schema.tables"+
			"WHERE table_schema ='PUBLIC' and table_name = "+
			"'"+
			name+
			"'",
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error in Query(TableExists)"+os.Getenv("$PGDATABASE")+name, err)
	}
	// If not an empty array
	for rows.Next() {
		return true
	}
	// Array is empty
	return false
}

// Return all rows pertaining to table
func ReturnRows(name string) pgx.Rows {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed Connection("+os.Getenv("$PGDATABASE")+"):"+name, err)
	}
	rows, err = conn.Query(context.Background(),
		"SELECT * from "+name,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed Query(Return Rows):"+os.Getenv("$PGDATABASE")+":"+name, err)
	}
	return rows
}
