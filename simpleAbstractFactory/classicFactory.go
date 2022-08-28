package main

/*
	classicFactory is the concrete abstract factory of abstractChessFactory
*/
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
