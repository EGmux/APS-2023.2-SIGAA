package utils

import (
	"context"
	"os"
	"os/exec"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

type RepositoryConn struct{}

type RowElem struct {
	ElemName string
	ElemType string
}

// func (c *RepositoryConn)tion to init transactions in DB
func (c *RepositoryConn) initTransaction() (pgx.Tx, error) {
	var err error
	tx, err := conn.Begin(context.Background())
	return tx, err
}

// Commit a transaction to DB, can't rollback
func (c *RepositoryConn) commitTransaction(tx pgx.Tx) error {
	err := tx.Commit(context.Background())
	return err
}

// Aux table for testing, empty by default, that is no rows are present
func (c *RepositoryConn) CreateEmptyTable(
	id string,
	idtype string,
	name string,
	args ...RowElem,
) error {
	tx, err := c.initTransaction()
	if err != nil {
		logs.InitTransaction("CreateEmptyTable", err)
	}
	header := id + " " + idtype
	for _, arg := range args {
		header += " ," + arg.ElemName + " " + arg.ElemType
	}
	_, err = tx.Exec(
		context.Background(),
		"create table "+name+" ("+header+")",
	)
	c.commitTransaction(tx)
	if err != nil {
		logs.CommitTransaction("CreateEmptyTable", err)
	}
	return err
}

// Insert row in table
func (c *RepositoryConn) InsertRow(
	idtype string,
	idval string,
	name string,
	args ...RowElem,
) error {
	label := idtype
	vals := idval
	for _, arg := range args {
		label += " ," + arg.ElemName
		vals += " ," + "'" + arg.ElemType + "'"
	}
	tx, err := c.initTransaction()
	if err != nil {
		logs.InitTransaction("InsertRows", err)
	}
	_, err = tx.Exec(
		context.Background(),
		"insert into "+name+" ("+label+") values("+vals+")",
	)
	c.commitTransaction(tx)
	if err != nil {
		logs.CommitTransaction("InsertRows", err)
	}
	return err
}

// Drop table
func (c *RepositoryConn) DropTable(name string) error {
	tx, err := c.initTransaction()
	if err != nil {
		logs.InitTransaction("DropTable", err)
	}
	_, err = tx.Exec(
		context.Background(),
		"drop table if exists "+name,
	)
	c.commitTransaction(tx)
	if err != nil {
		logs.CommitTransaction("DropTable", err)
	}
	return err
}

// Aux method for DB connection
func (c *RepositoryConn) ConnectDB(isTest ...*pgx.Conn) error {
	var err error
	for _, arg := range isTest {
		conn = arg
	}
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	return err
}

// Close connection to DB
func (c *RepositoryConn) CloseConnectionDB(name string) error {
	err := conn.Close(context.Background())
	return err
}

// The below func (c *RepositoryConn)tions are explicitly tested

// Check if table exists in DB
func (c *RepositoryConn) TableExists(name string) (bool, error) {
	rows, err := conn.Query(
		context.Background(),
		"SELECT * FROM information_schema.tables WHERE table_name = '"+name+"'",
	)
	defer rows.Close()
	// If not an empty array
	for rows.Next() {
		return true, err
	}
	// Array is empty
	return false, err
}

// INPLACE, pass ptr address! Return all rows pertaining to table
func (c *RepositoryConn) returnRows(structaddr any, table string, id string, args ...string) error {
	columns := c.convertToColumns(id, args...)
	err := pgxscan.Select(
		context.Background(),
		conn,
		structaddr,
		`SELECT `+columns+` FROM `+table,
	)
	return err
}

// Create new table based on .sql file in disk, PATH can either be absolute or relative
func (c *RepositoryConn) ImportSQL(PATH string) (string, error) {
	out, err := exec.Command("psql", "-d", os.Getenv("PGDATABASE"), "-f", PATH).Output()
	return string(out), err
}

// convert args to columns
func (c *RepositoryConn) convertToColumns(id string, args ...string) string {
	columns := id
	for _, arg := range args {
		columns += " ," + arg + " "
	}
	return columns
}

// INPLACE, pass ptr address! Insert rows into struct based on db query
func (c *RepositoryConn) InsertIntoSQLStruct(
	structptr any,
	tag DBQUERY,
	table string,
	id string,
	args ...string,
) error {
	var err error
	switch tag {
	case SELECT_ALL:
		err = c.returnRows(structptr, table, id, args...)
	}
	return err
}
