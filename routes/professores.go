package routes

import (
	"github.com/FMichetti/api-go-gin/controllers"
	"github.com/FMichetti/api-go-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func ProfessoresRequest(r *gin.Engine) {

	protected := r.Group("/api")

	protected.Use(middlewares.JwtAuthMiddleware("professor"))
	protected.GET("/professores", controllers.ExibeTodosProfessores)
	protected.POST("/professores", controllers.CriaNovoProfessor)
	protected.GET("/professores/:id", controllers.BuscaProfessorPorID)
	protected.GET("/professores/name", controllers.BuscaProfessorPorNome)
	protected.DELETE("/professores/:id", controllers.DeletaProfessorPorID)
	protected.PATCH("/professores/:id", controllers.EditaProfessor)
}
