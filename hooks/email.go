package hooks

import (
	"net/smtp"
	"log"

	"github.com/EkoAgustina/go-ms-portfolio/utils"
)

// SendEmail sends an email using SMTP with the specified parameters.
// It constructs the email message from the provided recipient, subject, and body.
// The sender's email and password are loaded from environment variables.
//
// Parameters:
// - to: The recipient's email address.
// - subject: The subject line of the email.
// - body: The body content of the email.
//
// On success, it logs a message indicating the email was sent. On failure, it logs the error encountered.
func SendEmail(to string, subject string, body string) {
    from := utils.LoadEnv("EMAIL_FROM")
    pass := utils.LoadEnv("EMAIL_PASSWORD")

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    log.Printf("Sending email to %s with subject: %s", to, subject)

    err := smtp.SendMail("smtp.gmail.com:587",
        smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
        from, []string{to}, []byte(msg))

    if err != nil {
        log.Printf("smtp error: %s while sending to %s", err, to)
        return
    }
    log.Println("Successfully sent to " + to)
}
