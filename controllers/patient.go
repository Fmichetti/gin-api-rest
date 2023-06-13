package controllers

import (
	"net/http"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Description Exibe todos os pacientes
// @Tags Patient
// @Produce json
// @Success 200 {object} []models.Patient
// @Router /patient [get]
func ExibeTodosPacientes(c *gin.Context) {
	var patient []models.Patient

	config.DB.Find(&patient)

	c.JSON(http.StatusOK, patient)
}

// @Description Cadastra um novo paciente
// @Tags Patient
// @Produce json
// @Param patient body models.Patient true "Patient object"
// @Success 200 {object} []models.Patient
// @Router /patient [post]
func CriaNovoPaciente(c *gin.Context) {
	var patient models.Patient

	var err error

	if err = c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Generate a new UUID
	patient.ID = uuid.New()

	err = config.DB.Create(&patient).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, patient)
}

// @Description Realiza uma busca baseada pelo ID
// @Tags Patient
// @Produce json
// @Success 200 {object} models.Patient
// @Router /patient/{id} [get]
func BuscaPacientePorID(c *gin.Context) {
	var patient models.Patient
	var err error

	id := c.Params.ByName("id")

	uuidToFind, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Paciente não encontrado",
		})

		return
	}

	err = config.DB.First(&patient, "id = ?", uuidToFind).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Paciente não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, patient)
}

// @Description Realiza uma edição baseada pelo ID
// @Tags Patient
// @Produce json
// @Param patient body models.Patient true "Patient object"
// @Success 200
// @Router /patient/{id} [patch]
func EditaPaciente(c *gin.Context) {
	var patient models.Patient
	id := c.Params.ByName("id")
	db := config.DB.First(&patient, id)

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	config.DB.Model(&patient).UpdateColumns(patient)

	//verifica se foi editado algum paciente
	if db.RowsAffected < 1 {

		c.JSON(http.StatusOK, gin.H{
			"data": "Paciente não encontrado!",
		})

		return
	}

	c.JSON(http.StatusOK, patient)
}

// @Description Realiza uma busca baseada pelo CPF
// @Tags Patient
// @Produce json
// @Success 200 {object} models.Patient
// @Router /patient/{cpf} [get]
func BuscaPacientePorCPF(c *gin.Context) {
	var patient models.Patient
	cpf := c.Params.ByName("cpf")
	config.DB.Where(&models.Patient{CPF: cpf}).First(&patient)

	if patient.CPF == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Paciente não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, patient)
}

// @Description Realiza uma busca baseada pelo ID
// @Tags Patient
// @Produce json
// @Success 200
// @Router /patient/{id} [delete]
func DeletaPacientePorID(c *gin.Context) {
	var patient models.Patient
	id := c.Params.ByName("id")
	db := config.DB.Delete(&patient, id)

	//verifica se foi removido algum paciente
	if db.RowsAffected < 1 {

		c.JSON(http.StatusOK, gin.H{
			"data": "Paciente não existente!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Paciente deletado com sucesso!",
	})
}
