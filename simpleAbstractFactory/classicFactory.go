package main

type classicFactory struct{}

func (af classicFactory) makePawn() pieceInterface {
	return &classicPiece{
		piece: piece{"pawn", 1},
	}
}
func (af classicFactory) makeQueen() pieceInterface {
	return &classicPiece{
		piece: piece{"queen", 9},
	}
}
