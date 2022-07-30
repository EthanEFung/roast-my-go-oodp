package main

/*
  What makes the template method awkward in go
	is that it primarily relies on inheritance and
	Go does not have this support. This means that
	the concrete implementation would still require
	the user to explicitly detail the methods of
	the shared interface
*/
func main() {
	base := otp{}
	smsOneTimePassword := &sms{base}
	o := otp{
		operations: smsOneTimePassword,
	}
	o.createAndSendPassword(4)

	emailOneTimePassword := &email{base}
	o = otp{
		operations: emailOneTimePassword,
	}
	o.createAndSendPassword(4)
}
