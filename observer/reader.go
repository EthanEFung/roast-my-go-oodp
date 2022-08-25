package main

import "fmt"

/*
  reader is a concrete subscriber which implements an update function
  accessed by publishers
*/
type reader struct {
	name string
}

func (sub *reader) update(a *article) {
	fmt.Printf("%s here's a new article:\n%s\n%s\n\n", sub.name, a.name, a.content)
}
func newReader(name string) *reader {
	return &reader{name}
}
