package data

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
	Capacity    int
	Mandatory   bool
	Timetable   string
	Assessments []string
	Professor   []*Professor
	Grades      []*Grade
	Id          float64
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
