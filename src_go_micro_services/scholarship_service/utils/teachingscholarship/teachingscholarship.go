package teachingscholarship

import (
	"scholarship/utils/class"
	"scholarship/utils/professor"
	"scholarship/utils/scholarship"
)

type TeachingScholarship struct{
	Scholarship scholarship.Scholarship
	Class class.Class
	Semester string
	Professor professor.Professor

}