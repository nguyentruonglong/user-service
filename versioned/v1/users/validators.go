package users

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserCreateValidator struct {
	userCreateSerializer UserCreateSerializer
	validatedData        UserModel `json:"-"`
}

func NewUserCreateValidator() UserCreateValidator {
	return UserCreateValidator{}
}

func (s *UserCreateValidator) Bind(c *gin.Context) error {
	if err := c.ShouldBindJSON(&s.userCreateSerializer); err != nil {
		return err
	}

	s.validatedData.Email = s.userCreateSerializer.Email
	s.validatedData.Password = s.userCreateSerializer.Password
	s.validatedData.Name = s.userCreateSerializer.Name
	s.validatedData.Phone = s.userCreateSerializer.Phone
	s.validatedData.Address = s.userCreateSerializer.Address
	s.validatedData.Country = s.userCreateSerializer.Country
	s.validatedData.Zipcode = s.userCreateSerializer.Zipcode
	s.validatedData.CreatedAt = time.Now()
	s.validatedData.UpdatedAt = time.Now()

	return nil
}
