package drawable

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/canvas"
	"github.com/owenoclee/tfw-go/cli-client/geom"
)

type Box struct {
	Bounds geom.Rect
	Pieces map[BoxPiece]rune
}

var _ Drawable = &Box{}

func (b *Box) Draw(s canvas.Screen) KeyCallbacks {
	if !b.Bounds.IsValid() {
		panic("invalid Box bounds")
	}

	for x := b.Bounds.TopLeft.X; x <= b.Bounds.BottomRight.X; x++ {
		topCellKind := HorizontalBoxPiece
		bottomCellKind := HorizontalBoxPiece
		switch x {
		case b.Bounds.TopLeft.X:
			topCellKind = TopLeftBoxPiece
			bottomCellKind = BottomLeftBoxPiece
		case b.Bounds.BottomRight.X:
			topCellKind = TopRightBoxPiece
			bottomCellKind = BottomRightBoxPiece
		}
		s.SetContent(geom.Vector{x, b.Bounds.TopLeft.Y}, b.Pieces[topCellKind], tcell.StyleDefault)
		s.SetContent(geom.Vector{x, b.Bounds.BottomRight.Y}, b.Pieces[bottomCellKind], tcell.StyleDefault)
	}
	for y := b.Bounds.TopLeft.Y + 1; y <= b.Bounds.BottomRight.Y-1; y++ {
		s.SetContent(geom.Vector{b.Bounds.TopLeft.X, y}, b.Pieces[VerticalBoxPiece], tcell.StyleDefault)
		s.SetContent(geom.Vector{b.Bounds.BottomRight.X, y}, b.Pieces[VerticalBoxPiece], tcell.StyleDefault)
	}

	return nil
}

type BoxPiece int

const (
	TopLeftBoxPiece BoxPiece = iota
	BottomLeftBoxPiece
	TopRightBoxPiece
	BottomRightBoxPiece
	HorizontalBoxPiece
	VerticalBoxPiece
)

var DefaultBoxPieces = map[BoxPiece]rune{
	TopLeftBoxPiece:     '┌',
	BottomLeftBoxPiece:  '└',
	TopRightBoxPiece:    '┐',
	BottomRightBoxPiece: '┘',
	HorizontalBoxPiece:  '─',
	VerticalBoxPiece:    '│',
}
