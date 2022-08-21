package main

import "fmt"

func main() {
	/*
		imagine we have a game that comprises of two opposing teams.
	*/
	game := newGame()

	/*
		what we know about each team is that each player will either
		be dressed in red, or dressed in blue
	*/
	game.addRedPlayer("Offense")
	game.addRedPlayer("Offense")
	game.addRedPlayer("Offense")
	game.addRedPlayer("Offense")

	game.addBluePlayer("Defense")
	game.addBluePlayer("Defense")
	game.addBluePlayer("Defense")
	game.addBluePlayer("Defense")

	/*
		since we know that the two teams should be wearing the
		same thing, we have the opportunity to reduce the memory consumption
		of our game by having two dress objects (red and blue) and having all
		team players reference those objects as the attire they are wearing.
	*/
	dressFactoryInstance := getDressFactoryInstance()

	for _, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("dress type: %v\n", dress.getColor())
	}
}
