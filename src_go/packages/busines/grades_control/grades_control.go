package gradescontrol

import (
	grades "sigaa.ufpe/packages/data/grades_data"
	gradessql "sigaa.ufpe/packages/data/grades_data/grades_SQL"
)

func Get_All_Grades_From_Student(username string) []grades.Grade{
	return gradessql.Get_All_Grades_From_Student(username)
}