package main

/*
	device is the receiver interface.
	Typically a receiver object will have a series of methods for business logic.
*/
type device interface {
	on()
	off()
}
