package auth

import (
	"fmt"
	"net/http"
	"strings"
	"user_service/versioned/v1/users"

	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

var userModel = new(users.UserModel)

func AuthValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		fmt.Println("Request ID:", uuid.String())
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}
