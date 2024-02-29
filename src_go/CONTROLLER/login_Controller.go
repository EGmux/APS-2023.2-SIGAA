package login_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Aluno struct{
	User string 
	Senha string
}

var db = []Aluno{
	{User: "hac", Senha: "1234"},
	{User: "Enzo", Senha: "ola"},
}

func loginController() *gin.Engine{

	r := gin.Default()
	r.LoadHTMLGlob("VIEW/*")
	r.GET("/login", func(ctx *gin.Context){
		ctx.HTML(http.StatusOK, "login.html",gin.H{

		})

	})
		


	r.GET("/credenciais", func(ctx *gin.Context) {

		
		ctx.HTML(http.StatusOK, "criarCredencial.html", gin.H{
			"title": "Lista de Usuarios",
			"users": db,
		})
		
	})

	
	
	return r
}





func Set_Login_Controller(){
	r := loginController()
	r.Run(":8080")
}