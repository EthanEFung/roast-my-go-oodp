package main

/*
	printablePrototype is a mixed interface of types that can print and clone
*/
type printablePrototype interface {
	printable
	prototype[printablePrototype]
}
