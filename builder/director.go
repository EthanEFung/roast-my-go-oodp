package main

type director struct {
	builder builder
}

func newDirector(b builder) *director {
	return &director{b}
}

func(d *director) setBuilder(b builder) {
	d.builder = b
}

func(d *director) buildCharacter() character {
	d.builder.setName()
	d.builder.setAttack()
	d.builder.setDefense()
	d.builder.setSpeed()
	d.builder.setSpecial()
	return d.builder.getCharacter()
}