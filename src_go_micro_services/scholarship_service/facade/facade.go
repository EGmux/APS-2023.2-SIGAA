package facade

import (
	scholarshipcontrol "scholarship/scholarship_control"
	"scholarship/utils/student"
	"scholarship/utils/teachingscholarship"
)

func GetUser(username string) []student.Student{
	return scholarshipcontrol.GetUser(username)
}

func GetAllTeachingScholarships() []teachingscholarship.TeachingScholarship{
	return scholarshipcontrol.GetAllTeachingScholarships()
}

func GetAvailableTeachingScholarships() []teachingscholarship.TeachingScholarship {
	return scholarshipcontrol.GetAvailableTeachingScholarships()
}

func Apply_Student_To_Teaching_Scholarship(User string, Scholarship_Id string) {
	scholarshipcontrol.Apply_Student_To_Teaching_Scholarship(User, Scholarship_Id)
}