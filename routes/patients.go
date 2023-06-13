package routes

import (
	"github.com/FMichetti/api-go-gin/controllers"
	"github.com/FMichetti/api-go-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func PatientsRequest(r *gin.Engine) {

	protected := r.Group("/api")

	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/patient", controllers.ExibeTodosPacientes)
	protected.GET("/patient/:id", controllers.BuscaPacientePorID)
	protected.GET("/patient/cpf/:cpf", controllers.BuscaPacientePorCPF)
	protected.POST("/patient", controllers.CriaNovoPaciente)
	protected.PATCH("/patient/:id", controllers.EditaPaciente)
	protected.DELETE("/patient/:id", controllers.DeletaPacientePorID)
}
