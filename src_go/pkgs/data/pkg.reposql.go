package data

import (
	"fmt"

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
	path string,
) error {
	var err error
	err = util.DB.ConnectDB()
	if err != nil {
		return errors.New("Fail to connectDB")
	}
	var exists bool
	exists, err = util.DB.TableExists(tableName)
	if err != nil {
		errors.New("Error table existence")
	}
	if exists {
		err = util.DB.Query(repoType, sqlQuery)
	} else {
		if err != nil {
			out, err := util.DB.ImportTable(path)
			fmt.Printf(out)
			return err
		}
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
	err = MapSQLStruct(r, credential, m, d)
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
		"./assets/credential.sql",
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
		"./assets/proaes.sql",
	)
	err = ConvertBetweenSQLandStruct(&ProaesSQLRepo, s.repo.GetPROAES, mappingPROAES, SQLTOSTRUCT)
	return err
}

func (s SQLRepository) CreateEnrollmentRepo() error {
	err := createIfExistORReturnSQLRepo(
		"enrollments",
		&EnrollmentSQLRepo,
		"SELECT from enrollment (idenrollment, name, status) ",
		"./assets/enrollment.sql",
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
		"./assets/professor.sql",
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
		"./assets/class.sql",
	)
	err = ConvertBetweenSQLandStruct(&ClassSQLRepo, s.repo.GetClasses, mappingClass, SQLTOSTRUCT)
	return err
}

func (s SQLRepository) CreateStudentRepo() error {
	err := createIfExistORReturnSQLRepo(
		"student",
		&StudentSQLRepo,
		"SELECT from students (idstudent, name, transcript) ",
		"./assets/students.sql",
	)
	err = ConvertBetweenSQLandStruct(
		&StudentSQLRepo,
		s.repo.GetStudents,
		mappingStudent,
		SQLTOSTRUCT,
	)
	return err
}
