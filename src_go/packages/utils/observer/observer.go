package observer

import (
	API_GMAIL_Professor "sigaa.ufpe/packages/api_gmail_professor"
	teachingscholarship "sigaa.ufpe/packages/data/scholarship_data/teachingScholarship"
	"sigaa.ufpe/packages/utils/adapter"
)


func Notify_From_Teaching_Scholarship_Update(TeachingScholarship teachingscholarship.TeachingScholarship){


	API_GMAIL_Professor.Send_Professor_email(adapter.Adapt_Teaching_To_Professor_API(TeachingScholarship))
}