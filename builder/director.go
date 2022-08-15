package main

/*
	director is an object that is reponsible for creating objects
*/
type director struct {
	builder builder
}

func newDirector(b builder) *director {
	return &director{b}
}

/*
	the director should have a method that allows the director to change
	the a builder object
*/
func(d *director) setBuilder(b builder) {
	d.builder = b
}

/*
	generally the builder would have the ability to change builders instructions 
	to match the use case
*/
func(d *director) buildCharacter() character {
	d.builder.setName()
	d.builder.setAttack()
	d.builder.setDefense()
	d.builder.setSpeed()
	d.builder.setSpecial()
	return d.builder.getCharacter()
}