package routes

import (

	// docs "github.com/FMichetti/api-go-gin/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {

	r := gin.Default()

	// CORS middleware with custom configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Add your frontend's origin here
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	}))

	AlunosRequest(r)
	AuthRequest(r)
	UserRequest(r)
	ProfessoresRequest(r)
	QuestoesRequest(r)
	TurmasRequest(r)

	// add swagger
	// swagger url:
	// http://localhost:8080/docs/index.html
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
