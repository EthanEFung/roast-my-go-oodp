package main

import "fmt"

type charmander struct {
	stats
}
func (element *charmander) getType() string {
	return "fire"
}
func (element *charmander) printStats() {
	fmt.Println("--CHARMANDER--")
	fmt.Println("Attack:", element.stats.attack)
	fmt.Println("Defense:", element.stats.defense)
	fmt.Println("Speed:", element.stats.speed)
	fmt.Println("Special:", element.stats.special)
}
func (element *charmander) accept(v visitor) {
	v.visitCharmander(element)
}
