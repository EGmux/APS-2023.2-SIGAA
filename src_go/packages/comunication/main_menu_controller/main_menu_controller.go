package main_menu_controller

import (
	//"fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"sigaa.ufpe/packages/busines/facade"
	"sigaa.ufpe/packages/busines/facade"
	student "sigaa.ufpe/packages/data/students_data"
)


func main_menu_controller() *gin.Engine{

	var username string

	r := gin.Default()

	r.LoadHTMLGlob("packages/view/MainMenu/*")

	r.GET("/mainMenu", func(ctx *gin.Context) {
		var students []student.Student
		username = ctx.Query("studentUser")
		students = facade.GetUser(username)
		var student student.Student = students[0]
		fmt.Println(student)
		fmt.Println(student.Disciplines)
		//ctx.String(http.StatusOK, "%s , %s",username, //student)
		ctx.HTML(http.StatusOK, "mainMenu.html", gin.H{
			"student":student,
		})
	})


	r.GET("/scholar", func(ctx *gin.Context)  {
		ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:8082/scholarship?studentUser="+username)
	})

	r.GET("/enrollment", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:8084/enrollment?studentUser="+username)
	})

	r.GET("/deferral", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:8085/deferral?studentUser="+username)
	})


	return r

}





func Set_Main_Menu_Controller(){

	r := main_menu_controller()
	r.Run(":8081")
}