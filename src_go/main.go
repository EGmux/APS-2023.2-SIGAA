package main

import "sigaa.ufpe/controllers"

var db = make(map[string]string)

func main() {
	controllers.Set_Login_Controller()
	// Listen and Server in 0.0.0.0:8080
}
