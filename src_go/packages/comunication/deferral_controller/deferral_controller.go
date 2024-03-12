package deferralcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sigaa.ufpe/packages/busines/facade"
)


func deferral_controller() *gin.Engine{

	r := gin.Default()

	r.LoadHTMLGlob("packages/view/Deferral/*")

	r.GET("/deferral", func(ctx *gin.Context) {
		student_name := ctx.Query("studentUser")
		student := facade.GetUser(student_name)[0]
		fmt.Println(student)
		ctx.HTML(http.StatusOK, "deferral.html", gin.H{
			"student": student,
		})
	})

	r.POST("/deferral/apply", func(ctx *gin.Context) {
		username := ctx.Query("studentUser")
		//student := facade.GetUser(username)
		facade.Update_Student_Deferral(username)
		
		ctx.String(http.StatusAccepted, "Student: "+username+" have is now on deferral")
	})

	r.POST("/deferral/reEnroll", func(ctx *gin.Context) {
		username := ctx.Query("studentUser")

		facade.Retake_Student_Enrollment(username)

		ctx.String(http.StatusAccepted, "Student: "+username+ "is reenrolled")


	})

	return r

}



func Set_Deferral_Controller(){

	r := deferral_controller()
	r.Run(":8085")
}