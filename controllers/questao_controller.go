package controllers

import (
	"net/http"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// @Summary Cria uma nova Questao
// @Tags Questoes
// @Produce json
// @Param questao body models.Questao true "Objeto Questao a ser criado"
// @Success 201 {object} models.Questao
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

// @Summary Exibe todas as Questoes
// @Tags Questoes
// @Produce json
// @Success 200 {array} models.Questao
// @Router /questoes [get]
func ListarQuestoes(c *gin.Context) {
	var questoes []models.Questao

	config.DB.Find(&questoes)

	c.JSON(http.StatusOK, questoes)
}

// ListAlternativasQuestao retrieves a list of alternatives for a specific question.
// @Summary List alternatives for a question
// @Tags Alternativas
// @Produce json
// @Param questao_id path string true "Question ID" Format(uuid)
// @Success 200 {array} models.Alternativa
// @Router /alternativas/{questao_id} [get]
func ListAlternativasQuestao(c *gin.Context) {
	questaoID := c.Param("questao_id") // Recupera o ID da questão a partir dos parâmetros da URL

	var questao models.Questao
	if err := config.DB.Preload("Alternativas").First(&questao, questaoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Questão não encontrada"})
		return
	}

	c.JSON(http.StatusOK, questao.Alternativas)
}

// @Summary Deleta uma questao pelo ID
// @Tags Questao
// @Produce json
// @Param id path string true "ID da questao a ser deleteda"
// @Success 200 {object} gin.H
// @Router /questoes/{id} [delete]
func ExcluirQuestao(c *gin.Context) {
	var questao models.Questao
	id := c.Params.ByName("id")
	db := config.DB.Delete(&questao, id)

	//verifica se aconteceu a remocao

	if db.RowsAffected > 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Aluno nao existente!",
		})
	}
}
