package engine

/*
0-5: from square
6-11: to square
12-13: promotion piece (knight, bishop, rook, queen)
14-15: castling (01), en passant (10), promotion (11)
*/
type Move uint16

const (
	NoneFlag = iota << 14
	CastlingFlag
	EnPassantFlag
	PromotionFlag
)

func NewMove(from Square, to Square) Move {
	return Move(uint16(from) | uint16(to)<<6)
}

func NewPromotionMove(from Square, to Square, promotion uint8) Move {
	return Move(uint16(from) | uint16(to)<<6 | uint16(promotion-1)<<12 | PromotionFlag)
}

func NewMoveCastle(side uint8) Move {
	// TODO
	switch (side) {
	case WhiteKingside:
	case WhiteQueenside:
	case BlackKingside:
	case BlackQueenside:
	}
	return Move(0)
}

// ex: c2c1q
// ex: b2b4
func NewMoveFromStr(moveStr string) Move {
	from := NewSquareFromStr(moveStr[0:2])
	to := NewSquareFromStr(moveStr[2:4])

	if len(moveStr) == 5 {
		promotion := CharToPiece[moveStr[4]]
		return NewPromotionMove(from, to, promotion)
	} else {
		return NewMove(from, to)
	}
}

func (m Move) From() Square {
	return Square(m & 0b111111)
}

func (m Move) To() Square {
	return Square(m >> 6 & 0b111111)
}

func (m Move) Promotion() uint8 {
	return uint8(m>>12&0b11 + 1)
}

func (m Move) IsPromotion() bool {
	return m&PromotionFlag == PromotionFlag
}

func (m Move) IsCastling() bool {
	return m&CastlingFlag == CastlingFlag
}

func (m Move) IsEnPassant() bool {
	return m&EnPassantFlag == EnPassantFlag
}

func (m Move) Type() uint8 {
	return uint8(m & PromotionFlag)
}

func (m Move) ToString() string {
	if m.IsPromotion() {
		return m.From().ToString() + m.To().ToString() + string(PieceToChar[m.Promotion()])
	} else {
		return m.From().ToString() + m.To().ToString()
	}
}
