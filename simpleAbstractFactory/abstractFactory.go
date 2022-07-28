package main

type abstractChessFactory interface {
	makePawn() piece
	makeQueen() piece
}
