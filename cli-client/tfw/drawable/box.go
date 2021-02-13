package drawable

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type Box struct {
	bounds geo.Rect
	Pieces map[BoxPiece]rune
	Child  Drawable
}

var _ Drawable = &Box{}

func (b *Box) Draw(s tfw.Screen) KeyCallbacks {
	if !b.bounds.IsValid() {
		panic("invalid Box bounds")
	}
	pieces := b.Pieces
	if pieces == nil {
		pieces = DefaultBoxPieces
	}

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

	child := b.Child
	if child == nil {
		child = &Blank{}
	}
	child.SetBounds(b.bounds.Shrink(1))
	return child.Draw(s)
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
