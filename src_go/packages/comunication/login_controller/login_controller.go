package login_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sigaa.ufpe/packages/busines/facade"
	"sigaa.ufpe/packages/data/students_data"
)

// Create API endpoints for user signup and login
func loginController() *gin.Engine {
	r := gin.Default()

	// Necessary /*/* to mach the first set of dirs
	// and the second to match the .html in each dir
	r.LoadHTMLGlob("packages/view/Login/*")

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.POST("/login/auth", func(ctx *gin.Context) {
		var student student.Student
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
		var newStudent student.Student
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

	return r
}


func Set_Login_Controller() {
	
	r := loginController()
	r.Run(":8080")
}
