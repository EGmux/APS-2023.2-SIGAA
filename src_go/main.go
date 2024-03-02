package main

import (
	"sigaa.ufpe/packages/controllers"
	"sigaa.ufpe/packages/utils/queries"
)

var db = make(map[string]string)

func main() {
	println("test 🤛")
	queries.CheckExistence("credentials")
	controllers.Set_Login_Controller()
	// Listen and Server in 0.0.0.0:8080
}
