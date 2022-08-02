package main

import "strings"

type folder struct {
	name string
	children []printablePrototype
}
func(concrete *folder) clone() printablePrototype {
	clone := &folder{
		name: concrete.name + "_clone",
	}
	for _, proto := range concrete.children {
		clone.children = append(clone.children, proto.clone())
	}
	return clone
}
func(concrete *folder) print(indentation int) string {
	var result string
	prefix := strings.Repeat(" ", indentation)
	result += prefix + concrete.name + "\n" 
	for _, proto := range concrete.children {
		result += prefix
		result += proto.print(indentation + indentation)
	}
	return result
}