package main

type piece struct {
	name  string
	value int
}

type pieceInterface interface {
	name() string
	value() int
	style() string
}
