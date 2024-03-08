package main

import (
	_ "github.com/joho/godotenv/autoload"
	// teachingscholarshipsql "sigaa.ufpe/pkgs/data/scholarship_data/teachingScholarship/teachingScholarship_SQL"
	// data_studentsdata_studentsSQL "sigaa.ufpe/pkgs/data/students_data/students_SQL"
)

var db = make(map[string]string)

func main() {
	channel := make(chan bool)
	println("test 🤛")

	// teachingscholarshipsql.Init_TeachingScholarship_DB()
	// data_studentsdata_studentsSQL.Init_Students_DB()
	// go teachingscholarshipcontroller.Set_Teaching_Scholarship_Controller()
	// go schollarshipcontroller.Set_Schollarship_Controller()
	// go main_menu_controller.Set_Main_Menu_Controller()
	// go login_controller.Set_Login_Controller()
	// Listen and Server in 0.0.0.0:8080
	<-channel
}
