package data

type IRepo interface {
	GetCredentials() (*[]*Credentials, error)
	GetClasses() (*[]*Class, error)
	GetCourses() (*[]*Course, error)
	GetEnrollments() (*[]*Enrollment, error)
	GetPROAES() (*[]*PROAES, error)
	GetProfessors() (*[]*Professor, error)
	GetTeachingScholarships() (*[]*TeachingScholarship, error)
	GetStudents() (*[]*Student, error)
	GetScholarships() (*[]*Scholarship, error)
}

type Repo struct{}

// Return the Credentials in struct format
func (r *Repo) GetCredentials() (*[]*Credentials, error) {
	return &CredentialRepo, nil
}

func (r *Repo) GetClasses() (*[]*Class, error) {
	return &ClassRepo, nil
}

func (r *Repo) GetCourses() (*[]*Course, error) {
	return &CourseRepo, nil
}

func (r *Repo) GetEnrollments() (*[]*Enrollment, error) {
	return &EnrollmentRepo, nil
}

func (r *Repo) GetPROAES() (*[]*PROAES, error) {
	return &PROAESRepo, nil
}

func (r *Repo) GetProfessors() (*[]*Professor, error) {
	return &ProfessorRepo, nil
}

func (r *Repo) GetTeachingScholarships() (*[]*TeachingScholarship, error) {
	return &TeachingScholarshipRepo, nil
}

func (r *Repo) GetStudents() (*[]*Student, error) {
	return &StudentRepo, nil
}

func (r *Repo) GetScholarships() (*[]*Scholarship, error) {
	return &ScholarshipRepo, nil
}
