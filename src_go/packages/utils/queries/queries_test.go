package queries

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

type QueriesTestSuite struct {
	suite.Suite
	conn *pgx.Conn
	rows pgx.Rows
	err  error
	tx   pgx.Tx
}

func (suite *QueriesTestSuite) SetupSuite() {
	ConnectDB(suite.conn)
}

func (suite *QueriesTestSuite) SetupTest() {
	CreateEmptyTable(
		"id",
		"serial primary key",
		"empty",
		RowElem{"test", "int"},
	)
}

func (suite *QueriesTestSuite) TearDownTest() {
	DropTable("empty")
}

func (suite *QueriesTestSuite) TearDownSuite() {
	CloseConnectionDB("TearDownSuite")
}

func (suite *QueriesTestSuite) TestTableExists() {
	// Just created table
	justcreated := "empty"
	suite.True(TableExists(justcreated))
	// Table that was never created
	suite.False(TableExists("shouldnotexist"))
	// Table that just had data insertion
	modifiedtable := justcreated
	InsertRow("id", "13", modifiedtable, RowElem{"test", "12"})
	suite.True(TableExists(modifiedtable))
}

func (suite *QueriesTestSuite) TestReturnRows() {
	// Must be zero rows returned
	justcreated := "empty"
	rows = ReturnRows(justcreated)
	rows.Close()
	suite.Equal(rows.CommandTag().RowsAffected(), int64(0))
	// Must be at most 1 row returned
	modifiedtable := justcreated
	InsertRow("id", "12", modifiedtable, RowElem{"test", "13"})
	rows = ReturnRows(modifiedtable)
	// Close tho rows so CommandTag can be read
	rows.Close()
	suite.Equal(rows.CommandTag().RowsAffected(), int64(1))
}

func TestQueriesTestSuite(t *testing.T) {
	suite.Run(t, new(QueriesTestSuite))
}
