package login_controller

import (
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
	r.LoadHTMLGlob("packages/login/*/*")

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.POST("/login/auth", func(ctx *gin.Context) {
		var student student.Student
		if err := ctx.Bind(&student); err != nil{
			ctx.String(http.StatusBadRequest, "Erro ao processar o formulario: %v", err)
			return
		}

		
		if facade.IsValidUser(student.User, student.Password) {
            // Redirecionar para a página principal em caso de autenticação bem-sucedida
            ctx.String(http.StatusOK, "O usuario esta autenticado")
            return
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

	return r
}


func Set_Login_Controller() {
	facade.InitDB()
	defer facade.CloseDB()
	r := loginController()
	r.Run(":8080")
}
