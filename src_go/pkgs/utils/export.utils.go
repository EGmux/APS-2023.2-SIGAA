package utils

const (
	SELECT_ALL DBQUERY = iota
)

type DBQUERY int64

// for multithreading
// type RepoConn RepositoryConn

var DB RepositoryConn
