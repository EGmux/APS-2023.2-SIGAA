package types

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"sigaa.ufpe/packages/utils/queries"
)

type SQLRepoSitory struct {
	conn *pgxpool.Pool
}

type CredentialsRepo struct {
	Id       string `db:"id"`
	Passwd   string `db:"passwd"`
	Username string `db:"username"`
}

type PROAESRepo struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type StudentRepo struct {
	Id         string `db:"id"`
	Name       string `db:"name"`
	Transcript string `db:"transcript"`
}

type EnrollmentRepo struct {
	Id     string `db:"id"`
	Name   string `db:"name"`
	Status string `db:"status"`
}

type ProfessorRepo struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type ClassRepo struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	Professor string `db:"professor"`
	Students  string `db:"students"`
}

type SubjectRepo struct {
	Id      string `db:"id"`
	Name    string `db:"name"`
	Content string `db:"content"`
	Deps    string `db:"deps"`
}

var (
	credentialRepo []*CredentialsRepo
	proaesRepo     []*PROAESRepo
	studentRepo    []*StudentRepo
	enrollmentRepo []*EnrollmentRepo
	professorRepo  []*ProfessorRepo
	classRepo      []*ClassRepo
	subjectRepo    []*SubjectRepo
)
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
