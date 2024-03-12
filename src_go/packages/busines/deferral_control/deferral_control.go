package deferralcontrol

import data_studentsdata_studentsSQL "sigaa.ufpe/packages/data/students_data/students_SQL"

func Update_Student_Deferral(username string){
	data_studentsdata_studentsSQL.Update_Student_Deferral(username)
}

func Retake_Student_Enrollment(username string){
	data_studentsdata_studentsSQL.Retake_Student_Enrollment(username)
}