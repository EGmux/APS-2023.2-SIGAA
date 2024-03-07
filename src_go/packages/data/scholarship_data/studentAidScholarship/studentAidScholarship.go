package studentaidscholarship

import scholarship "sigaa.ufpe/packages/data/scholarship_data"

type StudentAidScholarship struct{
	Scholarship scholarship.Scholarship
	Partial bool
}