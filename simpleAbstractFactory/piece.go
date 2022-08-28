package main

type piece struct {
	name  string
	value int
}

/*
	pieceInterface is the interface of the piece object, and the return type of our
	abstract factory methods
*/
type pieceInterface interface {
	name() string
	value() int
	style() string
}
