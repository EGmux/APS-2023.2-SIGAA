package observer

import (
	"scholarship/adapter"
	apigmailprofessor "scholarship/apiGmailProfessor"
	"scholarship/utils/teachingscholarship"
)

func Notify_From_Teaching_Scholarship_Update(TeachingScholarship teachingscholarship.TeachingScholarship){


	apigmailprofessor.Send_Professor_email(adapter.Adapt_Teaching_To_Professor_API(TeachingScholarship))


}