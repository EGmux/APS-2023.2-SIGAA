package main

import (
	login_controller "sigaa.ufpe/packages/comunication"
	data_studentsdata_studentsSQL "sigaa.ufpe/packages/data/students_data/students_SQL"
)

var db = make(map[string]string)

func main() {
	println("test 🤛")

	data_studentsdata_studentsSQL.InitDB()
	login_controller.Set_Login_Controller()
	// Listen and Server in 0.0.0.0:8080
}
