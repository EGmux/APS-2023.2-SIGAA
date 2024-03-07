package scholarshipcontrol

import (
	teachingscholarship "sigaa.ufpe/packages/data/scholarship_data/teachingScholarship"
	teachingscholarshipsql "sigaa.ufpe/packages/data/scholarship_data/teachingScholarship/teachingScholarship_SQL"
)


func GetAvailableTeachingScholarships() []teachingscholarship.TeachingScholarship{
	return teachingscholarshipsql.GetAvailableTeachingScholarships()
}

func GetAllTeachingScholarships() []teachingscholarship.TeachingScholarship{
	return teachingscholarshipsql.GetAllTeachingScholarships()
}