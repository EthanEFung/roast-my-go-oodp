package main

import (
	"fmt"
	"log"
)


func main() {
	f := newFormFacade() 
	values := formValues{
		username: "adamsandlerfan",
		email: "test@test.com",
		password: "123456789012",
		dob: "05/20/2001",
	}
	fmt.Printf("submitting...\n")
  if err := f.submit(values); err != nil {
		log.Fatal(fmt.Errorf("submission error %v", err))
	}
	fmt.Printf("success\n")
}