package main

/*
	blueDress is a concrete flyweight that is consumed by the `player` struct
*/
type blueDress struct{}

func newBlueDress() *blueDress {
	return &blueDress{}
}

func (b *blueDress) getColor() string {
	return "blue"
}
