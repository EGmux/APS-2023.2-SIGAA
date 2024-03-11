package enrollmentcontrol

import (
	class "sigaa.ufpe/packages/data/class_data"
	data_classdata_classsql "sigaa.ufpe/packages/data/class_data/class_SQL"
	data_studentsdata_studentsSQL "sigaa.ufpe/packages/data/students_data/students_SQL"
)

func GetAllClasses() []class.Class{
	return data_classdata_classsql.GetAllClasses()
}

func GetSelectedClasses(selectedClasses []string) []class.Class{
	return data_classdata_classsql.GetSelectedClasses(selectedClasses)
}

func Update_Student_Enrollment(student string, classes []class.Class){
	data_studentsdata_studentsSQL.Update_Student_Enrollment(student, classes)
}