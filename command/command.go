package main

/*
	command is the interface that the senders will utilize to issue
	calls to a receiver.
*/
type command interface {
	name() string
	/*
		each concrete command should have a command method that will be used
		by invokers. 
	*/
	execute()
}
