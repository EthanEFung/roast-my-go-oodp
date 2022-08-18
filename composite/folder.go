package main

import (
	"fmt"
)

/*
	folder is a concrete composite
*/
type folder struct {
	name string
	/*
		what we find is that our folder struct also contains child composites
	*/
	children []composite
}

func (c folder) getName() string {
	return c.name
}

/*
	notice that our search method recursively calls the search method on all the
	children of the the folder
*/
func (c folder) search(str string) bool {
	var found bool
	fmt.Printf("searching for term \"%v\" in folder %v\n", str, c.name)
	for _, f := range c.children {
		here := f.search(str)
		if here {
			found = here
		}
	}
	return found
}
