package email_services

import (
	"bytes"
	"encoding/json"
	"errors"
	"text/template"
	"user-service/api/models"

	"gorm.io/gorm"
)

// GetEmailTemplate retrieves an email template by code from the database.
func GetEmailTemplate(db *gorm.DB, templateCode string) (*models.EmailTemplate, error) {
	var template models.EmailTemplate
	if err := db.Where("code = ?", templateCode).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

// RenderTemplate renders the email template with the provided data.
func RenderTemplate(tmpl *models.EmailTemplate, data map[string]interface{}) (string, error) {
	// Unmarshal the Params field
	var params []string
	if err := json.Unmarshal([]byte(tmpl.Params), &params); err != nil {
		return "", err
	}

	// Validate if all required parameters are present in the data
	for _, param := range params {
		if _, ok := data[param]; !ok {
			return "", errors.New("missing required parameter: " + param)
		}
	}

	// Parse and render the template
	t, err := template.New(tmpl.Name).Parse(tmpl.Body)
	if err != nil {
		return "", err
	}

	var renderedBody bytes.Buffer
	if err := t.Execute(&renderedBody, data); err != nil {
		return "", err
	}

	return renderedBody.String(), nil
}

// SendEmail sends an email using the specified template and data.
func SendEmail(db *gorm.DB, templateCode string, data map[string]interface{}) error {
	// Get the email template
	tmpl, err := GetEmailTemplate(db, templateCode)
	if err != nil {
		return err
	}

	// Render the template with the provided data
	body, err := RenderTemplate(tmpl, data)
	if err != nil {
		return err
	}

	// Here you would add your email sending logic, e.g., using an SMTP server or an email API.
	// For demonstration purposes, we'll just print the rendered email body.
	println(body)

	return nil
}
