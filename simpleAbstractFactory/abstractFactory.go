package main

/*
	abstractChessFactory is the interface that produces the chess pieces of a themed set
*/
type abstractChessFactory interface {
	makePawn() piece
	makeQueen() piece
}
