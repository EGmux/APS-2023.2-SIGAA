package scholarshipcontrol

import (
	scholarshipsql "scholarship/scholarship_sql"
	"scholarship/student"
)


func GetUser(username string) []student.Student{
	return scholarshipsql.GetUser(username)
}