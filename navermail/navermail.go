package navermail

import "gopkg.in/gomail.v2"

func Send(
	fromUsername string,
	fromPassword string,
	to string,
	subject string,
	body string,
) error {
	m := gomail.NewMessage()
	m.SetHeader("From", fromUsername+"@naver.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.naver.com", 465, fromUsername, fromPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
