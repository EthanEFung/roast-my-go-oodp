package main

/*
	player is the concrete context object that consumes the flyweight `dress`
	notice that this object comprises of immutable data (aka intrinsic data) from
	the flyweight dress, and mutable data such as playerType and lat long which will
	change for the player over the course of their session.
*/
type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := dressFactoryInstance.getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
