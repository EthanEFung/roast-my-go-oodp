package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		imagine that our clients would like to interact with a form.
		instead of dealing directly with a form struct
		the client is limited to its facade.
	*/
	f := newFormFacade()
	values := formValues{
		username: "adamsandlerfan",
		email:    "test@test.com",
		password: "123456789012",
		dob:      "05/20/2001",
	}
	fmt.Printf("submitting...\n")
	err := f.submit(values) // using the facade pattern all the client needs is to use the submit method
	if err != nil {
		log.Fatal(fmt.Errorf("submission error %v", err))
	}
	fmt.Printf("success\n")
	/*
		this pattern seems simple on the outside, but if we dive a little deeper
		all the complexity that is surrounding state management,
		asynchronous calls, and form validations is abstracted away.
	*/
}
