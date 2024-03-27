package main

import (
	sql "main/SQL"
	enrollmentcontroller "main/enrollmentController"
)


func main(){

	canal := make(chan bool)
	
	sql.InitDB()
	enrollmentcontroller.Set_Enrollment_Controller()


	<- canal
}