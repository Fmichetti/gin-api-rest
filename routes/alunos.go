package routes

import (
	"github.com/FMichetti/api-go-gin/controllers"
	"github.com/FMichetti/api-go-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func AlunosRequest(r *gin.Engine) {

	protected := r.Group("/api")

	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/alunos", controllers.ExibeTodosAlunos)
	// protected.POST("/alunos", controllers.CriaNovoAluno)
	// protected.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	// protected.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	// protected.PATCH("/alunos/:id", controllers.EditaAluno)
	// protected.DELETE("/alunos/:id", controllers.DeletaAlunoPorID)
}
