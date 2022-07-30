package main

import "fmt"


type pikachu struct {
	stats
}

func (element *pikachu) getType() string {
	return "electric"
}
func (element *pikachu) printStats() {
	fmt.Println("--PIKACHU--")
	fmt.Println("Attack:", element.stats.attack)
	fmt.Println("Defense:", element.stats.defense)
	fmt.Println("Speed:", element.stats.speed)
	fmt.Println("Special:", element.stats.special)
}
func (element *pikachu) accept(v visitor) {
	v.visitPikachu(element)
}

