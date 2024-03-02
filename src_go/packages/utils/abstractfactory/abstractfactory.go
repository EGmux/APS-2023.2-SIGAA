package abstractfactory

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type sqlrepo SQLRepoSitory

type IAbstractFactory interface {
	GetFactory()
	CreateCourseRepo() bool
	CreateSubjectRepo() bool
	CreateClassRepo() bool
	CreateEnrollmentRepo() bool
	CreateCredentialRepo() bool
	CreateTranscriptRepo() bool
	CreatePROAESRepo() bool
	CreateProfessorRepo() bool
}

type SQLRepoSitory struct {
	conn *pgxpool.Pool
}

// TODO: implement
func (repo *sqlrepo) CreateCourseRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateCredentialRepo() bool {
	return true
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
