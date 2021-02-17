package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type box struct {
	bounds       geo.Rect
	visible      bool
	borderPieces map[BoxPiece]rune
	child        tfw.Drawable
}

func NewBox(borderPieces map[BoxPiece]rune, child tfw.Drawable) *box {
	if borderPieces == nil {
		panic("borderPieces cannot be nil")
	}
	return &box{
		visible:      true,
		borderPieces: borderPieces,
		child:        child,
	}
}

var _ tfw.DrawableWithChild = &box{}

func (b *box) Draw(s tfw.Screen) tfw.KeyCallbacks {
	if !b.bounds.IsValid() {
		panic("invalid Box bounds")
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
		s.SetContent(geo.Vector{x, b.bounds.TopLeft.Y}, b.borderPieces[topCellKind], tcell.StyleDefault)
		s.SetContent(geo.Vector{x, b.bounds.BottomRight.Y}, b.borderPieces[bottomCellKind], tcell.StyleDefault)
	}
	for y := b.bounds.TopLeft.Y + 1; y <= b.bounds.BottomRight.Y-1; y++ {
		s.SetContent(geo.Vector{b.bounds.TopLeft.X, y}, b.borderPieces[VerticalBoxPiece], tcell.StyleDefault)
		s.SetContent(geo.Vector{b.bounds.BottomRight.X, y}, b.borderPieces[VerticalBoxPiece], tcell.StyleDefault)
	}

	if b.child == nil || !b.child.Visible() {
		return nil
	}
	b.child.SetBounds(b.bounds.Shrink(1))
	return b.child.Draw(s)
}

func (b *box) SetBounds(r geo.Rect) {
	b.bounds = r
}

func (b *box) SetVisible(visible bool) {
	b.visible = visible
}

func (b *box) Visible() bool {
	return b.visible
}

func (b *box) SetChild(child tfw.Drawable) {
	b.child = child
}

func (b *box) Child() tfw.Drawable {
	return b.child
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

var PrettyBorder = map[BoxPiece]rune{
	TopLeftBoxPiece:     '┌',
	BottomLeftBoxPiece:  '└',
	TopRightBoxPiece:    '┐',
	BottomRightBoxPiece: '┘',
	HorizontalBoxPiece:  '─',
	VerticalBoxPiece:    '│',
}
