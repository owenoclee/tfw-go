package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type Box struct {
	bounds geo.Rect
	Pieces map[BoxPiece]rune
	Child  tfw.Drawable
}

var _ tfw.Drawable = &Box{}

func (b *Box) Draw(s tfw.Screen) tfw.KeyCallbacks {
	if !b.bounds.IsValid() {
		panic("invalid Box bounds")
	}
	pieces := b.Pieces
	if pieces == nil {
		pieces = DefaultBoxPieces
	}

	s.ClearRegion(b.bounds)

	for x := b.bounds.TopLeft.X; x <= b.bounds.BottomRight.X; x++ {
		topCellKind := HorizontalBoxPiece
		bottomCellKind := HorizontalBoxPiece
		switch x {
		case b.bounds.TopLeft.X:
			topCellKind = TopLeftBoxPiece
			bottomCellKind = BottomLeftBoxPiece
		case b.bounds.BottomRight.X:
			topCellKind = TopRightBoxPiece
			bottomCellKind = BottomRightBoxPiece
		}
		s.SetContent(geo.Vector{x, b.bounds.TopLeft.Y}, pieces[topCellKind], tcell.StyleDefault)
		s.SetContent(geo.Vector{x, b.bounds.BottomRight.Y}, pieces[bottomCellKind], tcell.StyleDefault)
	}
	for y := b.bounds.TopLeft.Y + 1; y <= b.bounds.BottomRight.Y-1; y++ {
		s.SetContent(geo.Vector{b.bounds.TopLeft.X, y}, pieces[VerticalBoxPiece], tcell.StyleDefault)
		s.SetContent(geo.Vector{b.bounds.BottomRight.X, y}, pieces[VerticalBoxPiece], tcell.StyleDefault)
	}

	if b.Child == nil {
		return nil
	}
	b.Child.SetBounds(b.bounds.Shrink(1))
	return b.Child.Draw(s)
}

func (b *Box) SetBounds(r geo.Rect) {
	b.bounds = r
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
