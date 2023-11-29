package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/FMichetti/api-go-gin/utils/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegistraAluno(c *gin.Context) {
	var input struct {
		Username  string `json:"username" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
		UserType  string `json:"user_type" binding:"required"`
		AlunoData struct {
			Nome           string    `json:"nome"`
			Matricula      string    `json:"matricula"`
			DataNascimento time.Time `json:"data_nascimento"`
			TurmaID        uint      `json:"turma_id"`
		} `json:"aluno_data"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.UserType = input.UserType

	userId, err := CriaUsuario(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aluno := models.Aluno{
		UserID:         userId,
		Nome:           input.AlunoData.Nome,
		Matricula:      input.AlunoData.Matricula,
		DataNascimento: input.AlunoData.DataNascimento,
		TurmaID:        input.AlunoData.TurmaID,
	}

	err = config.DB.Create(&aluno).Error

	if err != nil {
		status := DeletaUserPorID(userId)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": status})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aluno registrado com sucesso!", "id": aluno.ID})
}

func CriaUsuario(user models.User) (uint, error) {
	var err error

	u := user

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}

	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = strings.ReplaceAll(user.Username, " ", "")

	err = config.DB.Create(&u).Error

	if err != nil {
		return 0, err
	}

	return u.ID, nil
}

func Login(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	u.UserType = input.UserType

	token, err := LoginCheck(u.Username, u.Password, u.UserType)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func LoginCheck(username string, password string, user_type string) (string, error) {

	var err error

	u := models.User{}

	err = config.DB.Model(models.User{}).Where("username = ? AND user_type = ?", username, user_type).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func DeletaUserPorID(id uint) string {
	var user models.User

	db := config.DB.Delete(&user, id)

	//verifica se foi removido algum aluno
	if db.RowsAffected < 1 {
		return "Erro ao remover usuario"
	}

	return "OK"
}
