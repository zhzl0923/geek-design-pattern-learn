package flyweight

const (
	Red   = "red"
	Black = "black"
)

var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "車",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
	// ... 其他棋子
}

type ChessPieceUnit struct {
	ID    int
	Name  string
	Color string
}

func NewChessPieceUnit(id int, name, color string) *ChessPieceUnit {
	return &ChessPieceUnit{ID: id, Name: name, Color: color}
}

type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

func NewChessPiece(unit *ChessPieceUnit, x, y int) *ChessPiece {
	return &ChessPiece{Unit: unit, X: x, Y: y}
}

type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

// NewChessBoard 初始化棋盘
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for _, unit := range units {
		board.chessPieces[unit.ID] = &ChessPiece{
			Unit: NewChessPieceUnit(unit.ID, unit.Name, unit.Color),
			X:    0,
			Y:    0,
		}
	}
	return board
}

// Move 移动棋子
func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}
