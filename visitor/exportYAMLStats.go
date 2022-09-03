package main

import "fmt"

/*
	exportYAMLStats is a visitor that contains methods unique to pokemon we can export from.
*/
type exportYAMLStats struct {
}

func (visitor *exportYAMLStats) visitPikachu(element *pikachu) {
	/*
		The actual implementation of this should create some sort of yaml file
	*/
	fmt.Println("Exporting YAML")
	element.printStats()
}
func (visitor *exportYAMLStats) visitCharmander(element *charmander) {
	/*
		The actual implementation of this should create some sort of yaml file
	*/
	fmt.Println("Exporting YAML")
	element.printStats()
}
