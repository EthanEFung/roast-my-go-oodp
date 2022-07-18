package main

/* implements the card product interface */
type monsterCard struct {
	name string
}

func (c monsterCard) Name() string {
	return c.name
}
func (c monsterCard) Play(g *game, area string) {

}