package main

import "fmt"

func main() {
	mFactory := monsterFactory{}
	eFactory := energyFactory{}
	pikachu := mFactory.create("pikachu")
	electric := eFactory.create("electric")
	fmt.Println("Cards", pikachu, electric)
}