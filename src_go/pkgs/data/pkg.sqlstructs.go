package data

import (
	"errors"
	"math/rand"
)

// Type Constraints

type SQLTablesPtrs interface {
	CredentialsSQL | ProfessorSQL | PROAESSQL | EnrollmentSQL | ClassSQL | StudentSQL | CourseSQL | ScholarshipSQL | TeachingScholarshipSQL | GradeSQL
}

// To be mapped structs

type CourseSQL struct {
	Name       string      `db:"name"`
	Department string      `db:"department"`
	Curriculum string      `db:"curriculum"`
	Classes    []*ClassSQL `db:"classes"`
	Id         float64     `db:"idcourse"`
}

type ScholarshipSQL struct {
	Value   int        `db:"value"`
	Student StudentSQL `db:"student"`
	Id      float64    `db:"idscholarship"`
}

type TeachingScholarshipSQL struct {
	Scholarship ScholarshipSQL `db:"scholarship"`
	Class       ClassSQL       `db:"class"`
	Semester    string         `db:"semester"`
	Professor   ProfessorSQL   `db:"professor"`
	Id          float64        `db:"idteachingscholarship"`
}

type CredentialsSQL struct {
	Id       float64 `db:"idcredential"`
	Passwd   string  `db:"passwd"`
	Username string  `db:"username"`
	Logged   bool    `db:"logged"`
}

type GradeSQL struct {
	Grade string  `db:"grade"`
	Id    float64 `db:"idgrade"`
}

type PROAESSQL struct {
	Id    float64 `db:"idproaes"`
	Email string  `db:"email"`
}

type StudentSQL struct {
	Id               float64         `db:"idstudent"`
	Name             string          `db:"name"`
	CPF              string          `db:"cpf"`
	StudentClassSQL_ StudentClassSQL `db:"studentclass"`
}

type EnrollmentSQL struct {
	Id         float64     `db:"idenrollment"`
	Name       string      `db:"name"`
	Status     bool        `db:"status"`
	CPF        string      `db:"cpf"`
	Transcript []*ClassSQL `db:"transcript"`
}

type ProfessorSQL struct {
	Id         float64 `db:"idprofessor"`
	Name       string  `db:"name"`
	Email      string  `db:"email"`
	CPF        string  `db:"cpf"`
	Department string  `db:"department"`
}

type StudentClassSQL struct {
	Transcript []*ClassSQL   `db:"transcript"`
	Students   []*StudentSQL `db:"students"`
}

type ClassSQL struct {
	Id                 float64           `db:"idclass"`
	Capacity           int               `db:"capacity"`
	Required           bool              `db:"required"`
	Timetable          string            `db:"timetable"`
	Assesment          string            `db:"assesment"`
	Professor          []*ProfessorSQL   `db:"professor"`
	Grades             []*GradeSQL       `db:"grades"`
	StudentClassSQL_   StudentClassSQL   `db:"studentclass"`
	RecursiveClassSQL_ RecursiveClassSQL `db:"recursiveclass"`
}

type RecursiveClassSQL struct {
	Equivalency []*ClassSQL
}

// TODO: convert to singleton pattern
var (
	CredentialSQLRepo          []*CredentialsSQL
	ProaesSQLRepo              []*PROAESSQL
	StudentSQLRepo             []*StudentSQL
	EnrollmentSQLRepo          []*EnrollmentSQL
	ProfessorSQLRepo           []*ProfessorSQL
	ClassSQLRepo               []*ClassSQL
	CourseSQLRepo              []*CourseSQL
	ScholarshipSQLRepo         []*ScholarshipSQL
	TeachingScholarshipSQLRepo []*TeachingScholarshipSQL

	// Default constructors
)

func SetDefaultSQLRepos() {
	ClassSQLRepo = make([]*ClassSQL, 1)
	CredentialSQLRepo = make([]*CredentialsSQL, 1)
	StudentSQLRepo = make([]*StudentSQL, 1)
	TeachingScholarshipSQLRepo = make([]*TeachingScholarshipSQL, 1)
	ProfessorSQLRepo = make([]*ProfessorSQL, 1)
	CourseSQLRepo = make([]*CourseSQL, 1)
	EnrollmentSQLRepo = make([]*EnrollmentSQL, 1)
	ScholarshipSQLRepo = make([]*ScholarshipSQL, 1)
}

func credentialsSQLC() *CredentialsSQL {
	passwds := []string{"helloword", "hacked", "bugged", "DEAD", "good", "emacs"}
	usernames := []string{
		"Caros",
		"Monica",
		"Selminha",
		"Juliano",
		"Tereza",
		"Marcaos",
		"Marcelo",
		"Tobias",
	}
	Logged := []bool{true, false}
	return &CredentialsSQL{
		Id:       0,
		Passwd:   passwds[rand.Intn(len(passwds))],
		Username: usernames[rand.Intn(len(usernames))],
		Logged:   Logged[rand.Intn(len(Logged))],
	}
}

func proaesSQLC() *PROAESSQL {
	emails := []string{
		"bolinha@gmail.com",
		"cinufpe@gmail.com",
		"morning@cin.ufpe.br",
		"textmaxe@outlook.com",
	}
	return &PROAESSQL{
		Id:    0,
		Email: emails[rand.Intn(len(emails))],
	}
}

