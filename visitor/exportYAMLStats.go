package main

import "fmt"

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