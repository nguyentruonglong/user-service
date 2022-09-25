package auth

import (
	"net/http"
	users "user_service/versioned/v1/users"

	"github.com/gin-gonic/gin"
)

var userModel = new(users.UserModel)

func AuthValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"message": "Request does not contain an access token"})
			c.Abort()
			return
		}
		err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthGenerateJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtGenerateValidator := NewJWTGenerateValidator()
		err := jwtGenerateValidator.Bind(c)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		}

		validatedData := &jwtGenerateValidator.validatedData

		username := validatedData.Username
		password := validatedData.Password
		query, countInt, err := userModel.FindOneUser(&users.UserModel{Email: username})

		if err != nil || countInt == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Bad request"})
			return
		}

		isValidPassword, err := userModel.CheckPassword(password)
		if err != nil && isValidPassword {
			c.JSON(http.StatusNotFound, gin.H{"message": "Bad request"})
			return
		}

		token, err := GenerateJWT(query.Email, query.Password)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
		c.Next()
	}
}
