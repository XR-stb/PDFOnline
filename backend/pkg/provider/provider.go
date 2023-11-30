package provider

import "github.com/sirupsen/logrus"

type Interface interface {
	Init() error
	Send(to, subject, body string) error
}

var Default Interface = new(std)

func Init() error {
	smtp := new(SMTP)
	err := smtp.Init()
	if err != nil {
		logrus.Warnf("SMTP provider initialize failed, error: %v", err)
	} else {
		Default = smtp
	}

	return nil
}

func Send(to, subject, body string) error {
	return Default.Send(to, subject, body)
}
