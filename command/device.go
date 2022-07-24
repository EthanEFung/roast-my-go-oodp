package main

/*
device is the receiver interface
*/
type device interface {
	on() 
	off()
}