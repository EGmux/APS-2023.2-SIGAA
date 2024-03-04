package queries

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/jackc/pgx/v5"
)

var (
	conn *pgx.Conn
	rows pgx.Rows
	tx   pgx.Tx
	err  error
)

type RowElem struct {
	ElemName string
	ElemType string
}

// Function to init transactions in DB
func initTransaction() {
	tx, err = conn.Begin(context.Background())
	if err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Couldn't begin transaction in :"+os.Getenv("PGDATABASE"),
			err,
		)
	}
}

// Commit a transaction to DB, can't rollback
func commitTransaction() {
	err = tx.Commit(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed commit in :"+os.Getenv("PGDATABASE"))
	}
}

// Aux table for testing, empty by default, that is no rows are present
func CreateEmptyTable(
	id string,
	idtype string,
	name string,
	args ...RowElem,
) {
	initTransaction()
	header := id + " " + idtype
	for _, arg := range args {
		header += " ," + arg.ElemName + " " + arg.ElemType
	}
	_, err = tx.Exec(
		context.Background(),
		"create table "+name+" ("+header+")",
	)
	if err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Failed creating table in :"+os.Getenv("PGDATABASE"),
			err,
		)
	}
	commitTransaction()
}

// Insert row in table
func InsertRow(
	idtype string,
	idval string,
	name string,
	args ...RowElem,
) {
	label := idtype
	vals := idval
	for _, arg := range args {
		label += " ," + arg.ElemName
		vals += " ," + "'" + arg.ElemType + "'"
	}
	initTransaction()
	_, err = tx.Exec(
		context.Background(),
		"insert into "+name+" ("+label+") values("+vals+")",
	)
	if err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Failed creating table in :"+os.Getenv("PGDATABASE"),
			err,
		)
	}
	commitTransaction()
}

// Drop table
func DropTable(name string) {
	initTransaction()
	_, err = tx.Exec(
		context.Background(),
		"drop table if exists "+name,
	)
	if err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Failed dropping table in :"+os.Getenv("PGDATABASE"),
			err,
		)
	}
	commitTransaction()
}

// Aux method for DB connection
func ConnectDB(isTest ...*pgx.Conn) {
	for _, arg := range isTest {
		conn = arg
	}
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error Conecction to DB:"+os.Getenv("PGDATABASE"), err)
	}
}

// Close connection to DB
func CloseConnectionDB(name string) {
	err = conn.Close(context.Background())
	if err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Can't close connection to DB:"+os.Getenv("PGDATABASE")+":"+name,
			err,
		)
	}
}

// The below functions are explicitly tested

// Check if table exists in DB
func TableExists(name string) bool {
	rows, err = conn.Query(
		context.Background(),
		"SELECT * FROM information_schema.tables WHERE table_name = '"+name+"'",
	)
	defer rows.Close()
	if err != nil {

		fmt.Fprintln(os.Stderr, "Error in Query(TableExists)"+os.Getenv("$PGDATABASE")+name, err)
		fmt.Fprintln(os.Stderr, "Error in Query(TableExists)"+os.Getenv("PGDATABASE")+":"+name, err)
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
	rows, err = conn.Query(context.Background(),
		"SELECT * from "+name,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed Query(Return Rows):"+os.Getenv("PGDATABASE")+":"+name, err)
	}
	return rows
}

// Create new table based on .sql file in disk, PATH can either be absolute or relative
func ImportSQL(PATH string) string {
	out, err := exec.Command("psql", "-d", os.Getenv("PGDATABASE"), "-f", PATH).Output()
	if err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Failed while running IMPORTSQL:"+os.Getenv("PGDATABASE")+":"+PATH,
			err,
		)
	}
	return string(out)
}
