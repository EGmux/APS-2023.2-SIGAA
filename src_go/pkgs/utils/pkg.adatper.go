package utils

//
// import (
// 	"os"
// 	"strconv"
//
// 	API_GMAIL_Professor "sigaa.ufpe/pkgs/api_gmail_professor"
// 	// teachingscholarship "sigaa.ufpe/pkgs/data/scholarship_data/teachingScholarship"
// )
//
// func Adapt_Teaching_To_Professor_API(
// 	Scholarship teachingscholarship.TeachingScholarship,
// ) API_GMAIL_Professor.Professor_Email_API {
// 	var Professor_Email API_GMAIL_Professor.Professor_Email_API
//
// 	Professor_Email.Professor = Scholarship.Professor
// 	Professor_Email.Subject = "New aplication for Teaching Scholarship!"
// 	Professor_Email.Message = " The Student " + Scholarship.Scholarship.Student.User + " have applied for the teaching scholarship of the class " + Scholarship.Class.Name + " for the semester " + Scholarship.Semester + " under the professor " + Scholarship.Professor.Name + ". The schollarship is under the value of " + strconv.Itoa(
// 		Scholarship.Scholarship.Value,
// 	)
// 	Professor_Email.Scholarship = Scholarship
// 	Professor_Email.Student = Scholarship.Scholarship.Student
// 	Professor_Email.Professor.Email = os.Getenv("EMAIL_RECEIVER")
//
// 	return Professor_Email
// }
