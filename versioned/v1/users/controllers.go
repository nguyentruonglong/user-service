package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(UserModel)

func (u UserController) GetUser(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))

	query, _, err := userModel.FindOneUser(&UserModel{ID: idInt})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Bad request"})
		return
	}
	serializer := UserSerializer{c, query}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func (u UserController) GetUsers(c *gin.Context) {
	offsetInt, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		offsetInt = 0
	}

	limitInt, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		limitInt = 50
	}

	query, _, err := userModel.FindManyUser(limitInt, offsetInt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Bad request"})
		return
	}

	serializer := UsersSerializer{c, query}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func (u UserController) CreateUser(c *gin.Context) {
	userCreateValidator := NewUserCreateValidator()
	if err := userCreateValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	userModel.HashPassword(&userCreateValidator.validatedData)
	if err := userModel.SaveOne(&userCreateValidator.validatedData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	serializer := UserSerializer{c, userCreateValidator.validatedData}
	c.JSON(http.StatusCreated, serializer.Response())
}
