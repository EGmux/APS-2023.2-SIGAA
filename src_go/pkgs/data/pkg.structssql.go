package data

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
	Assessments        []string          `db:"assessments"`
	Professor          []*ProfessorSQL   `db:"professors"`
	Grades             []*GradeSQL       `db:"grades"`
	StudentClassSQL_   StudentClassSQL   `db:"studentclass"`
	RecursiveClassSQL_ RecursiveClassSQL `db:"recursiveclass"`
}

type RecursiveClassSQL struct {
	Equivalency []*ClassSQL
}
