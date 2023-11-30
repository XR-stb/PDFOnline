package provider

import "fmt"

type std struct{}

func (s *std) Init() error {
	return nil
}

func (s *std) Send(to, subject, body string) error {
	fmt.Printf("Sending email to \"%s\" with subject \"%s\" and body \"%s\"\n", to, subject, body)
	return nil
}
