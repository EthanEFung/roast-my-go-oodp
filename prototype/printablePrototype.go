package main
type printablePrototype interface {
	printable
	prototype[printablePrototype]
}