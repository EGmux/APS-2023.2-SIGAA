package main_menu_controller

import (
	//"fmt"
	"fmt"
	"html/template"
	"net/http"
	"strings"

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
		historic2 := strings.Replace(student.Historic, "\n", "<br>",-1)
		historic := template.HTML(strings.Replace(historic2,"\t", "&nbsp;&nbsp;&nbsp;&nbsp;",-1))
		ctx.HTML(http.StatusOK, "mainMenu.html", gin.H{
			"student":student,
			"historic":historic,
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

	r.GET("/certification", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:8086/certification?studentUser="+username)
	})


	return r

}





func Set_Main_Menu_Controller(){

	r := main_menu_controller()
	r.Run(":8081")
}