package enrollmentcontroller

import (
	"fmt"
	"log"
	"main/facade"
	"main/utils/class"
	"main/utils/student"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	consul "github.com/hashicorp/consul/api"
)

func enrollmentcontroller(){

	r := gin.Default()

	r.LoadHTMLGlob("./view/*")
	var username string
	var students []student.Student
	var student student.Student
	var classes []class.Class

	r.GET("/enrollment", func(ctx *gin.Context) {
		username = ctx.Query("studentUser")
		fmt.Println("------- Student Received:", username)
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


	r.GET("/deferral", func(ctx *gin.Context) {
		fmt.Println("-------- INSIDE deferral")
		fmt.Println("Student Received: ", student)
		fmt.Println(student)
		ctx.HTML(http.StatusOK, "deferral.html", gin.H{
			"student": student,
		})
	})

	
	r.POST("/deferral/apply", func(ctx *gin.Context) {
		username := ctx.Query("studentUser")
		//student := facade.GetUser(username)
		facade.Update_Student_Deferral(username)
		
		ctx.String(http.StatusAccepted, "Student: "+username+" is now on deferral")
	})

	r.POST("/deferral/reEnroll", func(ctx *gin.Context) {
		username := ctx.Query("studentUser")

		facade.Retake_Student_Enrollment(username)

		ctx.String(http.StatusAccepted, "Student: "+username+ "is reenrolled")


	})

	registerServiceWithConsul()
	r.Run(":8082")
}






func registerServiceWithConsul() {
	config := consul.DefaultConfig()
	config.Address = os.Getenv("CONSUL_HOST")+":"+os.Getenv("CONSUL_PORT")
	fmt.Println(config.Address)

	consulClient, err := consul.NewClient(config)
	if err != nil {
		log.Fatal("Error creating Consul client: ", err)
	}

	// Criar um novo serviço
	registration := new(consul.AgentServiceRegistration)
	registration.ID = "enrollment"
	registration.Name = "enrollment"
	registration.Port = 8082 // Porta do serviço

	// Registrar o serviço
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("Error registering service with Consul: ", err)
	}

	fmt.Println("Service registered with Consul")
}

func Set_Enrollment_Controller(){
	enrollmentcontroller()
}