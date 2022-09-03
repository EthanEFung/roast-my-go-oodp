package main

/*
	pokemon is the element interface of the visitor pattern. Without method overloading
	in go, each element in the vistor pattern will require an `accept` method that allows
	a visitor to operate on the element.
*/
type pokemon interface {
	getType() string
	accept(visitor)
	printStats()
}
