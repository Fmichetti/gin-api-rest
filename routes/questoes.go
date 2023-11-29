package routes

import (
	"github.com/FMichetti/api-go-gin/controllers"
	"github.com/FMichetti/api-go-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func QuestoesRequest(r *gin.Engine) {

	protected := r.Group("/api")

	protected.Use(middlewares.JwtAuthMiddleware("professor"))
	protected.POST("/questoes", controllers.CriarQuestao)
	protected.GET("/questoes", controllers.ListarQuestoes)
	protected.GET("/questoes/:questao_id/alternativas", controllers.ListAlternativasQuestao)
	protected.DELETE("/questoes/:questao_id/", controllers.ExcluirQuestao)

}
