package main

import (
	"fmt"
)

type folder struct {
	name  string
	files []composite
}

func (c folder) getName() string {
	return c.name
}
func (c folder) search(str string) bool {
	var found bool
	fmt.Printf("searching for term \"%v\" in folder %v\n", str, c.name)
	for _, f := range c.files {
		here := f.search(str)
		if here {
			found = here
		}
	}
	return found
}
