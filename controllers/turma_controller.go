package controllers

import (
	"net/http"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// @Summary Cria uma nova Turma
// @Tags Questoes
// @Produce json
// @Param turma body models.Turma true "Objeto Turma a ser criado"
// @Success 201 {object} models.Turma
// @Router /questoes [post]
func CriarTurma(c *gin.Context) {
	var turma models.Turma

	// Parse JSON request body into Turma struct
	if err := c.ShouldBindJSON(&turma); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// Execute a criação no banco de dados usando o GORM
	if err := config.DB.Create(&turma).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, turma)

}

// @Summary Lista todos as turmas
// @Tags Turmas
// @Produce json
// @Success 200 {array} models.Turma
// @Router /turmas [get]
func ExibeTodasTurmas(c *gin.Context) {
	var turmas []models.Turma

	config.DB.Find(&turmas)

	c.JSON(http.StatusOK, turmas)
}
