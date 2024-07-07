package email_services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/smtp"
	"text/template"
	"user-service/api/models"
	"user-service/config"

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

// SendEmail sends an email using the specified template, data, and recipient.
func SendEmail(db *gorm.DB, templateCode string, recipient string, data map[string]interface{}, cfg *config.AppConfig) error {
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

	// Determine the email provider and send the email
	switch cfg.EmailConfig.Provider {
	case "mailjet":
		return sendMailjetEmail(cfg, tmpl.Subject, recipient, body)
	case "sendgrid":
		return sendSendgridEmail(cfg, tmpl.Subject, recipient, body)
	default:
		return sendGenericEmail(cfg, tmpl.Subject, recipient, body)
	}
}

// sendMailjetEmail sends an email using Mailjet.
func sendMailjetEmail(cfg *config.AppConfig, subject, recipient, body string) error {
	mailjetConfig := cfg.EmailConfig.Mailjet
	return sendSMTPEmail(mailjetConfig.SMTPServer, mailjetConfig.SMTPPort, mailjetConfig.SMTPUser, mailjetConfig.SMTPPassword, mailjetConfig.SenderEmail, subject, recipient, body)
}

// sendSendgridEmail sends an email using Sendgrid.
func sendSendgridEmail(cfg *config.AppConfig, subject, recipient, body string) error {
	sendgridConfig := cfg.EmailConfig.Sendgrid
	return sendSMTPEmail(sendgridConfig.SMTPServer, sendgridConfig.SMTPPort, sendgridConfig.SMTPUser, sendgridConfig.SMTPPassword, sendgridConfig.SenderEmail, subject, recipient, body)
}

// sendGenericEmail sends an email using a generic SMTP server.
func sendGenericEmail(cfg *config.AppConfig, subject, recipient, body string) error {
	genericConfig := cfg.EmailConfig.Generic
	return sendSMTPEmail(genericConfig.SMTPServer, genericConfig.SMTPPort, genericConfig.SMTPUser, genericConfig.SMTPPassword, genericConfig.SenderEmail, subject, recipient, body)
}

// sendSMTPEmail is a helper function to send an email using an SMTP server.
func sendSMTPEmail(smtpServer string, smtpPort int, smtpUser, smtpPass, senderEmail, subject, recipient, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpServer)

	// Prepare email
	from := senderEmail
	to := recipient
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", from, to, subject, body)

	// Send email
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	println("Email sent to:", recipient)
	println("Body:", body)

	return nil
}
