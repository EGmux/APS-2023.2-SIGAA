package scholarshipcontroller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"scholarship/facade"
	"scholarship/student"

	"github.com/gin-gonic/gin"
	consul "github.com/hashicorp/consul/api"
)


func scholarshipcontroller(){

	r := gin.Default()

	r.LoadHTMLFiles("./view/*")

	r.GET("/scholarship", func(ctx *gin.Context) {
		var students []student.Student
		username := ctx.Query("studentUser")
		log.Print("Username received: ",username)
		students = facade.GetUser(username)
		var student student.Student = students[0]
		fmt.Println(student)
		fmt.Println(student.Disciplines)

		//historic2 := strings.Replace(student.Historic, "\n", "<br>",-1)
		//historic := template.HTML(strings.Replace(historic2,"\t", "&nbsp;&nbsp;&nbsp;&nbsp;",-1))
		historic := "This is a sample historic"
		ctx.HTML(http.StatusOK, "mainMenu.html", gin.H{
			"student":student,
			"historic":historic,
		})
	})

	registerServiceWithConsul()

	r.Run(":8081")
}

func registerServiceWithConsul(){
	config := consul.DefaultConfig()
	config.Address = os.Getenv("CONSUL_HOST")+":"+os.Getenv("CONSUL_PORT")
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

func Set_scholarship_controller(){

	scholarshipcontroller()
}
