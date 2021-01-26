package utils

import (
	"crypto/tls"
	"log"
	"workspace/practice/tcp-dialer/model"

	"gopkg.in/gomail.v2"
)

// SendEmail ...
func SendEmail(to string, subject string, content string, header string, smtp model.SMTP) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", smtp.Sender)
	mailer.SetHeader("To", to)
	// mailer.SetAddressHeader("Cc", "test@xapiens.id", "Test")
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", content)
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		smtp.Host,
		smtp.Port,
		smtp.Sender,
		smtp.Password,
	)

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Println("SendEmail error :", err)
		// panic(err)
	}
}
