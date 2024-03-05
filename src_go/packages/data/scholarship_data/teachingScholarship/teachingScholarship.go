package teachingscholarship

import (
	class "sigaa.ufpe/packages/data/class_data"
	professor "sigaa.ufpe/packages/data/professor_data"
	scholarship "sigaa.ufpe/packages/data/scholarship_data"
)


type TeachingScholarship struct{
	Scholarship scholarship.Scholarship
	Class class.Class 
	Semester string
	Professor professor.Professor

}