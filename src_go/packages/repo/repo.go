package repo

import (
	"sigaa.ufpe/packages/repo/structs"
)

type IRepo interface {
	GetCredentials() (*[]*structs.Credentials, error)
	GetClasses() (*[]*structs.Class, error)
	GetCourses() (*[]*structs.Course, error)
	GetEnrollments() (*[]*structs.Enrollment, error)
	GetPROAES() (*[]*structs.PROAES, error)
	GetProfessors() (*[]*structs.Professor, error)
	GetTeachingScholarships() (*[]*structs.TeachingScholarship, error)
	GetStudents() (*[]*structs.Student, error)
	GetScholarships() (*[]*structs.Scholarship, error)
}

type Repo struct{}

// Return the Credentials in struct format
func (r *Repo) GetCredentials() (*[]*structs.Credentials, error) {
	return &structs.CredentialRepo, nil
}

func (r *Repo) GetClasses() (*[]*structs.Class, error) {
	return &structs.ClassRepo, nil
}

func (r *Repo) GetCourses() (*[]*structs.Course, error) {
	return &structs.CourseRepo, nil
}

func (r *Repo) GetEnrollments() (*[]*structs.Enrollment, error) {
	return &structs.EnrollmentRepo, nil
}

func (r *Repo) GetPROAES() (*[]*structs.PROAES, error) {
	return &structs.PROAESRepo, nil
}

func (r *Repo) GetProfessors() (*[]*structs.Professor, error) {
	return &structs.ProfessorRepo, nil
}

func (r *Repo) GetTeachingScholarships() (*[]*structs.TeachingScholarship, error) {
	return &structs.TeachingScholarshipRepo, nil
}

func (r *Repo) GetStudents() (*[]*structs.Student, error) {
	return &structs.StudentRepo, nil
}

func (r *Repo) GetScholarships() (*[]*structs.Scholarship, error) {
	return &structs.ScholarshipRepo, nil
}
