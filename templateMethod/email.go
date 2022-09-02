package main

import "fmt"

/*
	email in a traditional oodp language would be a subclass of otp. Here the email
	struct stores a reference to a base otp and abides by the
	`operations` interface
*/
type email struct {
	otp
}

func (s *email) getRandomPassword(len int) string {
	/*
		notice here that if we want to utilize otp's method, the developer must explicitly
		call `otp.getRandomPassword`. This is a design decision in go which prioritizes
		composition over inheritance.
	*/
	return s.otp.getRandomPassword(len)
}

func (s *email) savePassword(otp string) {
	s.otp.savePassword(otp)
}

func (s *email) getMessage(otp string) string {
	return s.otp.getMessage(otp)
}

func (s *email) sendNotification(message string) error {
	/*
		here however we alter the one of the step in the operation not to use otp's
		sendNotification but instead write out an operation that does something different
	*/
	fmt.Printf("Sending email: %s\n", message)
	return nil
}
