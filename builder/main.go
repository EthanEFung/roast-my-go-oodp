package main

import "fmt"

func main() {
	mageBuilder := newMageBuilder()
	knightBuilder := newKnightBuilder()

	director := newDirector(mageBuilder)

	mage := director.buildCharacter()
	
	director.setBuilder(knightBuilder)

	knight := director.buildCharacter()

	fmt.Println("mage", mage)
	fmt.Println("knight", knight)
}