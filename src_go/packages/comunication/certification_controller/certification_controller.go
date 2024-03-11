package certificationcontroller

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"sigaa.ufpe/packages/busines/facade"
)


func certification_controller() *gin.Engine{

	r := gin.Default()
	var Certification string
	var username string

	r.LoadHTMLGlob("packages/view/Certification/*")

	r.GET("/certification", func(ctx *gin.Context) {
		username = ctx.Query("studentUser")
		students := facade.GetUser(username)
		student := students[0]
		historic2 := strings.Replace(student.Historic, "\n", "<br>",-1)
		historic := template.HTML(strings.Replace(historic2,"\t", "&nbsp;&nbsp;&nbsp;&nbsp;",-1))

		classes_information := facade.GetSelectedClasses(student.Disciplines)

		Certification = "The Federal University of Pernambuco\n\n Cetificates for the due purpouses tha teh student "+username+" is regularly enrolled in our university, being hims/hers activity schedule:\n"

		for _, class := range(classes_information){

			Certification += "Class: "+class.Name+"\n"
			Certification += "\tTimetable: "+class.Timetable+"\n"
			Certification += "\tProfessor: "+class.Professor.Name+"\n"
			Certification += "\tSemester: "+class.Semester+"\n"

		}

		Certification2 := strings.Replace(Certification, "\n", "<br>",-1)
		Certification3 := template.HTML(strings.Replace(Certification2, "\t","&nbsp;&nbsp;&nbsp;&nbsp;",-1))

		ctx.HTML(http.StatusOK, "certification.html", gin.H{
			"username":username,
			"student":student,
			"historic":historic,
			"certification":Certification3,
		})

	})

	r.GET("/certification/download", func(ctx *gin.Context) {

		tempFile, err := os.CreateTemp(os.TempDir(),"certificate.txt")
		if err != nil{
			ctx.String(http.StatusInternalServerError, "Error while creating temporary file")
			return
		}
		defer tempFile.Close()

		_, err = tempFile.WriteString(Certification)
		if err != nil{
			ctx.String(http.StatusInternalServerError, "Error while writing on temporary file")
			return
		}

		ctx.Header("Content-Disposition", "attachment; filename="+username+"_enrollment_certificate.txt")
		ctx.File(tempFile.Name())
	})

	return r
}


func Set_certification_controller(){

	r := certification_controller()
	r.Run(":8086")
}