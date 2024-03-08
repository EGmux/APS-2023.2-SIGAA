package data

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Control test environment
type debug struct {
	droptable  bool
	showerrors bool
}

// scope global vars
var (
	credentialsRepo   []*Credentials
	credentialSQLRepo []*CredentialsSQL

	proaesRepo     []*PROAES
	proaesSQLRRepo []*PROAESSQL

	professorRepo    []*Professor
	professorSQLRepo []*ProfessorSQL

	classRepo    []*Class
	classSQLRepo []*ClassSQL

	gradeRepo    []*Grade
	gradeSQLRepo []*GradeSQL

	studentRepo    []*Student
	studentSQLRepo []*StudentSQL

	utils utilsTestSuite
)

// Acess options and test suites
type utilsTestSuite struct {
	opts debug
	iter int
}

// Intialize tables for each new test, each call is a single row
func (u utilsTestSuite) TestSetup() error {
	for i := 0; i < u.iter; i++ {
		// Struct's
		AddRows(&credentialsRepo)
		AddRows(&proaesRepo)
		AddRows(&classRepo)
		AddRows(&gradeRepo)
		AddRows(&studentRepo)
		// SQLRepo's
		AddRowsSQ(&credentialSQLRepo)
		AddRowsSQ(&proaesSQLRRepo)
		AddRowsSQ(&classSQLRepo)
		AddRowsSQ(&gradeSQLRepo)
		AddRowsSQ(&studentSQLRepo)
	}
	return nil
}

// func (suite *UtilsTestSuite) SetupSuite() {
// 	suite.opts.showerrors = true
// }

// Reset state for each new test
func (u *utilsTestSuite) TearDownTest() {
	// SQLRepo's
	credentialSQLRepo = nil
	proaesSQLRRepo = nil
	classSQLRepo = nil
	gradeSQLRepo = nil
	studentSQLRepo = nil
	gradeSQLRepo = nil
	// Struct's
	credentialsRepo = nil
	proaesRepo = nil
	professorRepo = nil
	classRepo = nil
	gradeRepo = nil
	studentRepo = nil
}

// deepcheck the Credentials struct
func deepCheckCredentials(
	pos int,
	u []*CredentialsSQL,
	v []*Credentials,
) (bool, error) {
	if !cmp.Equal(v[pos].Id, u[pos].Id) ||
		!cmp.Equal(v[pos].Logged, u[pos].Logged) ||
		!cmp.Equal(v[pos].Password, u[pos].Passwd) ||
		!cmp.Equal(v[pos].User, u[pos].Username) {
		// TODO: need to improve the error messages
		return false, errors.New("Credentials bug")
	}
	return true, nil
}

// deepcheck the Professsor struct
func deepCheckProfessor(
	pos int,
	u []*ProfessorSQL,
	v []*Professor,
) (bool, error) {
	if !cmp.Equal(u[pos].Id, v[pos].Name) ||
		!cmp.Equal(u[pos].Email, v[pos].Email) ||
		!cmp.Equal(u[pos].CPF, v[pos].CPF) {
		return false, errors.New("Professor bug")
	}
	return true, nil
}

// // deepcheck the Grades struct
func deepCheckGrades(
	pos int,
	u []*GradeSQL,
	v []*Grade,
) (bool, error) {
	if cmp.Equal(u[pos].Grade, v[pos].Grade) ||
		cmp.Equal(u[pos].Id, v[pos].Id) {
		return false, errors.New("Grades bug")
	}
	return true, nil
}

// deepcheck the Class struct
func deepCheckClass(
	pos int,
	u []*ClassSQL,
	v []*Class,
) (bool, error) {
	rangeOver := func() bool {
		for p := range u[pos].Professor {
			t, _ := deepCheckProfessor(p, u[pos].Professor, v[pos].Professor)
			v, _ := deepCheckGrades(p, u[pos].Grades, v[pos].Grades)
			if !t || !v {
				break
			}
			return true
		}
		return false
	}
	if !cmp.Equal(u[pos].Id, v[pos].Id) ||
		!cmp.Equal(u[pos].Required, v[pos].Mandatory) ||
		!cmp.Equal(u[pos].Timetable, v[pos].Timetable) ||
		!cmp.Equal(u[pos].Assesment, v[pos].Assesment) ||
		!cmp.Equal(u[pos].Capacity, v[pos].Capacity) || rangeOver() {
		return false, errors.New("Class Bug")
	}
	return true, nil
}

