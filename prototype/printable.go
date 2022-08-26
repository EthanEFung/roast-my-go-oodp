package main

/*
	printable is an interface describing types have a print method
*/
type printable interface {
	print(indentation int) string
}
