package main

/*
	composite is the interface that is shared between all
	branch and leaf objects of a the pattern's tree
*/
type composite interface {
	search(string) bool
	getName() string
}
