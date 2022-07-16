package main

type knightBuilder struct {
	name string
	attack int
	defense int
	speed int
	special int
}

func newKnightBuilder() *knightBuilder {
	return &knightBuilder{}
}

func(b *knightBuilder) setName() {
	b.name = "knight"
} 

func(b *knightBuilder) setAttack() {
	b.attack = 8
}

func(b *knightBuilder) setDefense() {
	b.defense = 7
}

func(b *knightBuilder) setSpeed() {
	b.speed = 3
}

func(b *knightBuilder) setSpecial() {
	b.special = 3 
}

func(b *knightBuilder) getCharacter() character {
	return character{
		name: b.name,
		attack: b.attack,
		defense: b.defense,
		speed: b.speed,
		special: b.special,
	}
}