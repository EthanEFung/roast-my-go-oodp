package main

import "fmt"

type sms struct {
	otp
}

func (s *sms) getRandomPassword(len int) string {
	return s.otp.getRandomPassword(len)
}

func (s *sms) savePassword(otp string) {
	s.otp.savePassword(otp)
}

func (s *sms) getMessage(otp string) string {
	return s.otp.getMessage(otp)
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("Sending sms: %s\n", message)
	return nil
}
