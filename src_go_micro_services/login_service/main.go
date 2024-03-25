package main

import (
	//data_studentsdata_studentsSQL "login_service/SQL"
	logincontroller "login_service/login_controller"
)

func main() {

	canal := make(chan bool)
	//data_studentsdata_studentsSQL.InitDB()
	logincontroller.Set_Login_Controller()


	<-canal
}
