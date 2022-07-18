package main

/** the product interface */
type card interface {
	Name() string
	Play(game *game, area string)
}