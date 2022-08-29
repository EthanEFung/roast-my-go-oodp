package main

/*
	energyCard implements the `card` interface
*/
type energyCard struct {
	name string
}
func (c energyCard) Name() string {
	return c.name
}
func (c energyCard) Play(g *game, area string) {

}