package data

// Aliases

type IRepoFactory interface {
	ReturnRepos() (IRepo, error)
	CreateClassRepo() error
	CreateEnrollmentRepo() error
	CreateCredentialRepo() error
	CreatePROAESRepo() error
	CreateProfessorRepo() error
	CreateStudentRepo() error
}
