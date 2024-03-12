package gradescontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sigaa.ufpe/packages/busines/facade"
)



func grades_controller() *gin.Engine{
	
	r := gin.Default()

	var username string

	r.LoadHTMLGlob("packages/view/Grades/*")

	r.GET("/grades", func(ctx *gin.Context) {
		username = ctx.Query("studentUser")

		student := facade.GetUser(username)[0]

		grades := facade.Get_All_Grades_From_Student(username)

		ctx.HTML(http.StatusOK, "grades.html", gin.H{
			"username":username,
			"student":student,
			"grades":grades,

		})
	})


	return r
}


func Set_Grades_Controller(){

	r := grades_controller()
	r.Run(":8087")
}