package scholarshipcontrol

import (
	"scholarship/observer"
	scholarshipsql "scholarship/scholarship_sql"
	"scholarship/utils/student"
	"scholarship/utils/teachingscholarship"
)


func GetUser(username string) []student.Student{
	return scholarshipsql.GetUser(username)
}

func GetAllTeachingScholarships()[]teachingscholarship.TeachingScholarship{
	return scholarshipsql.GetAllTeachingScholarships()
}

func GetAvailableTeachingScholarships() []teachingscholarship.TeachingScholarship{
	return scholarshipsql.GetAvailableTeachingScholarships()
}

func Apply_Student_To_Teaching_Scholarship(User string, Scholarship_Id string)  {


	Scholarships := scholarshipsql.Get_Teaching_Scholarship_By_Id(Scholarship_Id)
	Scholarship := Scholarships[0]
	scholarshipsql.Associate_Student_To_Teaching_Scholarship(Scholarship.Scholarship.Id, User)

	observer.Notify_From_Teaching_Scholarship_Update(Scholarship)

}