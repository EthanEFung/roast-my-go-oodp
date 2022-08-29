package main

import "fmt"

func main() {
	/*
		in our example, we are constructing a deck of pokemon cards that need at minimum
		monster cards and energy cards. So we focus on having a factory for each type of
		card, that only needs to pass in the name / title of the card in order to create
		these cards for our deck.
	*/
	mFactory := monsterFactory{}
	eFactory := energyFactory{}
	pikachu := mFactory.create("pikachu")
	electric := eFactory.create("electric")
	fmt.Println("Cards", pikachu, electric)
}