package abstractfactory

import (
	"os"

	"sigaa.ufpe/packages/data/repo"
	"sigaa.ufpe/packages/data/repo/sqlrepo"
)

// Aliases

type IRepoFactory interface {
	ReturnRepos() (repo.IRepo, error)
	CreateClassRepo() error
	CreateEnrollmentRepo() error
	CreateCredentialRepo() error
	CreatePROAESRepo() error
	CreateProfessorRepo() error
	CreateStudentRepo() error
}

// Return Interface to
func GetRepos() (IRepoFactory, error) {
	tech := os.Getenv("DBTECH")
	switch tech {
	case "SQL":
		return &sqlrepo.SQLRepository{}, nil
		// case "JSON":
		// 	return &JSONRepository{}, nil
	}
	return nil, nil
}
