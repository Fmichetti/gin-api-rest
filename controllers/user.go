package controllers

import (
	"net/http"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeUsuarioPorID(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	config.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuario não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
