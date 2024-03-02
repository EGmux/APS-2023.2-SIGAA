package abstractfactory

type IAbstractFactory interface {
	GetFactory()
	CreateCourseRepo() bool
	CreateSubjectRepo() bool
	CreateClassRepo() bool
	CreateEnrollmentRepo() bool
	CreateCredentialRepo() bool
	CreateTranscriptRepo() bool
	CreatePROAESRepo() bool
	CreateProfessorRepo() bool
}
<<<<<<< HEAD
=======

type SQLRepoSitory struct {
	conn *pgxpool.Pool
}

// TODO: implement
// If table exits return else create it
func (repo *sqlrepo) CreateCourseRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateCredentialRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreatePROAESRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateTranscriptRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateEnrollmentRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateProfessorRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateClassRepo() bool {
	return true
}

// TODO: implement
func (repo *sqlrepo) CreateSubjectRepo() bool {
	return true
}
>>>>>>> 95fd195 (update(main.go): testcode)
