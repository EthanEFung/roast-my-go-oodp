package main

import "fmt"

type email struct {
	otp
}

func (s *email) getRandomPassword(len int) string {
	return s.otp.getRandomPassword(len)
}

func (s *email) savePassword(otp string) {
	s.otp.savePassword(otp)
}

func (s *email) getMessage(otp string) string {
	return s.otp.getMessage(otp)
}

func (s *email) sendNotification(message string) error {
	fmt.Printf("Sending email: %s\n", message)
	return nil
}
