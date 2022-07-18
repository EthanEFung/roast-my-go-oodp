package main
type cardFactory interface {
	create(name string) card
}
