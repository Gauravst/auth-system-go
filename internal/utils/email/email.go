package email

import "net/smtp"

func SendEmail(username, password, host, from, port string, toList []string, body []byte) error {
	auth := smtp.PlainAuth("", username, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		return err
	}

	return nil
}
