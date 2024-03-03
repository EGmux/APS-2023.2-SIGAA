package types

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"sigaa.ufpe/packages/utils/queries"
)

type sqlrepo SQLRepoSitory

type SQLRepoSitory struct {
	conn *pgxpool.Pool
}

// TODO: implement
// If table exits return else create it
func (repo *sqlrepo) CreateCourseRepo() bool {
	return true
}

// TODO: implement
// Check if table exists if not create it otherwise retrieve it
func (repo sqlrepo) CreateCredentialRepo() *sqlrepo {
	if queries.TableExists("credentials") {
		tableRows := queries.ReturnRows("credentials")
	}
}

// TODO: implement
func (repo *sqlrepo) CreatePROAESRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateTranscriptRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateEnrollmentRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateProfessorRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateClassRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateSubjectRepo() bool {
	return true
}
