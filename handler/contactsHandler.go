package handler

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/ZafirProjects/portfolio/view/contact"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type ContactHandler struct{}

type FormData struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (h *ContactHandler) HandleContactRender(c echo.Context) error {
	return render(c, contact.ContactRender())
}

func (h *ContactHandler) HandleSendEmail(c echo.Context) error {
	errs := godotenv.Load()
	if errs != nil {
		fmt.Println("Error loading .env file")
	}

	// Sender data.
	from := "real.zafirmon@gmail.com"
	password := os.Getenv("PASSWORD")

	// Receiver email address.
	to := []string{
		"real.zafirmon@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	var data FormData
	if err := c.Bind(&data); err != nil {
		return err
	}
	fmt.Println(data.Email)
	fmt.Println(data.Email)
	message := []byte("Subject: Email from " + data.Email + " from your Portfolio website!\r\n" + data.Message)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return render(c, contact.RenderForm())
}
