package main

/*
	redDress is a concrete flyweight that is consumed by the `player` struct
*/
type redDress struct {
}

func newRedDress() *redDress {
	return &redDress{}
}

func (d *redDress) getColor() string {
	return "red"
}
