package controllers

import (
	"net/http"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// @Summary Lista todos os professores
// @Tags Professores
// @Produce json
// @Success 200 {array} models.Professor
// @Router /professores [get]
func ExibeTodosProfessores(c *gin.Context) {
	var professores []models.Professor

	config.DB.Find(&professores)

	c.JSON(http.StatusOK, professores)
}

// @Summary Cria um novo professor
// @Tags Professores
// @Produce json
// @Param professor body models.Professor true "Objeto Professor a ser criado"
// @Success 201 {object} models.Professor
// @Router /professores [post]
func CriaNovoProfessor(c *gin.Context) {
	var professor models.Professor
	//verifica se o JSON esta correto
	if err := c.ShouldBindJSON(&professor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//verifica se nao existe nenhum aluno com esse userID
	var aluno models.Aluno
	id := professor.UserID

	config.DB.Where("user_id = ?", id).First(&aluno)
	if aluno.ID != 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User ID nao pode ser atribuido ao professor!",
		})

		return
	}

	//verifica se foi adicionado algum professor
	if err := config.DB.Create(&professor).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, professor)
}

// @Summary Busca um professor pelo ID
// @Tags Professores
// @Produce json
// @Param id path string true "ID do professor"
// @Success 200 {object} models.Professor
// @Router /professores/{id} [get]
func BuscaProfessorPorID(c *gin.Context) {
	var professor models.Professor
	id := c.Params.ByName("id")
	config.DB.First(&professor, id)

	if professor.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Professor não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, professor)
}

// @Summary Busca professores pelo Nome
// @Tags Professores
// @Produce json
// @Param nome query string true "Nome do professor a ser buscado"
// @Success 200 {array} models.Professor
// @Router /professores/{nome} [post]
func BuscaProfessorPorNome(c *gin.Context) {
	var professores []models.Professor
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
	config.DB.Where("nome LIKE ?", "%"+requestBody.Nome+"%").Find(&professores)

	if len(professores) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Professor não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, professores)
}

// @Summary Deleta um professor pelo ID
// @Tags Professores
// @Produce json
// @Param id path string true "ID do professor a ser deletado"
// @Success 200 {object} gin.H
// @Router /professores/{id} [delete]
func DeletaProfessorPorID(c *gin.Context) {
	var professor models.Professor
	id := c.Params.ByName("id")
	db := config.DB.Delete(&professor, id)

	//verifica se foi removido algum professor
	if db.RowsAffected < 1 {

		c.JSON(http.StatusOK, gin.H{
			"data": "Professor não existente!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Professor deletado com sucesso!",
	})
}

// @Summary Edita um professor pelo ID
// @Tags Professores
// @Produce json
// @Param id path string true "ID do professor a ser editado"
// @Param professor body models.Professor true "Objeto Professor com as alterações"
// @Success 200 {object} models.Professor
// @Router /professores/{id} [patch]
func EditaProfessor(c *gin.Context) {
	var professor models.Professor
	id := c.Params.ByName("id")
	db := config.DB.First(&professor, id)

	if err := c.ShouldBindJSON(&professor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	config.DB.Model(&professor).UpdateColumns(professor)

	//verifica se foi editado algum professor
	if db.RowsAffected < 1 {

		c.JSON(http.StatusOK, gin.H{
			"data": "Professor não existente!",
		})

		return
	}

	c.JSON(http.StatusOK, professor)
}
