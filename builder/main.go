package main

import "fmt"

func main() {
	/*
	each builder will have methods that are exposed to manipulate
	the properties of the object.
	*/
	mageBuilder := newMageBuilder()
	knightBuilder := newKnightBuilder()

	/*
	instead of having a singular function to build a mage or a knight,
	the builder pattern delegates the responsibility of creating to a
	director object.
	*/
	director := newDirector(mageBuilder)

	/*
	the director builds the object and is resposibile for the
	proper configurations
	*/
	mage := director.buildCharacter()
	
	director.setBuilder(knightBuilder)

	knight := director.buildCharacter()

	fmt.Println("mage", mage)
	fmt.Println("knight", knight)
}