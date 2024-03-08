package data

import (
	"errors"
)

type RepoType int

// ENUM
const (
	CREDENTIALS RepoType = iota
	PROFESSOR
	EPROAES
	ENROLLMENT
	CLASS
	COURSE
	STUDENT
	NOREPO
)

// Type Constraints
type TablePtrs interface {
	Credentials | Professor | PROAES | Enrollment | Class | Student | TeachingScholarship | Scholarship | Course | Grade
}

// Default structs
type Class struct {
	Capacity  int
	Mandatory bool
	Timetable string
	Assesment string
	Professor []*Professor
	Grades    []*Grade
	Id        float64
	// Never copy this struct
	StudentClass_   StudentClass
	RecursiveClass_ RecursiveClass
}

type RecursiveClass struct {
	Equivalency []*Class
}

type Course struct {
	Name       string
	Department string
	Curriculum string
	Classes    []*Class
	Id         float64
}

type Enrollment struct {
	Id         float64
	Name       string
	Status     bool
	CPF        string
	Transcript []*Class
}

type Grade struct {
	Grade string
	Id    float64
}

type PROAES struct {
	Email string
	Id    float64
}

type Professor struct {
	Name       string
	Email      string
	CPF        string
	Department string
	Id         float64
}

type Scholarship struct {
	Value   int
	Student Student
	Id      float64
}

type TeachingScholarship struct {
	Scholarship Scholarship
	Class       Class
	Semester    string
	Professor   Professor
	Id          float64
}

type Student struct {
	Name          string
	CPF           string
	Id            float64
	StudentClass_ StudentClass
}

type StudentClass struct {
	Transcript []*Class
	Student    []*Student
}

type Credentials struct {
	User     string
	Password string
	Logged   bool
	Id       float64
}

// TODO: convert to singleton pattern
var (
	ClassRepo               []*Class
	CourseRepo              []*Course
	EnrollmentRepo          []*Enrollment
	PROAESRepo              []*PROAES
	ProfessorRepo           []*Professor
	TeachingScholarshipRepo []*TeachingScholarship
	StudentRepo             []*Student
	CredentialRepo          []*Credentials
	ScholarshipRepo         []*Scholarship

	// Default constructors
)

// Set default repo for constructor so nil pointers can be used
func SetDefaultRepo() {
	ClassRepo = make([]*Class, 1)
	CredentialRepo = make([]*Credentials, 1)
	StudentRepo = make([]*Student, 1)
	TeachingScholarshipRepo = make([]*TeachingScholarship, 1)
	ProfessorRepo = make([]*Professor, 1)
	CourseRepo = make([]*Course, 1)
	EnrollmentRepo = make([]*Enrollment, 1)
	ScholarshipRepo = make([]*Scholarship, 1)
	TeachingScholarshipRepo = make([]*TeachingScholarship, 1)
}

func credentialsC() *Credentials {
	return &Credentials{
		Id:       0,
		Password: "",
		User:     "",
		Logged:   false,
	}
}

func proaesC() *PROAES {
	return &PROAES{
		Id:    0,
		Email: "",
	}
}

func classC() *Class {
	return &Class{
		Id:              0,
		Capacity:        0,
		Mandatory:       false,
		Timetable:       "",
		Assesment:       "",
		Professor:       []*Professor{professorC()},
		Grades:          []*Grade{gradesC()},
		StudentClass_:   StudentClass{},
		RecursiveClass_: RecursiveClass{},
	}
}

func gradesC() *Grade {
	return &Grade{
		Grade: "",
		Id:    0,
	}
}

func professorC() *Professor {
	return &Professor{
		Id:         0,
		Name:       "",
		Email:      "",
		CPF:        "",
		Department: "",
	}
}

func scholarshipC() *Scholarship {
	return &Scholarship{
		Value:   0,
		Student: *studentC(),
		Id:      0,
	}
}

func studentC() *Student {
	return &Student{
		Id:            0,
		Name:          "",
		CPF:           "",
		StudentClass_: StudentClass{},
	}
}

func teachSholarshipC() *TeachingScholarship {
	return &TeachingScholarship{
		Scholarship: Scholarship{},
		Class:       *classC(),
		Semester:    "",
		Professor:   *professorC(),
		Id:          0,
	}
}

func courseC() *Course {
	return &Course{
		Name:       "",
		Department: "",
		Curriculum: "",
		Classes:    []*Class{classC()},
		Id:         0,
	}
}

func enrollmentC() *Enrollment {
	return &Enrollment{
		Id:         0,
		Name:       "",
		Status:     false,
		CPF:        "",
		Transcript: []*Class{classC()},
	}
}

// generic function to add rows to table
func addRows[T TablePtrs](s *[]*T, f func() *T) error {
	defaultTable := *f()
	*s = append(*s, &defaultTable)
	return nil
}

// Default consructors
func AddRows[T TablePtrs](s *[]*T) error {
	if s == nil {
		SetDefaultRepo()
	}
	switch t := any(s).(type) {
	case *[]*Class:
		addRows(t, classC)
	case *[]*Student:
		addRows(t, studentC)
	case *[]*Credentials:
		addRows(t, credentialsC)
	case *[]*Professor:
		addRows(t, professorC)
	case *[]*Course:
		addRows(t, courseC)
	case *[]*PROAES:
		addRows(t, proaesC)
	case *[]*Enrollment:
		addRows(t, enrollmentC)
	case *[]*TeachingScholarship:
		addRows(t, teachSholarshipC)
	}
	return errors.New("Not a valid table")
}
