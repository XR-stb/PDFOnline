package provider

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"sync"

	"github.com/sirupsen/logrus"

	"backend/pkg/config"
)

var template = "To: %s\r\nSubject: %s\r\n\r\n%s\r\n"

type SMTP struct {
	client *smtp.Client
	mu     sync.Mutex
}

func (s *SMTP) Init() error {
	smtpConfig := config.SMTPConfig()

	c, err := smtp.Dial(fmt.Sprintf("%s:%d", smtpConfig.Host, smtpConfig.Port))
	if err != nil {
		return fmt.Errorf("smtp dial failed, error: %v", err)
	}

	if ok, _ := c.Extension("STARTTLS"); ok {
		if err := c.StartTLS(&tls.Config{
			ServerName: smtpConfig.Host,
		}); err != nil {
			return fmt.Errorf("starttls failed, error: %v", err)
		}
	}

	if err = c.Auth(smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)); err != nil {
		logrus.Warn("smtp plain auth failed, try login auth")
		err = c.Auth(LoginAuth(smtpConfig.Username, smtpConfig.Password, smtpConfig.Host))
		if err != nil {
			return fmt.Errorf("smtp auth failed, error: %v", err)
		}
	}

	err = c.Mail(smtpConfig.From)
	if err != nil {
		return fmt.Errorf("smtp mail failed, error: %v", err)
	}

	s.client = c

	return nil
}

func (s *SMTP) send(to, subject, body string) error {
	err := s.client.Rcpt(to)
	if err != nil {
		return err
	}

	wc, err := s.client.Data()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(wc, template, to, subject, body)
	if err != nil {
		return err
	}

	err = wc.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *SMTP) Send(to, subject, body string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.send(to, subject, body)
	if err != nil {
		logrus.Warnf("smtp send failed, error: %v, reconnect and try again", err)
		err = s.Init()
		if err != nil {
			return err
		}
		err = s.send(to, subject, body)
		if err != nil {
			return err
		}
	}

	return nil
}

type loginAuth struct {
	username, password string
	host               string
}

func LoginAuth(username, password, host string) smtp.Auth {
	return &loginAuth{username, password, host}
}

func (a loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	return "LOGIN", nil, nil
}

func (a loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		if bytes.EqualFold([]byte("username:"), fromServer) {
			return []byte(a.username), nil
		} else if bytes.EqualFold([]byte("password:"), fromServer) {
			return []byte(a.password), nil
		}
	}
	return nil, nil
}
