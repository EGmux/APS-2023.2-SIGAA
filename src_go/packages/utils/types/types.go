package types

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"sigaa.ufpe/packages/utils/logs"
	"sigaa.ufpe/packages/utils/queries"
)

// Aliases
type serial int64

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

func (repo SQLRepoSitory) createIfExistORReturnRepo(
	repoName string,
	tableName string,
	repoType interface{},
	id string,
	columns ...string,
) (interface{}, error) {
	err := queries.ConnectDB()
	if err != nil {
		logs.ConnectDB(repoName, err)
	}
	var exists bool
	exists, err = queries.TableExists(tableName)
	if err != nil {
		logs.TableExists(repoName, tableName, err)
	}
	if exists {
		switch t := repoType.(type) {
		case []*CredentialsRepo:
			err = queries.InsertIntoStruct(&t, queries.SELECT_ALL, tableName, id, columns...)
			return t, err
		case []*PROAESRepo:
			err = queries.InsertIntoStruct(&t, queries.SELECT_ALL, tableName, id, columns...)
			return t, err
		case []*StudentRepo:
			err = queries.InsertIntoStruct(&t, queries.SELECT_ALL, tableName, id, columns...)
			return t, err
		case []*EnrollmentRepo:
			err = queries.InsertIntoStruct(&t, queries.SELECT_ALL, tableName, id, columns...)
			return t, err
		case []*ProfessorRepo:
			err = queries.InsertIntoStruct(&t, queries.SELECT_ALL, tableName, id, columns...)
			return t, err
		case []*ClassRepo:
			err = queries.InsertIntoStruct(&t, queries.SELECT_ALL, tableName, id, columns...)
			return t, err
		case []*SubjectRepo:
			err = queries.InsertIntoStruct(&t, queries.SELECT_ALL, tableName, id, columns...)
			return t, err
		}
		if err != nil {
			logs.InsertRows(repoName, repoName, tableName, err)
		}
	}
	return nil, nil
}

// Check if table exists if not create it otherwise retrieve it
func (repo SQLRepoSitory) CreateCredentialRepo() ([]*CredentialsRepo, error, bool) {
	returnedrepo, err := repo.createIfExistORReturnRepo(
		"CredentialsRepo",
		"credentials",
		credentialRepo,
		"id",
		"passwd",
		"username",
	)
	t, ok := returnedrepo.([]*CredentialsRepo)
	return t, err, ok
}

func (repo SQLRepoSitory) CreatePROAESRepo() ([]*PROAESRepo, error, bool) {
	returnedrepo, err := repo.createIfExistORReturnRepo(
		"PROAESRepo",
		"PROAES",
		proaesRepo,
		"id",
		"name",
	)
	t, ok := returnedrepo.([]*PROAESRepo)
	return t, err, ok
}

func (repo SQLRepoSitory) CreateTranscriptRepo() ([]*StudentRepo, error, bool) {
	returnedrepo, err := repo.createIfExistORReturnRepo(
		"StudentRepo",
		"students",
		studentRepo,
		"id",
		"name",
		"transcript",
	)
	t, ok := returnedrepo.([]*StudentRepo)
	return t, err, ok
}

func (repo SQLRepoSitory) CreateEnrollmentRepo() ([]*EnrollmentRepo, error, bool) {
	returnedrepo, err := repo.createIfExistORReturnRepo(
		"EnrollmentRepo",
		"enrollments",
		enrollmentRepo,
		"id",
		"name",
		"status",
	)
	t, ok := returnedrepo.([]*EnrollmentRepo)
	return t, err, ok
}

func (repo SQLRepoSitory) CreateProfessorRepo() ([]*ProfessorRepo, error, bool) {
	returnedrepo, err := repo.createIfExistORReturnRepo(
		"ProfessorRepo", "professors", professorRepo, "id", "name",
	)
	t, ok := returnedrepo.([]*ProfessorRepo)
	return t, err, ok
}

func (repo SQLRepoSitory) CreateClassRepo() ([]*ClassRepo, error, bool) {
	returnedrepo, err := repo.createIfExistORReturnRepo(
		"ClassRepo",
		"classes",
		classRepo,
		"id",
		"location",
		"professor",
		"students",
	)
	t, ok := returnedrepo.([]*ClassRepo)
	return t, err, ok
}

func (repo SQLRepoSitory) CreateSubjectRepo() ([]*SubjectRepo, error, bool) {
	returnedrepo, err := repo.createIfExistORReturnRepo(
		"SubjectRepo",
		"subjects",
		subjectRepo,
		"id",
		"class",
		"content",
		"deps",
	)
	t, ok := returnedrepo.([]*SubjectRepo)
	return t, err, ok
}
