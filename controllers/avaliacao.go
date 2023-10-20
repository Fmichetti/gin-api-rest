package controllers

import (
	"net/http"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// @Summary Cria uma nova Avaliacao
// @Tags Avaliacao
// @Produce json
// @Param avaliacao body models.Avaliacao true "Objeto Avaliacao a ser criado"
// @Success 201 {object} models.Avaliacao
// @Router /questoes [post]
func CriarAvaliacao(c *gin.Context) {
	var avaliacao models.Avaliacao

	// Parse JSON request body into Avaliacao struct
	if err := c.ShouldBindJSON(&avaliacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// Execute a criação no banco de dados usando o GORM
	if err := config.DB.Create(&avaliacao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, avaliacao)
}

// ListQuestoesAvaliacao retrieves a list of questions for a specific test.
// @Summary List questions for a test
// @Tags Avaliacao
// @Produce json
// @Param avaliacao_id path string true "Avaliacao ID" Format(uuid)
// @Success 200 {array} models.Alternativa
// @Router /alternativas/{avaliacao_id} [get]
func ListQuestoesAvaliacao(c *gin.Context) {
	avaliacaoID := c.Param("avaliacao_id") // Recupera o ID da questão a partir dos parâmetros da URL

	var avaliacao models.Avaliacao
	if err := config.DB.Preload("Alternativas").First(&avaliacao, avaliacaoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Avaliação não encontrada"})
		return
	}

	c.JSON(http.StatusOK, avaliacao.Questoes)
}

// @Summary Deleta uma avaliacao pelo ID
// @Tags Avaliacao
// @Produce json
// @Param id path string true "ID da avaliacao a ser deleteda"
// @Success 200 {object} gin.H
// @Router /questoes/{id} [delete]
func ExcluirAvaliacao(c *gin.Context) {
	var avaliacao models.Avaliacao
	id := c.Params.ByName("id")
	db := config.DB.Delete(&avaliacao, id)

	//verifica se aconteceu a remocao

	if db.RowsAffected > 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Avaliação nao existente!",
		})
	}
}
