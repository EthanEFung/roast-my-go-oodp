package main

type energyCard struct {
	// implements the card interface
	name string
}
func (c energyCard) Name() string {
	return c.name
}
func (c energyCard) Play(g *game, area string) {

}