package main

import (
	"sigaa.ufpe/packages/controllers"
	"sigaa.ufpe/packages/utils/abstractfactory/Assets"
)

var db = make(map[string]string)

func main() {
	println("test 🤛")

	credentials.InitDB()
	controllers.Set_Login_Controller()
	// Listen and Server in 0.0.0.0:8080
}
