package login

type Student struct {
	User     string
	Password string
}

var DB = []Student{
	{User: "hac", Password: "1234"},
	{User: "Enzo", Password: "ola"},
}
