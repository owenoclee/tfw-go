package main

type vector struct {
	x int
	y int
}

type piece int

// pieces
const (
	TopLeftPiece piece = iota
	BottomLeftPiece
	TopRightPiece
	BottomRightPiece
	HorizontalPiece
	VerticalPiece
)

type cellMeta struct {
	pos  vector
	kind piece
}

func rectCells(topLeft vector, bottomRight vector) []cellMeta {
	if topLeft.x > bottomRight.x || topLeft.y > bottomRight.y {
		return nil
	}

	var metaCells []cellMeta
	for x := topLeft.x; x <= bottomRight.x; x++ {
		topCell := cellMeta{
			pos:  vector{x, topLeft.y},
			kind: HorizontalPiece,
		}
		bottomCell := cellMeta{
			pos:  vector{x, bottomRight.y},
			kind: HorizontalPiece,
		}
		switch x {
		case topLeft.x:
			topCell.kind = TopLeftPiece
			bottomCell.kind = BottomLeftPiece
		case bottomRight.x:
			topCell.kind = TopRightPiece
			bottomCell.kind = BottomRightPiece
		}
		metaCells = append(metaCells, topCell, bottomCell)
	}
	for y := topLeft.y + 1; y <= bottomRight.y-1; y++ {
		topCell := cellMeta{
			pos:  vector{topLeft.x, y},
			kind: VerticalPiece,
		}
		bottomCell := cellMeta{
			pos:  vector{bottomRight.x, y},
			kind: VerticalPiece,
		}
		metaCells = append(metaCells, topCell, bottomCell)
	}
	return metaCells
}
