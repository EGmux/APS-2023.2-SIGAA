package adapter

import (
	"os"
	apigmailprofessor "scholarship/apiGmailProfessor"
	"scholarship/utils/teachingscholarship"
	"strconv"
)



func Adapt_Teaching_To_Professor_API(Scholarship teachingscholarship.TeachingScholarship) apigmailprofessor.Professor_Email_API{

	var Professor_Email apigmailprofessor.Professor_Email_API

	Professor_Email.Professor = Scholarship.Professor
	Professor_Email.Subject = "New aplication for Teaching Scholarship!"
	Professor_Email.Message = " The Student " + Scholarship.Scholarship.Student.User + " have applied for the teaching scholarship of the class " + Scholarship.Class.Name + " for the semester " + Scholarship.Semester + " under the professor " + Scholarship.Professor.Name + ". The schollarship is under the value of " + strconv.Itoa(Scholarship.Scholarship.Value)
	Professor_Email.Scholarship = Scholarship
	Professor_Email.Student = Scholarship.Scholarship.Student
	Professor_Email.Professor.Email = os.Getenv("EMAIL_RECEIVER")

	return Professor_Email
}