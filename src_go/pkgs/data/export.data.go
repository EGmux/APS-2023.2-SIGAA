package data

import "os"

var IReposityFactory IRepoFactory

// Return Interface to
func GetReposFactory() (IRepoFactory, error) {
	tech := os.Getenv("DBTECH")
	switch tech {
	case "SQL":
		return &SQLRepository{}, nil
		// case "JSON":
		// 	return &JSONRepository{}, nil
	}
	return nil, nil
}
