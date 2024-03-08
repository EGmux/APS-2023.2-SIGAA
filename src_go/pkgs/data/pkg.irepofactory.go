package data

import (
	"os"
)

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

// Return Interface to
func GetRepos() (IRepoFactory, error) {
	tech := os.Getenv("DBTECH")
	switch tech {
	case "SQL":
		return &SQLRepository{}, nil
		// case "JSON":
		// 	return &JSONRepository{}, nil
	}
	return nil, nil
}
