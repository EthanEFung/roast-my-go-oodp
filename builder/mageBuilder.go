package main

type mageBuilder struct {
	name string
	attack int
	defense int
	speed int
	special int
}

func newMageBuilder() *mageBuilder {
	return &mageBuilder{}
}

func(b *mageBuilder) setName() {
	b.name = "mage"
} 

func(b *mageBuilder) setAttack() {
	b.attack = 3
}

func(b *mageBuilder) setDefense() {
	b.defense = 2
}

func(b *mageBuilder) setSpeed() {
	b.speed = 4
}

func(b *mageBuilder) setSpecial() {
	b.special = 10
}

func(b *mageBuilder) getCharacter() character {
	return character{
		name: b.name,
		attack: b.attack,
		defense: b.defense,
		speed: b.speed,
		special: b.special,
	}
}