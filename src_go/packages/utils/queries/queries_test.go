package queries

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/suite"
)

type QueriesTestSuite struct {
	suite.Suite
	conn *pgx.Conn
	err  error
	tx   pgx.Tx
}

type rowElem struct {
	elemName string
	elemType string
}

// Aux function to init transactions in DB
func (suite *QueriesTestSuite) initTransaction() {
	suite.tx, suite.err = suite.conn.Begin(context.Background())
	if suite.err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Couldn't begin transition in :"+os.Getenv("PGDATABASE"),
			suite.err,
		)
	}
}

// Commit a transaction to DB, can't rollback
func (suite *QueriesTestSuite) commitTransaction() {
	suite.err = suite.tx.Commit(context.Background())
	if suite.err != nil {
		fmt.Fprintln(os.Stderr, "Failed commit in :"+os.Getenv("PGDATABASE"))
	}
}

// Aux table for testing, empty by default, that is no rows are present
func (suite *QueriesTestSuite) createEmptyTable(
	id string,
	idtype string,
	name string,
	args ...rowElem,
) {
	suite.initTransaction()
	header := id + " " + idtype
	for _, arg := range args {
		header += " ," + arg.elemName + " " + arg.elemType
	}
	_, suite.err = suite.tx.Exec(
		context.Background(),
		// BUG: if args is nil, then this transaction always fails
		"create table "+name+" ("+header+")",
	)
	if suite.err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Failed creating table in :"+os.Getenv("PGDATABASE"),
			suite.err,
		)
	}
	suite.commitTransaction()
}

// Insert row in table
func (suite *QueriesTestSuite) insertRow(
	idtype string,
	idval string,
	name string,
	args ...rowElem,
) {
	label := idtype
	vals := idval
	for _, arg := range args {
		label += " ," + arg.elemName
		vals += " ," + "'" + arg.elemType + "'"
	}
	suite.initTransaction()
	_, suite.err = suite.tx.Exec(
		context.Background(),
		"insert into "+name+" ("+label+") values("+vals+")",
	)
	if suite.err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Failed creating table in :"+os.Getenv("PGDATABASE"),
			suite.err,
		)
	}
	suite.commitTransaction()
}

// Drop table
func (suite *QueriesTestSuite) dropTable(name string) {
	suite.initTransaction()
	_, suite.err = suite.tx.Exec(
		context.Background(),
		"drop table if exists "+name,
	)
	if suite.err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Failed dropping table in :"+os.Getenv("PGDATABASE"),
			suite.err,
		)
	}
	suite.commitTransaction()
}

// Aux method for DB connection
func (suite *QueriesTestSuite) connectDB() {
	suite.conn, suite.err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if suite.err != nil {
		fmt.Fprintln(os.Stderr, "Error Conecction to DB:"+os.Getenv("PGDATABASE"), suite.err)
	}
}

func (suite *QueriesTestSuite) SetupSuite() {
	suite.connectDB()
}

func (suite *QueriesTestSuite) SetupTest() {
	suite.createEmptyTable(
		"id",
		"serial primary key",
		"empty",
		rowElem{"test", "int"},
	)
}

func (suite *QueriesTestSuite) TearDownTest() {
	suite.dropTable("empty")
}

func (suite *QueriesTestSuite) TearDownSuite() {
	suite.conn.Close(context.Background())
}

func (suite *QueriesTestSuite) TestTableExists() {
	// Just created table
	justcreated := "empty"
	suite.True(TableExists(justcreated))
	// Table that was never created
	suite.False(TableExists("shouldnotexist"))
	// Table that just had data insertion
	suite.insertRow("id", "13", "empty", rowElem{"test", "12"})
	suite.True(TableExists("empty"))
}

func TestQueriesTestSuite(t *testing.T) {
	suite.Run(t, new(QueriesTestSuite))
}
