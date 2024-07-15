package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	templateID := os.Getenv("TEMPLATE_ID")
	fromEmail := os.Getenv("FROM_EMAIL")
	csvFile := os.Getenv("CSV_FILE")
	emailSubject := os.Getenv("EMAIL_SUBJECT")
	emailContent := os.Getenv("EMAIL_CONTENT")

	if sendgridAPIKey == "" || templateID == "" || fromEmail == "" || csvFile == "" || emailSubject == "" || emailContent == "" {
		log.Fatalf("Environment variables not set correctly")
	}

	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		log.Fatalf("Failed to read CSV headers: %s", err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		dynamicData := make(map[string]string)
		for i, header := range headers {
			dynamicData[header] = record[i]
		}

		toEmail := dynamicData["email"]
		if err := sendEmail(sendgridAPIKey, templateID, fromEmail, toEmail, dynamicData, emailSubject, emailContent); err != nil {
			log.Printf("Failed to send email to %s: %s", toEmail, err)
		}
	}
}

func sendEmail(apiKey, templateID, fromEmail, toEmail string, dynamicData map[string]string, subject, content string) error {
	from := mail.NewEmail("Example User", fromEmail)
	to := mail.NewEmail("Recipient", toEmail)
	message := mail.NewSingleEmail(from, subject, to, content, "")
	message.SetTemplateID(templateID)

	for key, value := range dynamicData {
		message.Personalizations[0].SetDynamicTemplateData(key, value)
	}

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	// Log status code and full response body
	log.Printf("Email sent to %s with status code %d. Response: %+v\n", toEmail, response.StatusCode, response)

	return nil
}
