package main

import "fmt"

/*
gMaps is a concrete "implementation" of the app interface
*/
type gmaps struct{}

func (m *gmaps) autocomplete(substr string) {
	fmt.Printf("Google searching for %v\n", substr)
}
