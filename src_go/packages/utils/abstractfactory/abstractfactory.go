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
