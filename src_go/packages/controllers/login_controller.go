package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sigaa.ufpe/packages/login"
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

	r.GET("/signUp", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signUp.html", gin.H{
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
