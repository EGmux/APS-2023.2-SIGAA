package scholarshipcontroller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"scholarship/facade"
	"scholarship/utils/student"

	"github.com/gin-gonic/gin"
	consul "github.com/hashicorp/consul/api"
)

func scholarshipcontroller() {

	var username string
	r := gin.Default()

	r.LoadHTMLGlob("./view/*")

	r.GET("/scholarship", func(ctx *gin.Context) {
		var students []student.Student
		username = ctx.Query("studentUser")
		log.Print("Username received: ", username)
		students = facade.GetUser(username)
		var student student.Student = students[0]
		fmt.Println(student)
		fmt.Println(student.Disciplines)

		//historic2 := strings.Replace(student.Historic, "\n", "<br>",-1)
		//historic := template.HTML(strings.Replace(historic2,"\t", "&nbsp;&nbsp;&nbsp;&nbsp;",-1))
		historic := "This is a sample historic"
		ctx.HTML(http.StatusOK, "mainMenu.html", gin.H{
			"student":  student,
			"historic": historic,
		})
	})

	r.GET("/scholar", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "schollarship_main.html", gin.H{
			"title":                "Scholarships",
			"title1":               "Teaching Scholarships",
			"teachingScholarships": facade.GetAllTeachingScholarships(),
		})
	})

	r.GET("/teach", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "teachingScholarship.html", gin.H{
			"title":                "Teaching Scholarships",
			"teachingScholarships": facade.GetAvailableTeachingScholarships(),
		})
	})

	r.POST("/teach/submit", func(ctx *gin.Context) {
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

	registerServiceWithConsul()

	r.Run(":8081")
}

func registerServiceWithConsul() {
	config := consul.DefaultConfig()
	config.Address = os.Getenv("CONSUL_HOST") + ":" + os.Getenv("CONSUL_PORT")
	fmt.Println(config.Address)

	consulClient, err := consul.NewClient(config)
	if err != nil {
		log.Fatal("Error creating Consul client: ", err)
	}

	// Criar um novo serviço
	registration := new(consul.AgentServiceRegistration)
	registration.ID = "scholarship"
	registration.Name = "scholarship"
	registration.Port = 8081 // Porta do serviço

	// Registrar o serviço
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("Error registering service with Consul: ", err)
	}

	fmt.Println("Service registered with Consul")

}

func Set_scholarship_controller() {

	scholarshipcontroller()
}
