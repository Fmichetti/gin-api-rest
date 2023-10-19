package routes

import (

	// docs "github.com/FMichetti/api-go-gin/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {

	r := gin.Default()

	AlunosRequest(r)
	AuthRequest(r)
	UserRequest(r)
	ProfessoresRequest(r)
	QuestoesRequest(r)

	// add swagger
	// swagger url:
	// http://localhost:8080/docs/index.html
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
