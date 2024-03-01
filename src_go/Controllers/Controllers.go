package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sigaa.ufpe/login"
)

// Create API endpoints for user signup and login
func loginController() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("VIEW/login")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.GET("/credentials", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signIn.html", gin.H{
			"title": "User list",
			"users": login.DB,
		})
	})

	return r
}

func Set_Login_Controller() {
	r := loginController()
	r.Run(":8080")
}
