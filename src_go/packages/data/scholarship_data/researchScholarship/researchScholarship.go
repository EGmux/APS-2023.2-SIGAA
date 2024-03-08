package researchscholarship

import (
	professor "sigaa.ufpe/packages/data/professor_data"
	scholarship "sigaa.ufpe/packages/data/scholarship_data"
)

type ResearchScholarship struct{
	
	Scholarship scholarship.Scholarship
	Semester string
	Topic string
	Professor professor.Professor
}