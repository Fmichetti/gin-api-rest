package middlewares

import (
	"fmt"
	"net/http"

	"github.com/FMichetti/api-go-gin/config"
	"github.com/FMichetti/api-go-gin/models"
	"github.com/FMichetti/api-go-gin/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(selected_role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		user_id, err := token.ExtractTokenID(c)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		userRole, err := getUserRoleFromDatabase(user_id)

		fmt.Println(userRole)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if userRole == selected_role {
			c.Next()

			return
		}

		c.String(http.StatusUnauthorized, "Unauthorized")

	}
}

// getUserRoleFromDatabase obtém a função do usuário com base no ID do usuário
func getUserRoleFromDatabase(userID uint) (string, error) {
	user, err := getUserByID(userID)
	if err != nil {
		return "", err
	}

	// A função do usuário é obtida a partir do campo UserType
	return user.UserType, nil
}

func getUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
