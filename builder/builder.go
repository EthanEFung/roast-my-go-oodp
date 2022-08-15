package main

/*
	builders have a number of methods that will be exposed to a director
*/
type builder interface {
	setName()
	setAttack()
	setDefense()
	setSpeed()
	setSpecial()
	getCharacter() character
}