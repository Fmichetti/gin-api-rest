package routes

import (
	"github.com/FMichetti/api-go-gin/controllers"
	"github.com/FMichetti/api-go-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRequest(r *gin.Engine) {
	protected := r.Group("/api")

	protected.Use(middlewares.JwtAuthMiddleware("professor"))

	protected.GET("/user/:id", controllers.ExibeUsuarioPorID)

	protected.GET("/users", controllers.ExibeTodosUsers)
}
