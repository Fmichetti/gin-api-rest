package routes

import (
	"github.com/FMichetti/api-go-gin/controllers"
	"github.com/FMichetti/api-go-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func AvaliacaoRequest(r *gin.Engine) {

	protected := r.Group("/api")

	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/avaliacoes", controllers.CriarAvaliacao)
	protected.GET("/avaliacoes", controllers.ListQuestoesAvaliacao)
	protected.DELETE("/avaliacoes/:id", controllers.ExcluirAvaliacao)
}
