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
	/*
		imagine we have a catalog of different chess piece styles for our application, and
		would like to make sure that these collections remain in tact (no mixing).
	*/
	classic := &classicFactory{}
	modern := &modernFactory{}

	/*
		by creating differnt factories that abide by the same interface, we can produce
		the same products without risk of the products falling outside of the family
		of objects.
	*/
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
