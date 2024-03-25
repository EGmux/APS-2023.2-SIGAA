package logincontroller

import (
	"fmt"
	"log"
	data_studentsdata_studentsSQL "login_service/SQL"
	"login_service/facade"
	"net/http"

	"github.com/gin-gonic/gin"
	consul "github.com/hashicorp/consul/api"
)



func logincontroller() {

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	r.LoadHTMLGlob("./view/*")

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})


	registerServiceWithConsul()
	r.POST("/login/auth", func(ctx *gin.Context) {
		var student data_studentsdata_studentsSQL.Student
		if err := ctx.Bind(&student); err != nil{
			fmt.Println("Erro de Bind")
			ctx.String(http.StatusBadRequest, "Erro ao processar o formulario: %v", err)
		}

		
		if facade.IsValidUser(student.User, student.Password) {
            // Redirecionar para a página principal em caso de autenticação bem-sucedida
            ctx.Redirect(http.StatusMovedPermanently, "/mainMenu?studentUser="+student.User)
	
        }
		
	})

	r.GET("/signUp", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signUp.html", gin.H{
			"title": "User list",
			"users": facade.GetAllUsers(),
		})
	})

	r.POST("/signUp/insert", func(ctx *gin.Context) {
		var newStudent data_studentsdata_studentsSQL.Student
		if err := ctx.Bind(&newStudent); err != nil {
			ctx.String(http.StatusBadRequest, "Erro ao processar o formulário: %v", err)
		}

		facade.InsertUser(newStudent.User, newStudent.Password)

		ctx.String(http.StatusOK, "Usuario adicionado com sucesso!")
	})

	r.GET("/mainMenu", func(ctx *gin.Context){
		var studentUser = ctx.Query("studentUser")
		ctx.Redirect(http.StatusMovedPermanently, "http://localhost:8081/mainMenu?studentUser="+studentUser)
	})
	
	r.Run(":8080")
}

func registerServiceWithConsul() {
	config := consul.DefaultConfig()
	consulClient, err := consul.NewClient(config)
	if err != nil {
		log.Fatal("Error creating Consul client: ", err)
	}

	// Criar um novo serviço
	registration := new(consul.AgentServiceRegistration)
	registration.ID = "login"
	registration.Name = "login"
	registration.Port = 8080 // Porta do serviço
	registration.Address = "localhost"
	registration.Check = &consul.AgentServiceCheck{
		HTTP:     "http://localhost:8080/health",
		Interval: "10s",
	}

	// Registrar o serviço
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("Error registering service with Consul: ", err)
	}

	fmt.Println("Service registered with Consul")
}

func Set_Login_Controller() {

	logincontroller()
}
