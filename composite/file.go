package main

import (
	"fmt"
	"strings"
)

type file struct {
	name string
	content string
}
func (c file) getName() string {
	return c.name
}
func (c file) search(str string) bool {
	index := strings.Index(c.content, str)
	if index == -1 {
		return false
	}
	fmt.Printf("term \"%v\" was found in file %v\n", str, c.name)
	return true
}