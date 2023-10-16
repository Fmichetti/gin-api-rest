package controllers

import (
	"net/http"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// @Description Exibe todos os alunos
// @Tags Alunos
// @Produce json
// @Success 200 {object} []models.Aluno
// @Router /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno

	config.DB.Find(&alunos)

	c.JSON(http.StatusOK, alunos)
}

// @Description Cria um novo aluno
// @Tags Alunos
// @Produce json
// @Param aluno body models.Aluno true "Aluno object"
// @Success 200 {object} []models.Aluno
// @Router /alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//verifica se foi adicionado algum aluno
	if err := config.DB.Create(&aluno).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, aluno)
}

// @Description Realiza uma busca baseada pelo ID
// @Tags Alunos
// @Produce json
// @Success 200 {object} models.Aluno
// @Router /alunos/{id} [get]
func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	config.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, aluno)
}

// @Description Realiza uma busca baseada pelo Nome
// @Tags Alunos
// @Produce json
// @Success 200 {object} models.Aluno
// @Router /alunos/{nome} [get]
func BuscaAlunoPorNome(c *gin.Context) {
	var alunos []models.Aluno
	var requestBody struct {
		Nome string `json:"nome"`
	}

	// Bind JSON data from the request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Query the database using a partial match on the "nome" field
	config.DB.Where("nome LIKE ?", "%"+requestBody.Nome+"%").Find(&alunos)

	if len(alunos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, alunos)
}

// @Description Realiza uma busca baseada pelo ID
// @Tags Alunos
// @Produce json
// @Success 200
// @Router /alunos/{id} [delete]
func DeletaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db := config.DB.Delete(&aluno, id)

	//verifica se foi removido algum aluno
	if db.RowsAffected < 1 {

		c.JSON(http.StatusOK, gin.H{
			"data": "Aluno não existente!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso!",
	})
}

// @Description Realiza uma edição baseada pelo ID
// @Tags Alunos
// @Produce json
// @Param aluno body models.Aluno true "Aluno object"
// @Success 200
// @Router /alunos/{id} [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	db := config.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	config.DB.Model(&aluno).UpdateColumns(aluno)

	//verifica se foi editado algum aluno
	if db.RowsAffected < 1 {

		c.JSON(http.StatusOK, gin.H{
			"data": "Aluno não existente!",
		})

		return
	}

	c.JSON(http.StatusOK, aluno)
}
