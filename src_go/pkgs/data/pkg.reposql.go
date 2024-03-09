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
func createIfExistORReturnSQLRepo[T SQLTablesPtrs](
	tableName string,
	repoType *[]*T,
	sqlQuery string,
	fallbackSqlQuery string,
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
		err = util.DB.Query(repoType, sqlQuery)
	}
	if err != nil {
		err = util.DB.Transaction(fallbackSqlQuery)
	}
	return err
}

// Convert data to struct used in the application
func ConvertBetweenSQLandStruct[T SQLTablesPtrs, V TablePtrs](
	r *[]*T,
	v func() (*[]*V, error),
	m func(*V, *T, Direction) error,
	d Direction,
) error {
	credential, err := v()
	if err != nil {
		return errors.New("failed retrieving repo")
	}
	err = ConvertSQLToList(r, credential, m, d)
	if err != nil {
		return errors.New("Failed to Convert SQL to List")
	}
	return nil
}

// Check if table exists if not create it otherwise retrieve it
func (s SQLRepository) CreateCredentialRepo() error {
	err := createIfExistORReturnSQLRepo(
		"credentials",
		&CredentialSQLRepo,
		"SELECT from credentials (idcredential, passwd, username, logged)",
		"INSERT INTO credentials (idcredential, passwd, username) VALUES (0, 'NULL', 'NULL')",
	)
	err = ConvertBetweenSQLandStruct(
		&CredentialSQLRepo,
		s.repo.GetCredentials,
		mappingCredentials,
		SQLTOSTRUCT,
	)
	return err
}

func (s SQLRepository) CreatePROAESRepo() error {
	err := createIfExistORReturnSQLRepo(
		"PROAES",
		&ProaesSQLRepo,
		"SELECT from PROAES (idproaes, email) ",
		"INSERT INTO credentials (idproaes, email) VALUES (0, 'NULL')",
	)
	err = ConvertBetweenSQLandStruct(&ProaesSQLRepo, s.repo.GetPROAES, mappingPROAES, SQLTOSTRUCT)
	return err
}

func (s SQLRepository) CreateEnrollmentRepo() error {
	err := createIfExistORReturnSQLRepo(
		"enrollments",
		&EnrollmentSQLRepo,
		"SELECT from enrollment (idenrollment, name, status) ",
		"INSERT INTO enrollment (idenrollment, name, status) VALUES (0, 'NULL')",
	)
	err = ConvertBetweenSQLandStruct(
		&EnrollmentSQLRepo,
		s.repo.GetEnrollments,
		mappingEnrollment,
		SQLTOSTRUCT,
	)
	return err
}

func (s SQLRepository) CreateProfessorRepo() error {
	err := createIfExistORReturnSQLRepo(
		"professors",
		&ProfessorSQLRepo,
		"SELECT from professors (idprofessor, name) ",
		"INSERT INTO professors (idprofessor, name) VALUES (0, 'NULL')",
	)
	err = ConvertBetweenSQLandStruct(
		&ProfessorSQLRepo,
		s.repo.GetProfessors,
		mappingProfessor,
		SQLTOSTRUCT,
	)
	return err
}

func (s SQLRepository) CreateClassRepo() error {
	err := createIfExistORReturnSQLRepo(
		"classes",
		&ClassSQLRepo,
		"SELECT from classes (idclass, location, professor, students) ",
		"INSERT INTO classes (idclass, location, professor, students) VALUES (0, 'NULL', 'NULL', 'NULL', 'NULL')",
	)
	err = ConvertBetweenSQLandStruct(&ClassSQLRepo, s.repo.GetClasses, mappingClass, SQLTOSTRUCT)
	return err
}

func (s SQLRepository) CreateStudentRepo() error {
	err := createIfExistORReturnSQLRepo(
		"student",
		&StudentSQLRepo,
		"SELECT from students (idstudent, name, transcript) ",
		"INSERT INTO students (idstudent, name, transcript) VALUES (0, 'NULL', 'NULL')",
	)
	err = ConvertBetweenSQLandStruct(
		&StudentSQLRepo,
		s.repo.GetStudents,
		mappingStudent,
		SQLTOSTRUCT,
	)
	return err
}
