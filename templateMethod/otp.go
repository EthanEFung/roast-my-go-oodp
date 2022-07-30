package main

import "fmt"

type operations interface {
	getRandomPassword(int) string
	savePassword(string)
	getMessage(string) string
	sendNotification(string) error
}

type otp struct {
	operations operations
}

func (o *otp) createAndSendPassword(pwLength int) error {
	otp := o.operations.getRandomPassword(pwLength)
	o.operations.savePassword(otp)
	message := o.operations.getMessage(otp)
	err := o.operations.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

func (o *otp) getRandomPassword(len int) string {
	randomOTP := "1234"
	fmt.Println("Generating password", randomOTP)
	return randomOTP
}

func (o *otp) savePassword(otp string) {
	fmt.Println("Saving password to the cache")
	for _, char := range otp {
		fmt.Print(byte(char))
	}
	fmt.Println()
}

func (o *otp) getMessage(otp string) string {
	return "Password is: " + otp
}
