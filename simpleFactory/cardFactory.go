package main

/*
	cardFactory is the factory interface with a method that produces products with the
	`card` interface
*/
type cardFactory interface {
	create(name string) card
}
