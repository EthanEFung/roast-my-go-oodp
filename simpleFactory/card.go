package main

/*
	card is the product interface produced by concrete cardFactories
*/
type card interface {
	Name() string
	Play(game *game, area string)
}