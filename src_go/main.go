package main

import (
	_ "github.com/joho/godotenv/autoload"

	//enrollmentcontroller "sigaa.ufpe/packages/comunication/enrollment_controller"
	enrollmentcontroller "sigaa.ufpe/packages/comunication/enrollment_controller"
	"sigaa.ufpe/packages/comunication/login_controller"
	"sigaa.ufpe/packages/comunication/main_menu_controller"
	schollarshipcontroller "sigaa.ufpe/packages/comunication/schollarship_controller"
	teachingscholarshipcontroller "sigaa.ufpe/packages/comunication/schollarship_controller/teachingScholarship"
	singleton_db "sigaa.ufpe/packages/utils/singleton"
)

var db = make(map[string]string)

func main() {
	channel := make(chan bool)
	println("test 🤛")

	singleton_db.InitDB()
	go teachingscholarshipcontroller.Set_Teaching_Scholarship_Controller()
	go schollarshipcontroller.Set_Schollarship_Controller()
	go main_menu_controller.Set_Main_Menu_Controller()
	go login_controller.Set_Login_Controller() //8080
	go enrollmentcontroller.Set_Enrollment_Controller() //8084
	// Listen and Server in 0.0.0.0:8080
	//var student []student.Student = facade.GetUser("hugo")
	//fmt.Println(student)
	<-channel
}
