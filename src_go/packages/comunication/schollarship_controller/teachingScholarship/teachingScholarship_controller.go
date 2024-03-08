package teachingscholarshipcontroller

import (
	"fmt"
	"net/http"

	//"strconv"

	"github.com/gin-gonic/gin"
	"sigaa.ufpe/packages/busines/facade"
)

func teachingScholarship_controller() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("packages/gui/view/Schollarship//*")

	r.GET("/teachingScholarship", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "teachingScholarship.html", gin.H{
			"title":                "Teaching Scholarships",
			"teachingScholarships": facade.GetAvailableTeachingScholarships(),
		})
	})

	r.POST("/teachingScholarship/submitScholarship", func(ctx *gin.Context) {
		Id := ctx.PostForm("Id")
		User := ctx.PostForm("User")

		//Int_Id, err := strconv.Atoi(Id)
		//if err != nil{
		//	panic(err)
		//}
		fmt.Println(Id, User)
		facade.Apply_Student_To_Teaching_Scholarship(User, Id)

		ctx.String(http.StatusOK, "Bind bem sucesdido, Id recebido: %v", Id)
	})

	return r
}

func Set_Teaching_Scholarship_Controller() {
	r := teachingScholarship_controller()
	r.Run(":8083")
}
