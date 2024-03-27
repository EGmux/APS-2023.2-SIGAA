package control

import (
	sql "main/SQL"
	"main/utils/class"
	"main/utils/student"
)

func GetUser(username string) []student.Student{
	return sql.GetUser(username)
}


func GetAllClasses() []class.Class{
	return sql.GetAllClasses()
}

func GetSelectedClasses(selectedClasses []string) []class.Class{
	return sql.GetSelectedClasses(selectedClasses)
}

func Update_Student_Enrollment(student string, classes []class.Class){
	sql.Update_Student_Enrollment(student, classes)
}

func Update_Student_Deferral(username string){
	sql.Update_Student_Deferral(username)
}

func Retake_Student_Enrollment(username string){
	sql.Retake_Student_Enrollment(username)
}