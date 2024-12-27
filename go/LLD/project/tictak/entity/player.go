package entity

type Player struct {
	name  string
	piece Piece
}

func NewPlayer(name string, piece Piece) *Player {
	return &Player{name: name, piece: piece}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetPiece() Piece {
	return p.piece
}
