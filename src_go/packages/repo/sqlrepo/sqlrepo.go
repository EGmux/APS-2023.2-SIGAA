package sqlrepo

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"sigaa.ufpe/packages/repo"
	"sigaa.ufpe/packages/repo/structs/sqlstructs"
	"sigaa.ufpe/packages/utils"
	"sigaa.ufpe/packages/utils/logs"
	"sigaa.ufpe/packages/utils/queries"
)

type SQLRepository struct {
	conn *pgxpool.Pool
	repo repo.Repo
}

// Retrieve each repo, acessible through the IRepo interface
func (s SQLRepository) ReturnRepos() (repo.IRepo, error) {
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
func createIfExistORReturnSQLRepo[T sqlstructs.SQLTablesPtrs](
	repoName string,
	tableName string,
	repoType *[]*T,
	id string,
	columns ...string,
) error {
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
		err = queries.InsertIntoSQLStruct(repoType, queries.SELECT_ALL, tableName, id, columns...)
	}
	if err != nil {
		logs.InsertRows(repoName, repoName, tableName, err)
	}
	return err
}

// Check if table exists if not create it otherwise retrieve it
func (s SQLRepository) CreateCredentialRepo() error {
	err := createIfExistORReturnSQLRepo(
		"CredentialsSQLRepo",
		"credentials",
		&sqlstructs.CredentialSQLRepo,
		"idcredential",
		"passwd",
		"username",
		"logged",
	)
	// Inplace modification, alter the fields in Repo
	credentialRepo, err := s.repo.GetCredentials()
	err = utils.ConvertSQLToList(&sqlstructs.CredentialSQLRepo, credentialRepo)
	return err
}

func (s SQLRepository) CreatePROAESRepo() error {
	err := createIfExistORReturnSQLRepo(
		"PROAESSQLRepo",
		"PROAES",
		&sqlstructs.ProaesSQLRepo,
		"idproaes",
		"email",
	)
	// Inplace modification, alter the fields in Repo
	proaesRepo, err := s.repo.GetPROAES()
	utils.ConvertSQLToList(&sqlstructs.ProaesSQLRepo, proaesRepo)
	return err
}

func (s SQLRepository) CreateEnrollmentRepo() error {
	err := createIfExistORReturnSQLRepo(
		"EnrollmentSQLRepo",
		"enrollments",
		&sqlstructs.EnrollmentSQLRepo,
		"idenrollment",
		"name",
		"status",
	)
	// Inplace modification, alter the fields in Repo
	enrolmmentRepo, err := s.repo.GetEnrollments()
	utils.ConvertSQLToList(&sqlstructs.EnrollmentSQLRepo, enrolmmentRepo)
	return err
}

func (s SQLRepository) CreateProfessorRepo() error {
	err := createIfExistORReturnSQLRepo(
		"ProfessorSQLRepo", "professors", &sqlstructs.ProfessorSQLRepo, "idprofessor", "name",
	)
	// Inplace modification, alter the fields in Repo
	professorRepo, err := s.repo.GetProfessors()
	utils.ConvertSQLToList(&sqlstructs.ProaesSQLRepo, professorRepo)
	return err
}

func (s SQLRepository) CreateClassRepo() error {
	err := createIfExistORReturnSQLRepo(
		"ClassSQLRepo",
		"classes",
		&sqlstructs.ClassSQLRepo,
		"id",
		"location",
		"professor",
		"students",
	)
	// Inplace modification, alter the fields in Repo
	classesRepo, err := s.repo.GetClasses()
	utils.ConvertSQLToList(&sqlstructs.ClassSQLRepo, classesRepo)
	return err
}

func (s SQLRepository) CreateStudentRepo() error {
	err := createIfExistORReturnSQLRepo(
		"StudentSQLRepo",
		"student",
		&sqlstructs.StudentSQLRepo,
		"id",
		"name",
		"transcript",
	)
	// Inplace modification, alter the fields in Repo
	studentRepo, err := s.repo.GetStudents()
	utils.ConvertSQLToList(&sqlstructs.StudentSQLRepo, studentRepo)
	return err
}
