package controllers

import (
	"net/http"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// @Description Cria uma nova Questao
// @Tags Questoes
// @Produce json
// @Param questao body models.Questao true "Questao object"
// @Success 200 {object} []models.Questao
// @Router /questoes [post]
func CriarQuestao(c *gin.Context) {
	var questao models.Questao

	// Parse JSON request body into Questao struct
	if err := c.ShouldBindJSON(&questao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// Execute a criação no banco de dados usando o GORM
	if err := config.DB.Create(&questao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, questao)

}

// @Description Exibe todas as Questoes
// @Tags Questoes
// @Produce json
// @Success 200 {object} []models.Questao
// @Router /questoes [get]
func ListarQuestoes(c *gin.Context) {
	var questoes []models.Questao

	config.DB.Find(&questoes)

	c.JSON(http.StatusOK, questoes)
}

// ListAlternativas lista as alternativas de uma determinada questão
func ListAlternativasQuestao(c *gin.Context) {
	questaoID := c.Param("questao_id") // Recupera o ID da questão a partir dos parâmetros da URL

	var questao models.Questao
	if err := config.DB.Preload("Alternativas").First(&questao, questaoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Questão não encontrada"})
		return
	}

	c.JSON(http.StatusOK, questao.Alternativas)
}
