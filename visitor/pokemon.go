package main

/*
pokemon is the element interface of the visitor pattern
*/
type pokemon interface {
	getType() string
	accept(visitor)
	printStats()
}