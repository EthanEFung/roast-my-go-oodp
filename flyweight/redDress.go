package main

type redDress struct {
}

func newRedDress() *redDress {
	return &redDress{}
}

func (d *redDress) getColor() string {
	return "red"
}
