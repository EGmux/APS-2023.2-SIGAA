package main

import (
	"login_controller"
)

var db = make(map[string]string)



func main() {
	
	login_controller.Set_Login_Controller()
	// Listen and Server in 0.0.0.0:8080
	
}