package main

import "fmt"

func main() {
	game := newGame()

	game.addRedPlayer("Offense")
	game.addRedPlayer("Offense")
	game.addRedPlayer("Offense")
	game.addRedPlayer("Offense")

	game.addBluePlayer("Defense")
	game.addBluePlayer("Defense")
	game.addBluePlayer("Defense")
	game.addBluePlayer("Defense")

	dressFactoryInstance := getDressFactoryInstance()

	for _, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("dress type: %v\n", dress.getColor())
	}
}
