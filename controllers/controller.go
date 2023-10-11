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

// // @Description Cria um novo aluno
// // @Tags Alunos
// // @Produce json
// // @Param aluno body models.Aluno true "Aluno object"
// // @Success 200 {object} []models.Aluno
// // @Router /alunos [post]
// func CriaNovoAluno(c *gin.Context) {
// 	var aluno models.Aluno
// 	if err := c.ShouldBindJSON(&aluno); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	config.DB.Create(&aluno)
// 	c.JSON(http.StatusOK, aluno)
// }

// // @Description Realiza uma busca baseada pelo ID
// // @Tags Alunos
// // @Produce json
// // @Success 200 {object} models.Aluno
// // @Router /alunos/{id} [get]
// func BuscaAlunoPorID(c *gin.Context) {
// 	var aluno models.Aluno
// 	id := c.Params.ByName("id")
// 	config.DB.First(&aluno, id)

// 	if aluno.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "Aluno não encontrado",
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, aluno)
// }

// // @Description Realiza uma busca baseada pelo ID
// // @Tags Alunos
// // @Produce json
// // @Success 200
// // @Router /alunos/{id} [delete]
// func DeletaAlunoPorID(c *gin.Context) {
// 	var aluno models.Aluno
// 	id := c.Params.ByName("id")
// 	db := config.DB.Delete(&aluno, id)

// 	//verifica se foi removido algum aluno
// 	if db.RowsAffected < 1 {

// 		c.JSON(http.StatusOK, gin.H{
// 			"data": "Aluno não existente!",
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": "Aluno deletado com sucesso!",
// 	})
// }

// // @Description Realiza uma edição baseada pelo ID
// // @Tags Alunos
// // @Produce json
// // @Param aluno body models.Aluno true "Aluno object"
// // @Success 200
// // @Router /alunos/{id} [patch]
// func EditaAluno(c *gin.Context) {
// 	var aluno models.Aluno
// 	id := c.Params.ByName("id")
// 	db := config.DB.First(&aluno, id)

// 	if err := c.ShouldBindJSON(&aluno); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}

// 	config.DB.Model(&aluno).UpdateColumns(aluno)

// 	//verifica se foi editado algum aluno
// 	if db.RowsAffected < 1 {

// 		c.JSON(http.StatusOK, gin.H{
// 			"data": "Aluno não existente!",
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, aluno)
// }

// // @Description Realiza uma busca baseada pelo CPF
// // @Tags Alunos
// // @Produce json
// // @Success 200 {object} models.Aluno
// // @Router /alunos/{cpf} [get]
// func BuscaAlunoPorCPF(c *gin.Context) {
// 	var aluno models.Aluno
// 	cpf := c.Params.ByName("cpf")
// 	config.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

// 	if aluno.CPF == "" {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "Aluno não encontrado",
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, aluno)
// }
