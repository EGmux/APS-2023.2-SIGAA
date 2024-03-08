package login_control

import (
	"os"

	"sigaa.ufpe/pkgs/busines/validation"
	"sigaa.ufpe/pkgs/data/repo/structs"
	"sigaa.ufpe/pkgs/utils/queries"
)

func IsValidUser(username string, password string) bool {
	return validation.IsValidUser(username, password)
}

func Init_StudentsDB() {
	queries.ConnectDB()
}

func CloseDB() {
	// NOTE: need to also consider JSON db
	queries.CloseConnectionDB(os.Getenv("PGDATABASE"))
}

func InsertUser(username string, password string) {
	queries.InsertRow(
		"float",
		"0",
		"credentials",
		queries.RowElem{"name", username},
		queries.RowElem{"password", password},
	)
}

func GetAllUsers() []structs.Credentials {
	return nil
}
