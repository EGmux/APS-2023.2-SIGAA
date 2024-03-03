package packages_data_students_data_students_json


type Student struct {
	User     string
	Password string
}

var DB = []Student{
	{User: "hac", Password: "1234"},
	{User: "Enzo", Password: "ola"},
}