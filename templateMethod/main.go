package main

/*
  What makes the template method awkward in go is that it primarily relies on
	inheritance and Go does not have this support. This means that the concrete
	implementation would still require the user to explicitly detail the methods of the
	shared interface.
*/
func main() {
	base := otp{} // here we create the base struct that is shared among the algorithms

	// here we pass the base as a reference to smsOneTimePassword
	smsOneTimePassword := &sms{base}

	o := otp{
		/*
			then we add it again to o (which will be used in the application)
		*/
		operations: smsOneTimePassword,
	}
	o.createAndSendPassword(4)

	// here we repeat the 
	emailOneTimePassword := &email{base}
	o = otp{
		operations: emailOneTimePassword,
	}
	o.createAndSendPassword(4)
}
