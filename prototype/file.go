package main

import "strings"

/*
	file is a concrete printablePrototype
*/
type file struct {
	name string
}

func (concrete *file) clone() printablePrototype {
	return &file{name: concrete.name + "_clone"}
}
func (concrete file) print(indentation int) string {
	prefix := strings.Repeat(" ", indentation*4)
	return prefix + concrete.name + "\n"
}
