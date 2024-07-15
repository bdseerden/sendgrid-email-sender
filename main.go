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

	if sendgridAPIKey == "" || templateID == "" || fromEmail == "" || csvFile == "" {
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
		if err := sendEmail(sendgridAPIKey, templateID, fromEmail, toEmail, dynamicData); err != nil {
			log.Printf("Failed to send email to %s: %s", toEmail, err)
		}
	}
}

func sendEmail(apiKey, templateID, fromEmail, toEmail string, dynamicData map[string]string) error {
	from := mail.NewEmail("Example User", fromEmail)
	to := mail.NewEmail(toEmail, toEmail)
	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(templateID)

	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	for key, value := range dynamicData {
		personalization.SetDynamicTemplateData(key, value)
	}
	message.AddPersonalizations(personalization)

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	// Log status code and full response body
	log.Printf("Email sent to %s with status code %d. Response: %+v\n", toEmail, response.StatusCode, response)

	return nil
}
