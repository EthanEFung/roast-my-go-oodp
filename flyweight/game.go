package main

type game struct {
	redTeam  []*player
	blueTeam []*player
}

func newGame() *game {
	return &game{
		redTeam:  make([]*player, 1),
		blueTeam: make([]*player, 1),
	}
}

func (g *game) addRedPlayer(playerType string) {
	player := newPlayer(playerType, RedDressType)
	g.redTeam = append(g.redTeam, player)
}

func (g *game) addBluePlayer(playerType string) {
	player := newPlayer(playerType, BlueDressType)
	g.blueTeam = append(g.blueTeam, player)
}
