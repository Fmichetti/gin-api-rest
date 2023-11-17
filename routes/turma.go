package routes

import (
	"github.com/FMichetti/api-go-gin/controllers"
	"github.com/FMichetti/api-go-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func TurmasRequest(r *gin.Engine) {

	protected := r.Group("/api")

	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/turmas", controllers.CriarTurma)
	protected.GET("/turmas", controllers.ExibeTodasTurmas)
}
