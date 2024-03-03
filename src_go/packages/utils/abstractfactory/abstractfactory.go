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
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> ce465bc (update(packages:abstractfactory): extract struct types to types.go)

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
<<<<<<< HEAD
=======
>>>>>>> 95fd195 (update(main.go): testcode)
=======
>>>>>>> fe4aa66 (update(packages:abstractfactory): extract struct types to types.go)
>>>>>>> ce465bc (update(packages:abstractfactory): extract struct types to types.go)
