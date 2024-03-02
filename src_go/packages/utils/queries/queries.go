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

func CreateTable(){
	
}


// Check if table exists in DB
<<<<<<< HEAD
func TableExists(name string) bool {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error Connection("+os.Getenv("PGDATABASE")+":"+name)
	}
	rows, err = conn.Query(
		context.Background(),
		// WARN: Remember to space end of each string
		"SELECT * FROM information_schema.tables WHERE table_name = '"+name+"'",
	)
	if err != nil {
<<<<<<< HEAD
<<<<<<< HEAD

		fmt.Fprintln(os.Stderr, "Error in Query(TableExists)"+os.Getenv("$PGDATABASE")+name, err)
=======
		fmt.Fprintln(os.Stderr, "Error in Query(TableExists)"+os.Getenv("PGDATABASE")+":"+name, err)
>>>>>>> c14c23d (fix(queries:TableExists): SQL query)
=======
		fmt.Fprintln(os.Stderr, "Error in Query(TableExists)"+os.Getenv("PGDATABASE")+":"+name, err)
=======
		fmt.Fprintln(os.Stderr, "Error in Query(TableExists)"+os.Getenv("$PGDATABASE")+name, err)
=======
func CheckExistence(name string) bool {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed connection to DB:Table:"+name, err)
	}
	rows, err = conn.Query(
		context.Background(),
		"SELECT * FROM information_schema.tables WHERE table_schema ='PUBLIC' and table_name = "+"'"+name+"'",
		// "WHERE table_schema LIKE 'public'"+
		// "AND table_type LIKE 'BASE TABLE'"+
		// "table_name = 'credentials'",
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR HERE:", err)
>>>>>>> fd23a1e (add(packages:queries): queries.go)
>>>>>>> 3327b58 (add(packages:queries): queries.go)
>>>>>>> 8b15822 (add(packages:queries): queries.go)
	}
	// If not an empty array
	for rows.Next() {
		return true
	}
	// Array is empty
	return false
}
<<<<<<< HEAD


// Return all rows pertaining to table
func ReturnRows(name string) pgx.Rows {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed Connection("+os.Getenv("PGDATABASE")+":"+name, err)
	}
	rows, err = conn.Query(context.Background(),
		"SELECT * from "+name,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed Query(Return Rows):"+os.Getenv("PGDATABASE")+":"+name, err)
	}
	return rows
}
=======
>>>>>>> fd23a1e (add(packages:queries): queries.go)
