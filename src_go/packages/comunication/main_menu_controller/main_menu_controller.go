package main_menu_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main_menu_controller() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("packages/gui/view/MainMenu/*")

	r.GET("/mainMenu", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "mainMenu.html", gin.H{})
	})

	r.GET("/scholar", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:8082/scholarship")
	})

	return r
}

func Set_Main_Menu_Controller() {
	r := main_menu_controller()
	r.Run(":8081")
}

