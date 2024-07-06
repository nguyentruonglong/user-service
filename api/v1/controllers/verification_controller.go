package controllers

import (
	"net/http"
	"user-service/api/errors"
	"user-service/api/models"
	"user-service/config"
	"user-service/tasks"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SendEmailVerificationCode handles the email verification code request
func SendEmailVerificationCode(c *gin.Context, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	var input models.EmailVerificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrInvalidEmailVerificationInput)
		return
	}

	task := tasks.EmailTask{
		TemplateCode: "EMAIL_VERIFICATION",
		Data: map[string]interface{}{
			"FirstName":        input.FirstName,
			"VerificationLink": input.VerificationLink,
		},
		Recipient: input.Email,
	}

	if err := tasks.PublishEmailTask("email_queue", task, cfg); err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrEmailTaskPublishingFailed)
		return
	}

	response := models.EmailVerificationResponse{
		Message: "Verification email sent",
	}

	c.JSON(http.StatusOK, response)
}
