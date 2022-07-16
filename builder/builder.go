package main

type builder interface {
	setName()
	setAttack()
	setDefense()
	setSpeed()
	setSpecial()
	getCharacter() character
}