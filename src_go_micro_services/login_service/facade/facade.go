package facade

import (
	data_studentsdata_studentsSQL "login_service/SQL"
	"login_service/control"
)

func GetAllUsers() [] data_studentsdata_studentsSQL.Student{
	return control.GetAllUsers()
}

func GetUser(username string) []data_studentsdata_studentsSQL.Student{
	return control.GetUser(username)
}

func IsValidUser(username string, password string) bool {
	return control.IsValidUser(username, password)
}

func InsertUser(username string, password string){
	control.InsertUser(username, password)
}