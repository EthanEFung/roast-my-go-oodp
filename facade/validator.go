package main

import (
	"fmt"
	"net/mail"
)

type validator interface {
	validate(values formValues) error 
}

type validatorInstance struct {}

func newValidator() validator {
	return validatorInstance{}
}

func (i validatorInstance) validate(values formValues) error {
	// if value.username == taken { error }
	if _, err := mail.ParseAddress(values.email); err != nil {
		return fmt.Errorf("invalid email address")
	}
	if len(values.password) < 12 {
		return fmt.Errorf("password is less than 12 characters")
	}

	return nil
}