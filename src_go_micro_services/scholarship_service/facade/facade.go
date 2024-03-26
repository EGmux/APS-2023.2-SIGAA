package facade

import (
	scholarshipcontrol "scholarship/scholarship_control"
	"scholarship/student"
)

func GetUser(username string) []student.Student{
	return scholarshipcontrol.GetUser(username)
}
