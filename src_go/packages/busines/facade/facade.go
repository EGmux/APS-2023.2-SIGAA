package facade

import (
	login_control "sigaa.ufpe/packages/busines/login_control"
	scholarshipcontrol "sigaa.ufpe/packages/busines/scholarship_control"
	teachingscholarshipcontrol "sigaa.ufpe/packages/busines/scholarship_control/teaching_scholarship_control"
	teachingscholarship "sigaa.ufpe/packages/data/scholarship_data/teachingScholarship"
	student "sigaa.ufpe/packages/data/students_data"
)

func IsValidUser(username string, password string) bool {
	return login_control.IsValidUser(username, password)
}

func InitDB() {
	login_control.Init_StudentsDB()
}

func CloseDB() {
	login_control.CloseDB()
}

func InsertUser(username string, password string) {
	login_control.InsertUser(username, password)
}

func GetAllUsers() []student.Student {
	return login_control.GetAllUsers()
}

func GetAllTeachingScholarships() []teachingscholarship.TeachingScholarship {
	return scholarshipcontrol.GetAllTeachingScholarships()
}

func GetAvailableTeachingScholarships() []teachingscholarship.TeachingScholarship {
	return scholarshipcontrol.GetAvailableTeachingScholarships()
}

func Apply_Student_To_Teaching_Scholarship(User string, Scholarship_Id string) {
	teachingscholarshipcontrol.Apply_Student_To_Teaching_Scholarship(User, Scholarship_Id)
}
