package control

import (
	data_studentsdata_studentsSQL "login_service/SQL"
)

func IsValidUser(username string, password string) bool {
	return data_studentsdata_studentsSQL.IsValidUser(username, password)
}

func InsertUser(username string, password string){
	data_studentsdata_studentsSQL.InsertUser(username, password)
}

func GetAllUsers() ([]data_studentsdata_studentsSQL.Student){
	return data_studentsdata_studentsSQL.GetAllUsers()
}

func GetUser(username string) []data_studentsdata_studentsSQL.Student{
	return data_studentsdata_studentsSQL.GetUser(username)
}