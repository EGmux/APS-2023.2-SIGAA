package types

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"sigaa.ufpe/packages/utils/queries"
)

type SQLRepoSitory struct {
	conn *pgxpool.Pool
}

// // TODO: implement
// // If table exits return else create it
// func (repo *sqlrepo) CreateCourseRepo() bool {
// 	return true
// }

// TODO: implement
// Check if table exists if not create it otherwise retrieve it
func (repo SQLRepoSitory) CreateCredentialRepo() bool {
	queries.ConnectDB()
	if queries.TableExists("credentials") {
		queries.InsertRow(
			"id",
			"10",
			"credentials",
			queries.RowElem{ElemName: "passwd", ElemType: "1234"},
			queries.RowElem{ElemName: "username", ElemType: "egb2"},
		)
		tableRows := queries.ReturnRows("credentials")
		for tableRows.Next() {
			println(tableRows.Scan())
		}
		tableRows.Close()
	}
	return false
}

// TODO: implement
// func (repo *sqlrepo) CreatePROAESRepo() bool {
// 	return true
// }
//
// // TODO: implement
// func (repo *sqlrepo) CreateTranscriptRepo() bool {
// 	return true
// }
//
// // TODO: implement
// func (repo *sqlrepo) CreateEnrollmentRepo() bool {
// 	return true
// }
//
// // TODO: implement
// func (repo *sqlrepo) CreateProfessorRepo() bool {
// 	return true
// }
//
// // TODO: implement
// func (repo *sqlrepo) CreateClassRepo() bool {
// 	return true
// }
//
// // TODO: implement
// func (repo *sqlrepo) CreateSubjectRepo() bool {
// 	return true
// }
