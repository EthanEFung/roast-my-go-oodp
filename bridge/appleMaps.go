package main

import "fmt"

/*
appleMaps is a concrete "implementation" of the app interface
*/
type appleMaps struct{}

func (app *appleMaps) autocomplete(substr string) {
	fmt.Printf("Apple Maps searching for %v\n", substr)
}
