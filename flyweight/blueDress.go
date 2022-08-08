package main

type blueDress struct {}

func newBlueDress() *blueDress {
	return &blueDress{}
}

func (b *blueDress) getColor() string {
	return "blue"
}