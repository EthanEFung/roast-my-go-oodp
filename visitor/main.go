/*
	The visitor pattern can be helpful if your existing code base has
	a lot of structs that needs new functionality with as minimal changes
	to those structs as possible. This pattern can be awkward in go because
	it is much easier to implement if the language uses method overloading.
	The lack there of in go was a design decision and we can imagine that
	truly implementing a visitor in go, can become tedious as the visitor
	interface and concretes would constantly need to be changed so that
	it could visit structs of different types.
*/
package main

func main() {
	p := &pikachu{stats{
		attack: 9,
		defense: 4,
		speed: 15,
		special: 9,
	}}
	c := &charmander{stats{
		attack: 11,
		defense: 5,
		speed: 7,
		special: 6,
	}}

	jsonExporter := &exportJSONStats{}
	p.accept(jsonExporter)
	c.accept(jsonExporter)

	yamlExporter := &exportYAMLStats{}
	p.accept(yamlExporter)
	c.accept(yamlExporter)
}