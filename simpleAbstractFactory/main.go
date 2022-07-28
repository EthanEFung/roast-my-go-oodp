/*
	The abstract factory can be a useful pattern if
	1. you have existing factory methods in your code base
	2. the factory methods produce products that can be categorized
	3. the categories are generally independent of each other
*/
package main

import "fmt"

func info(p pieceInterface) {
	fmt.Printf("name: %v\n", p.name())
	fmt.Printf("value: %v\n", p.value())
	fmt.Printf("style: %v\n", p.style())
}

func main() {
	classic := &classicFactory{}
	modern := &modernFactory{}

	classicPawn := classic.makePawn()
	classicQueen := classic.makeQueen()

	modernPawn := modern.makePawn()
	modernQueen := modern.makeQueen()

	fmt.Println()
	fmt.Println("classics")
	info(classicPawn)
	info(classicQueen)

	fmt.Println()
	fmt.Println("moderns")
	info(modernPawn)
	info(modernQueen)
}
