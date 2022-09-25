package auth

import (
	"github.com/gin-gonic/gin"
)

type JWTGenerateValidator struct {
	jwtGenerateSerializer JWTGenerateSerializer
	validatedData         JWTClaim `json:"-"`
}

func NewJWTGenerateValidator() JWTGenerateValidator {
	return JWTGenerateValidator{}
}

func (s *JWTGenerateValidator) Bind(c *gin.Context) error {
	if err := c.ShouldBindJSON(&s.jwtGenerateSerializer); err != nil {
		return err
	}

	s.validatedData.Username = s.jwtGenerateSerializer.Username
	s.validatedData.Password = s.jwtGenerateSerializer.Password

	return nil
}
