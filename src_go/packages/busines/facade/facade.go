package facade

import (
	login_control "sigaa.ufpe/packages/busines/login_control"
	student "sigaa.ufpe/packages/data/students_data"
)





func IsValidUser(username string, password string) bool{
	return login_control.IsValidUser(username, password)
}

func InitDB(){
	login_control.InitDB()
}

func CloseDB(){
	login_control.CloseDB()
}

func InsertUser(username string, password string){
	login_control.InsertUser(username, password)
}

func GetAllUsers() ([]student.Student){
	return login_control.GetAllUsers()
}

