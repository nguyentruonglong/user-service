package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(UserModel)

func (u UserController) GetUser(c *gin.Context) {
	id_int, err := strconv.Atoi(c.Param("id"))

	query, err := userModel.FindOneUser(&UserModel{ID: id_int})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Bad request"})
		return
	}
	serializer := UserSerializer{c, query}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func (u UserController) GetUsers(c *gin.Context) {
	offset_int, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		offset_int = 0
	}

	limit_int, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		limit_int = 50
	}

	query, _, err := userModel.FindManyUser(limit_int, offset_int)

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

	if err := SaveOne(&userCreateValidator.validatedData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	serializer := UserSerializer{c, userCreateValidator.validatedData}
	c.JSON(http.StatusCreated, serializer.Response())
}