// deepcheck the PROAEST struct
func deepChecPROAES(
	pos int,
	u []*PROAESSQL,
	v []*PROAES,
) (bool, error) {
	if !cmp.Equal(u[pos].Id, v[pos].Id) ||
		!cmp.Equal(u[pos].Email, v[pos].Email) {
		return false, errors.New("PROAES bug")
	}
	return true, nil
}

// deepcheck the Student struct
func deepCheckStudent(
	pos int,
	u []*StudentSQL,
	v []*Student,
) (bool, error) {
	if !cmp.Equal(u[pos].CPF, v[pos].CPF) ||
		!cmp.Equal(u[pos].Id, v[pos].Id) ||
		!cmp.Equal(u[pos].Name, v[pos].Name) {
		return false, errors.New("Student bug")
	}
	return true, nil
}

func deepCheck[T SQLTablesPtrs, V TablePtrs](
	te *testing.T,
	t []*T,
	v []*V,
	dc func(int, []*T, []*V) (bool, error),
) (bool, error) {
	for pos := range t {
		_, err := dc(pos, t, v)
		if err != nil {
			te.Error(err)
		}
	}
	return true, nil
}

// Check that a SQL table map to a Struct table
func TestMapSQLToStruct(t *testing.T) {
	utils.iter = 0
	utils.TestSetup()
	ConvertSQLToList(&credentialSQLRepo, &credentialsRepo)
	ConvertSQLToList(&proaesSQLRRepo, &proaesRepo)
	ConvertSQLToList(&professorSQLRepo, &proaesRepo)
	ConvertSQLToList(&classSQLRepo, &classRepo)
	ConvertSQLToList(&studentSQLRepo, &studentRepo)
	ConvertSQLToList(&gradeSQLRepo, &gradeRepo)

	deepCheck(t, credentialSQLRepo, credentialsRepo, deepCheckCredentials)
	deepCheck(t, proaesSQLRRepo, proaesRepo, deepChecPROAES)
	deepCheck(t, professorSQLRepo, professorRepo, deepCheckProfessor)
	deepCheck(t, classSQLRepo, classRepo, deepCheckClass)
	deepCheck(t, studentSQLRepo, studentRepo, deepCheckStudent)
	deepCheck(t, gradeSQLRepo, gradeRepo, deepCheckGrades)
	utils.TearDownTest()
}

// Check that a Struct table map to SQL table
func TestMapStructToSQL(t *testing.T) {
	utils.iter = 0
	utils.TestSetup()
	ConvertStructToSQL(&credentialSQLRepo, &credentialsRepo)
	ConvertStructToSQL(&credentialSQLRepo, &credentialsRepo)
	ConvertStructToSQL(&proaesSQLRRepo, &proaesRepo)
	ConvertStructToSQL(&professorSQLRepo, &proaesRepo)
	ConvertStructToSQL(&classSQLRepo, &classRepo)
	ConvertStructToSQL(&studentSQLRepo, &studentRepo)
	ConvertStructToSQL(&gradeSQLRepo, &gradeRepo)

	deepCheck(t, credentialSQLRepo, credentialsRepo, deepCheckCredentials)
	deepCheck(t, proaesSQLRRepo, proaesRepo, deepChecPROAES)
	deepCheck(t, classSQLRepo, classRepo, deepCheckClass)
	deepCheck(t, studentSQLRepo, studentRepo, deepCheckStudent)
	deepCheck(t, gradeSQLRepo, gradeRepo, deepCheckGrades)
	utils.TearDownTest()
}
