package main

/*
	modernFactory is the concrete abstract factory of abstractChessFactory
*/
type modernFactory struct{}

func (af modernFactory) makePawn() pieceInterface {
	return &modernPiece{
		piece: piece{"pawn", 1},
	}
}
func (af modernFactory) makeQueen() pieceInterface {
	return &modernPiece{
		piece: piece{"queen", 9},
	}
}
