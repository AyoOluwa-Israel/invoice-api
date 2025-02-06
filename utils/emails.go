package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/AyoOluwa-Israel/invoice-api/models"
	gomail "gopkg.in/gomail.v2"
)

func ConvertEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

type EmailData struct {
	Name    string
	Subject string
}




func SendEmail(message models.MessageStruct) error {

	// Set up the email parameters
	m := gomail.NewMessage()

	m.SetHeader("From", "AyoOluwa Israel <support@israelayooluwa.com>")
	m.SetHeader("To", message.Email)

	processedTemplate, err := template.ParseFiles("./templates/contact_us.html")

	if err != nil {
		log.Fatalf("Error Loading templates: %v", err)
	}

	data := EmailData{
		Name:    message.Name,
		Subject: "Thank You for Contacting Me!",
	}

	var body bytes.Buffer

	if err := processedTemplate.Execute(&body, data); err != nil {
		return fmt.Errorf("error rendering template: %v", err)
	}

	m.SetBody("text/html", body.String())
	m.SetHeader("Subject", data.Subject) 

	// Set up the SMTP server
	d := gomail.NewDialer("live.smtp.mailtrap.io", 2525, "api", "39d977e502de8560a4ada97189559b54")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	log.Println("Email sent successfully!")
	return nil // Return nil to indicate success

}
