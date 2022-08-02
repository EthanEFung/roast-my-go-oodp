package main

import "strings"

type file struct {
	name string
}
func (concrete *file) clone() printablePrototype {
	return &file{name: concrete.name + "_clone"}
}
func (concrete file) print(indentation int) string {
	prefix := strings.Repeat(" ", indentation)
	return prefix + concrete.name + "\n"
}