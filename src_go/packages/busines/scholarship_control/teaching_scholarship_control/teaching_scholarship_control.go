package teachingscholarshipcontrol

import (
	teachingscholarshipsql "sigaa.ufpe/packages/data/scholarship_data/teachingScholarship/teachingScholarship_SQL"
	"sigaa.ufpe/packages/utils/observer"
)

func Apply_Student_To_Teaching_Scholarship(User string, Scholarship_Id string) {
	Scholarships := teachingscholarshipsql.Get_Teaching_Scholarship_By_Id(Scholarship_Id)
	Scholarship := Scholarships[0]
	teachingscholarshipsql.Associate_Student_To_Teaching_Scholarship(
		Scholarship.Scholarship.Id,
		User,
	)

	observer.Notify_From_Teaching_Scholarship_Update(Scholarship)
}

