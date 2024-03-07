package scholarship

import student "sigaa.ufpe/packages/data/students_data"


type Scholarship struct{
	Id int
	Value int
	Student student.Student
}

