package main

type classicPiece struct {
	piece
}

func (p *classicPiece) name() string {
	return p.piece.name
}
func (p *classicPiece) value() int {
	return p.piece.value
}
func (p *classicPiece) style() string {
	return "classic"
}
