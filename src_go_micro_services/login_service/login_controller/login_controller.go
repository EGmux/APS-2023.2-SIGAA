package logincontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func logincontroller() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("./view/login.html")

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})
	return r
}

func Set_Login_Controller() {

	r := logincontroller()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Could not run the Login controller")
	}
}
