package main

import (
	"sigaa.ufpe/packages/controllers"
	"sigaa.ufpe/packages/utils/queries"
)

var db = make(map[string]string)

func main() {
	println("test 🤛")
<<<<<<< HEAD
	queries.TableExists("credentials")
=======
	queries.CheckExistence("credentials")
>>>>>>> 95fd195 (update(main.go): testcode)
	controllers.Set_Login_Controller()
	// Listen and Server in 0.0.0.0:8080
}
