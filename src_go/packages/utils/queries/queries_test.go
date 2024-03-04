package queries

import (
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
	"sigaa.ufpe/packages/utils/logs"
)

type QueriesTestSuite struct {
	suite.Suite
	conn *pgx.Conn
}

type testRepo struct {
	param1 string
	param2 string
}

func (suite *QueriesTestSuite) SetupSuite() {
	ConnectDB(suite.conn)
}

func (suite *QueriesTestSuite) SetupTest() {
	err := CreateEmptyTable(
		"id",
		"serial primary key",
		"empty",
		RowElem{"test", "text"},
		RowElem{"test2", "text"},
	)
	if err != nil {
		logs.CreateTable("SetupTest", "empty", err)
	}
	os.Setenv("ON_ERROR_STOP", "true")
}

func (suite *QueriesTestSuite) TearDownTest() {
	err := DropTable("empty")
	if err != nil {
		logs.DropTable("TearDownTest", "empty", err)
	}
	// Drops only if created in test
	err = DropTable("test")
	if err != nil {
		logs.DropTable("TeardDownTest", "empty", err)
	}
	os.Setenv("ON_ERROR_STOP", "false")
}

func (suite *QueriesTestSuite) TearDownSuite() {
	err := CloseConnectionDB("TearDownSuite")
	if err != nil {
		logs.CloseConnectionDB("TearDownTestSuite", err)
	}
}

func (suite *QueriesTestSuite) TestTableExists() {
	// // Just created table
	// justcreated := "empty"
	// suite.True(TableExists(justcreated))
	// // Table that was never created
	// suite.False(TableExists("shouldnotexist"))
	// // Table that just had data insertion
	// modifiedtable := justcreated
	// InsertRow("id", "13", modifiedtable, RowElem{"test", "12"}, RowElem{"test2", "45"})
	// suite.True(TableExists(modifiedtable))
}

// func (suite *QueriesTestSuite) TestReturnRows() {
// 	// Must be zero rows returned
// 	justcreated := "empty"
// 	InsertRows()
// 	rows = ReturnRows(justcreated)
// 	rows.Close()
// 	suite.Equal(rows.CommandTag().RowsAffected(), int64(0))
// 	// Must be at most 1 row returned
// 	modifiedtable := justcreated
// 	InsertRow("id", "12", modifiedtable, RowElem{"test", "13"}, RowElem{"test2", "23"})
// 	rows = ReturnRows(modifiedtable)
// 	// Close tho rows so CommandTag can be read
// 	rows.Close()
// 	suite.Equal(rows.CommandTag().RowsAffected(), int64(1))
// }

func (suite *QueriesTestSuite) TestImportSQL() {
	// // Importing a non existent PATH must fail
	// errMessage, _ := ImportSQL("")
	// suite.Equal(errMessage, "")
	// // Importing a valid PATH must not fail
	// errMessage, _ = ImportSQL("./testdata/test.sql")
	// suite.Equal(errMessage, "CREATE TABLE\n")
	// suite.True(TableExists("test"))
	// // Importing a invalid .sql file must fail
	// errMessage, _ = ImportSQL("./testdata/wrong.sql")
	// suite.Equal(errMessage, "")
}

func TestQueriesTestSuite(t *testing.T) {
	suite.Run(t, new(QueriesTestSuite))
}
