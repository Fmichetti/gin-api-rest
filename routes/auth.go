package routes

import (
	"github.com/FMichetti/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRequest(r *gin.Engine) {
	r.POST("/register", controllers.RegistraAluno)
	r.GET("/login", controllers.Login)
}
