package helper

import "gopkg.in/gomail.v2"

func SendEmail(recipient string, subject string, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "mohit84cseb22@bpitindia.edu.in")
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer("smtp.office365.com", 587, "mohit84cseb22@bpitindia.edu.in", "Abhishek@1998")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
