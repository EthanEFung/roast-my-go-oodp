package main

type modernPiece struct {
	piece
}

func (p *modernPiece) name() string {
	return p.piece.name
}
func (p *modernPiece) value() int {
	return p.piece.value
}
func (p *modernPiece) style() string {
	return "modern"
}
