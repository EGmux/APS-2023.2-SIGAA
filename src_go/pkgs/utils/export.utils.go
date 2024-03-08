package utils

const (
	SELECT_ALL DBQUERY = iota
)

type DBQUERY int64

var DB RepositoryConn