func classSQLC() *ClassSQL {
	capacities := []int{1, 0, 10, 20, 60, 45, 30}
	required := []bool{true, false}
	timetables := []string{"tue-08-10", "wed-10-12", "fri-08-10", "mon-15-17", "sun-17-19"}
	assessments := []string{"13-feb", "20-mar", "10-apr", "12-jan", "25-dec", "9-jun"}
	return &ClassSQL{
		Id:                 0,
		Capacity:           capacities[rand.Intn(len(capacities))],
		Required:           required[rand.Intn(len(required))],
		Timetable:          timetables[rand.Intn(len(timetables))],
		Assesment:          assessments[rand.Intn(len(assessments))],
		Professor:          []*ProfessorSQL{professorSQLC()},
		Grades:             []*GradeSQL{gradesSQLC()},
		StudentClassSQL_:   StudentClassSQL{},
		RecursiveClassSQL_: RecursiveClassSQL{},
	}
}

func gradesSQLC() *GradeSQL {
	grades := []string{"10.0", "A", "B", "0.8", "0,9", "0", "0.0"}
	return &GradeSQL{
		Grade: grades[rand.Intn(len(grades))],
		Id:    0,
	}
}

func professorSQLC() *ProfessorSQL {
	names := []string{
		"Janailson Santana",
		"Tiago leite",
		"Tiamoe e Pumba",
		"Boris Cafa",
		"tulio Margors",
	}
	emails := []string{"@cin.ufpe.br", "@gmail.com", "@outlookc.mo", "@icloud.com"}
	cpfs := []string{"12891345789", "90128945892", "90124782901", "991019202902", "982389208292"}
	departments := []string{"CTG", "CIN", "CAC", "DEE", "DMAT", "CCEN", "DQF"}
	return &ProfessorSQL{
		Id:         0,
		Name:       names[rand.Intn(len(names))],
		Email:      emails[rand.Intn(len(emails))],
		CPF:        cpfs[rand.Intn(len(cpfs))],
		Department: departments[rand.Intn(len(departments))],
	}
}

func scholarshipSQLC() *ScholarshipSQL {
	values := []int{100.00, 500.00, 2000.00, 1000.00, 4000}
	return &ScholarshipSQL{
		Value:   values[rand.Intn(len(values))],
		Student: *studentSQLC(),
		Id:      0,
	}
}

func studentSQLC() *StudentSQL {
	names := []string{
		"Caio santana",
		"Vitoria Baters",
		"Hugo Olavo",
		"Mariana Sentinela",
		"Bers Mornen",
		"Julio Tuliano",
	}
	cpfs := []string{"90238945892", "90128923784", "78338923478", "90124891234", "89109127823"}
	return &StudentSQL{
		Id:               0,
		Name:             names[rand.Intn(len(names))],
		CPF:              cpfs[rand.Intn(len(cpfs))],
		StudentClassSQL_: StudentClassSQL{},
	}
}

func courseSQLC() *CourseSQL {
	names := []string{
		"computer engineering",
		"material sciences",
		"chemistry",
		"Veterinary",
		"computer science",
		"biology",
	}
	department := []string{"cin", "ctg", "dmat", "dqf", "dee"}
	curricuilum := []string{
		"bunch of courses",
		"another bunch of courses",
		"more bounch of coursers",
	}
	return &CourseSQL{
		Name:       names[rand.Intn(len(names))],
		Department: department[rand.Intn(len(department))],
		Curriculum: curricuilum[rand.Intn(len(curricuilum))],
		Classes:    []*ClassSQL{classSQLC()},
		Id:         0,
	}
}

func enrollmentSQLC() *EnrollmentSQL {
	names := []string{"Barbara mikaecho", "Tulio olavo", "Diabrao e Mar", "fernando dias"}
	status := []bool{true, false}
	cpf := []string{"091289234218", "90129012342", "89127812341", "90124892342"}
	return &EnrollmentSQL{
		Id:         0,
		Name:       names[rand.Intn(len(names))],
		Status:     status[rand.Intn(len(status))],
		CPF:        cpf[rand.Intn(len(cpf))],
		Transcript: []*ClassSQL{classSQLC()},
	}
}

func teachingScholarshipSQLC() *TeachingScholarshipSQL {
	semesters := []string{"1", "2"}
	return &TeachingScholarshipSQL{
		Scholarship: *scholarshipSQLC(),
		Class:       *classSQLC(),
		Semester:    semesters[rand.Intn(len(semesters))],
		Professor:   *professorSQLC(),
		Id:          0,
	}
}

// generic function to add rows to table
func addRowsSQL[T SQLTablesPtrs](s *[]*T, f func() *T) error {
	defaultTable := *f()
	*s = append(*s, &defaultTable)
	return nil
}

// Default consructors
func AddRowsSQ[T SQLTablesPtrs](s *[]*T) error {
	if s == nil {
		SetDefaultSQLRepos()
	}
	switch t := any(s).(type) {
	case *[]*ClassSQL:
		addRowsSQL(t, classSQLC)
	case *[]*CredentialsSQL:
		addRowsSQL(t, credentialsSQLC)
	case *[]*ProfessorSQL:
		addRowsSQL(t, professorSQLC)
	case *[]*CourseSQL:
		addRowsSQL(t, courseSQLC)
	case *[]*PROAESSQL:
		addRowsSQL(t, proaesSQLC)
	case *[]*EnrollmentSQL:
		addRowsSQL(t, enrollmentSQLC)
	case *[]*TeachingScholarshipSQL:
		addRowsSQL(t, teachingScholarshipSQLC)
	case *[]*StudentSQL:
		addRowsSQL(t, studentSQLC)
	}
	return errors.New("Not a valid table")
}
