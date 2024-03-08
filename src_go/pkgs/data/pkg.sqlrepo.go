package data

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	util "sigaa.ufpe/pkgs/utils"
)

type SQLRepository struct {
	conn *pgxpool.Pool
	repo Repo
}

// Retrieve each repo, acessible through the IRepo interface
func (s SQLRepository) ReturnRepos() (IRepo, error) {
	var err error
	err = s.CreateCredentialRepo()
	if err != nil {
		return nil, err
	}
	err = s.CreateEnrollmentRepo()
	if err != nil {
		return nil, err
	}
	err = s.CreateProfessorRepo()
	if err != nil {
		return nil, err
	}
	err = s.CreateClassRepo()
	if err != nil {
		return nil, err
	}
	err = s.CreatePROAESRepo()
	if err != nil {
		return nil, err
	}
	err = s.CreateStudentRepo()
	if err != nil {
		return nil, err
	}
	return &s.repo, nil
}

// Fill sqlrepo lists with DB rows after a SELECT ALL query
func createIfExistORReturnSQLRepoT[T SQLTablesPtrs](
	repoName string,
	tableName string,
	repoType *[]*T,
	id string,
	columns ...string,
) error {
	err := errors.New("rerro")
	if err != nil {
		util.DB.ConnectDB()
	}
	var exists bool
	exists, err = util.DB.TableExists(tableName)
	if err != nil {
		util.DB.TableExists(tableName)
	}
	if exists {
		err = util.DB.InsertIntoSQLStruct(repoType, util.SELECT_ALL, tableName, id, columns...)
	}
	if err != nil {
		util.DB.InsertRow(repoName, repoName, tableName)
	}
	return err
}

// Check if table exists if not create it otherwise retrieve it
func (s SQLRepository) CreateCredentialRepo() error {
	err := createIfExistORReturnSQLRepoT(
		"CredentialsSQLRepo",
		"credentials",
		&CredentialSQLRepo,
		"idcredential",
		"passwd",
		"username",
		"logged",
	)
	// Inplace modification, alter the fields in Repo
	credentialRepo, err := s.repo.GetCredentials()
	err = ConvertSQLToList(&CredentialSQLRepo, credentialRepo)
	return err
}

func (s SQLRepository) CreatePROAESRepo() error {
	err := createIfExistORReturnSQLRepoT(
		"PROAESSQLRepo",
		"PROAES",
		&ProaesSQLRepo,
		"idproaes",
		"email",
	)
	// Inplace modification, alter the fields in Repo
	proaesRepo, err := s.repo.GetPROAES()
	err = ConvertSQLToList(&ProaesSQLRepo, proaesRepo)
	return err
}

func (s SQLRepository) CreateEnrollmentRepo() error {
	err := createIfExistORReturnSQLRepoT(
		"EnrollmentSQLRepo",
		"enrollments",
		&EnrollmentSQLRepo,
		"idenrollment",
		"name",
		"status",
	)
	// Inplace modification, alter the fields in Repo
	enrolmmentRepo, err := s.repo.GetEnrollments()
	err = ConvertSQLToList(&EnrollmentSQLRepo, enrolmmentRepo)
	return err
}

func (s SQLRepository) CreateProfessorRepo() error {
	err := createIfExistORReturnSQLRepoT(
		"ProfessorSQL", "professors", &ProfessorSQLRepo, "idprofessor", "name",
	)
	// Inplace modification, alter the fields in Repo
	professorRepo, err := s.repo.GetProfessors()
	err = ConvertSQLToList(&ProaesSQLRepo, professorRepo)
	return err
}

func (s SQLRepository) CreateClassRepo() error {
	err := createIfExistORReturnSQLRepoT(
		"ClassSQLRepo",
		"classes",
		&ClassSQLRepo,
		"id",
		"location",
		"professor",
		"students",
	)
	// Inplace modification, alter the fields in Repo
	classesRepo, err := s.repo.GetClasses()
	err = ConvertSQLToList(&ClassSQLRepo, classesRepo)
	return err
}

func (s SQLRepository) CreateStudentRepo() error {
	err := createIfExistORReturnSQLRepoT(
		"StudentSQLRepo",
		"student",
		&StudentSQLRepo,
		"id",
		"name",
		"transcript",
	)
	// Inplace modification, alter the fields in Repo
	studentRepo, err := s.repo.GetStudents()
	err = ConvertSQLToList(&StudentSQLRepo, studentRepo)
	return err
}
