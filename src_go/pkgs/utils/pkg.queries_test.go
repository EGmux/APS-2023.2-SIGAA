package utils

import (
	"os"
	// TODO: re-do with default test lib
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

type debug struct {
	droptable  bool
	showerrors bool
}

type QueriesTestSuite struct {
	suite.Suite
	conn *pgx.Conn
	repo []*TestRepo
	opts debug
}

type TestRepo struct {
	Id     string `db:"id"`
	Param1 string `db:"param1"`
	Param2 string `db:"param2"`
}

func (suite *QueriesTestSuite) SetupSuite() {
	DB.ConnectDB(suite.conn)
	suite.opts.droptable = false
	suite.opts.showerrors = true
}

func (suite *QueriesTestSuite) SetupTest() {
	err := DB.Transaction("CREATE TABLE empty (id serial primary key, param1 text, param2 text)")
	if err != nil {
		logs.CreateTable("SetupTest", "empty", err)
	}
	os.Setenv("ON_ERROR_STOP", "true")
}

func (suite *QueriesTestSuite) TearDownTest() {
	var err error
	// Clear the struct
	suite.repo = nil
	if !suite.opts.droptable {
		err = DB.DropTable("empty")
		if err != nil {
			logs.DropTable("TearDownTest", "empty", err)
		}
	}
	// Drops only if created in test
	if !suite.opts.droptable {
		err = DB.DropTable("test")
		if err != nil {
			logs.DropTable("TeardDownTest", "empty", err)
		}
	}
	os.Setenv("ON_ERROR_STOP", "false")
}

func (suite *QueriesTestSuite) TearDownSuite() {
	err := DB.CloseConnectionDB("TearDownSuite")
	if err != nil {
		logs.CloseConnectionDB("TearDownTestSuite", err)
	}
}

func (suite *QueriesTestSuite) TestTableExists() {
	justcreated := "empty"
	exist, err := DB.TableExists(justcreated)
	if err != nil && suite.opts.showerrors {
		logs.TableExists("TestTableExists", "justcreated", err)
	}
	if !suite.True(exist) {
		suite.T().Log("Just created table")
	}
	exist, err = DB.TableExists("shouldnotexist")
	if err != nil && suite.opts.showerrors {
		logs.TableExists("TestTableExists", "shouldnotexist", err)
	}
	if !suite.False(exist) {
		suite.T().Log("Table that was never created")
	}
	modifiedtable := justcreated
	DB.Transaction("INSERT INTO empty (id, param1, param2) VALUES (13, 12, 45)")
	exist, err = DB.TableExists(modifiedtable)
	if err != nil && suite.opts.showerrors {
		logs.InsertRows("TestTableExists", "repo", "modifiedtable", err)
	}
	if !suite.True(exist) {
		suite.T().Log("Table that just had data insertion")
	}
}

func (suite *QueriesTestSuite) TestReturnRows() {
	var err error
	err = DB.Query(&suite.repo, "SELECT (id, param1, param2) from empty")
	if err != nil {
		logs.InsertRows("TestReturnRows", "repo", "justcreated", err)
	}
	if !suite.Len(suite.repo, 0) {
		suite.T().Log("Failed Must be zero rows returned")
	}
	err = DB.Transaction("INSERT INTO  empty(id, param1, param2) values (12,13,23)")
	if err != nil {
		logs.InsertRows("TestReturnRows", "repo", "modifiedtable", err)
	}
	err = DB.Query(&suite.repo, "SELECT (id, param1, param2) from empty")
	if !suite.Len(suite.repo, 1) {
		suite.T().Log("Failed Must be at most 1 row returned")
	}
}

func (suite *QueriesTestSuite) TestImportSQL() {
	errMessage, err := DB.ImportTable("")
	if err != nil && suite.opts.showerrors {
		logs.ImportSQL("TestImportSql", "", err)
	}
	if !suite.Equal(errMessage, "") {
		suite.T().Log("Importing a non exitent PATH must fail")
	}
	errMessage, err = DB.ImportTable("./testdata/test.sql")
	if err != nil && suite.opts.showerrors {
		logs.ImportSQL("TestImportSQL", "./testdata/test.sql", err)
	}
	suite.Equal(errMessage, "CREATE TABLE\n")
	if !suite.True(DB.TableExists("test")) {
		suite.T().Log("Importing a valid PATH must not fail")
	}
	errMessage, err = DB.ImportTable("./testdata/wrong.sql")
	if err != nil && suite.opts.showerrors {
		logs.ImportSQL("TestImportSQL", "./testdata/wrong.sql", err)
	}
	if !suite.Equal(errMessage, "") {
		suite.T().Log("Importing an invalid .sql file must fail")
	}
}

func TestQueriesTestSuite(t *testing.T) {
	suite.Run(t, new(QueriesTestSuite))
}