package facade

import (
	deferralcontrol "sigaa.ufpe/packages/busines/deferral_control"
	enrollmentcontrol "sigaa.ufpe/packages/busines/enrollment_control"
	login_control "sigaa.ufpe/packages/busines/login_control"
	scholarshipcontrol "sigaa.ufpe/packages/busines/scholarship_control"
	teachingscholarshipcontrol "sigaa.ufpe/packages/busines/scholarship_control/teaching_scholarship_control"
	class "sigaa.ufpe/packages/data/class_data"
	teachingscholarship "sigaa.ufpe/packages/data/scholarship_data/teachingScholarship"
	student "sigaa.ufpe/packages/data/students_data"
)

func IsValidUser(username string, password string) bool {
	return login_control.IsValidUser(username, password)
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

func GetUser(username string) []student.Student{
	return login_control.GetUser(username)
}

func GetAllClasses() []class.Class{
	return enrollmentcontrol.GetAllClasses()
}

func GetSelectedClasses(selectedClasses []string) []class.Class{
	return enrollmentcontrol.GetSelectedClasses(selectedClasses)
}

func Update_Student_Enrollment(student string, classes []class.Class){
	enrollmentcontrol.Update_Student_Enrollment(student, classes)
}

func Update_Student_Deferral(username string){
	deferralcontrol.Update_Student_Deferral(username)
}