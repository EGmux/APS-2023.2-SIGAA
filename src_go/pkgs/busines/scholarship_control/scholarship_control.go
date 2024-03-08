package scholarshipcontrol

import (
	"sigaa.ufpe/pkgs/data/repo/structs"
	teachingscholarshipsql "sigaa.ufpe/pkgs/data/scholarship_data/teachingScholarship/teachingScholarship_SQL"
)

func GetAvailableTeachingScholarships() []structs.TeachingScholarship {
	return teachingscholarshipsql.GetAvailableTeachingScholarships()
}

func GetAllTeachingScholarships() []structs.TeachingScholarship {
	return teachingscholarshipsql.GetAllTeachingScholarships()
}
