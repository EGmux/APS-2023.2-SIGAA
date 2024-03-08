package utils

import (
	API_GMAIL_Professor "sigaa.ufpe/pkgs/api_gmail_professor"
	teachingscholarship "sigaa.ufpe/pkgs/data/scholarship_data/teachingScholarship"
)

func Notify_From_Teaching_Scholarship_Update(
	TeachingScholarship teachingscholarship.TeachingScholarship,
) {
	API_GMAIL_Professor.Send_Professor_email(
		Adapt_Teaching_To_Professor_API(TeachingScholarship),
	)
}
