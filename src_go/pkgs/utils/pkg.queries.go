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

// function to init transactions in DB
func (c *RepositoryConn) initTransaction() (pgx.Tx, error) {
	var err error
	tx, err := conn.Begin(context.Background())
	return tx, err
}

// commit a transaction to DB, can't rollback
func (c *RepositoryConn) commitTransaction(tx pgx.Tx) error {
	err := tx.Commit(context.Background())
	return err
}

// function to execute a new transaction in DB
func (c *RepositoryConn) newTransaction(transaction string) error {
	tx, err := c.initTransaction()
	if err != nil {
		logs.InitTransaction("InsertRows", err)
	}
	_, err = tx.Exec(
		context.Background(),
		transaction,
	)
	c.commitTransaction(tx)
	if err != nil {
		logs.CommitTransaction("InsertRows", err)
	}
	return err
}

// Drop table aux function
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

// NOTE: Below methods are explicitly tested

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
func (c *RepositoryConn) Query(structaddr any, sqlQuery string) error {
	err := pgxscan.Select(
		context.Background(),
		conn,
		structaddr,
		sqlQuery,
	)
	return err
}

// INPLACE, pass ptr address! Insert all rows pertaining to table
func (c *RepositoryConn) Transaction(sqlQuery string) error {
	err := c.newTransaction(sqlQuery)
	return err
}

// Create new table based on .sql file in disk, PATH can either be absolute or relative
func (c *RepositoryConn) ImportTable(PATH string) (string, error) {
	out, err := exec.Command("psql", "-d", os.Getenv("PGDATABASE"), "-f", PATH).Output()
	return string(out), err
}
