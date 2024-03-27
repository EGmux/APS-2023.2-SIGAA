package facade

import (
	"main/control"
	"main/utils/class"
	"main/utils/student"
)

func GetUser(username string) []student.Student{
	return control.GetUser(username)
}


func GetAllClasses() []class.Class{
	return control.GetAllClasses()
}

func GetSelectedClasses(selectedClasses []string) []class.Class{
	return control.GetSelectedClasses(selectedClasses)
}

func Update_Student_Enrollment(student string, classes []class.Class){
	control.Update_Student_Enrollment(student, classes)
}

func Update_Student_Deferral(username string){
	control.Update_Student_Deferral(username)
}

func Retake_Student_Enrollment(username string){
	control.Retake_Student_Enrollment(username)
}