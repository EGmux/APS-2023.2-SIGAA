package enrollmentcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sigaa.ufpe/packages/busines/facade"
	class "sigaa.ufpe/packages/data/class_data"
	student "sigaa.ufpe/packages/data/students_data"
)

func enrollment_controller()*gin.Engine {

	r := gin.Default()
	var username string
	var students []student.Student
	var student student.Student
	var classes []class.Class
	r.LoadHTMLGlob("packages/view/Enrollment/*")
	

	r.GET("/enrollment", func(ctx *gin.Context) {
		username = ctx.Query("studentUser")
		students = facade.GetUser(username)
		student = students[0]
		classes = facade.GetAllClasses()
		fmt.Println(classes)

		ctx.HTML(http.StatusOK, "enrollment.html",gin.H{
			"username":username,
			"student":student,
			"classes":classes,
		})
	})

	r.POST("/enrollment/apply", func(ctx *gin.Context) {
		studentUser := ctx.Query("studentUser")
		fmt.Println("Student:",studentUser)
		selection := ctx.PostFormArray("selectedClasses")
		if len(selection) == 0 {
			ctx.String(http.StatusBadRequest, "No class selected")
			return
		}
		classes = facade.GetSelectedClasses(selection)

		facade.Update_Student_Enrollment(studentUser, classes)

		ctx.String(http.StatusOK, "enrollment apply!")
	})


	return r
}



func Set_Enrollment_Controller(){
	r := enrollment_controller()
	r.Run(":8084")

}