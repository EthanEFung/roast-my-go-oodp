package main

import "fmt"

type exportJSONStats struct {
}
func (visitor *exportJSONStats) visitPikachu(element *pikachu) {
	/*
		The actual implementation of this should create some sort of json file
	*/
	fmt.Println("Exporting JSON")
	element.printStats()
}
func (visitor *exportJSONStats) visitCharmander(element *charmander) {
	fmt.Println("Exporting JSON")
	element.printStats()
}