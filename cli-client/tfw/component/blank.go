package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type Blank struct {
	bounds geo.Rect
}

var _ tfw.Drawable = &Blank{}

func (b *Blank) Draw(s tfw.Screen) tfw.KeyCallbacks {
	for _, cell := range b.bounds.CellLocations() {
		s.SetContent(cell, ' ', tcell.StyleDefault)
	}
	return nil
}

func (b *Blank) SetBounds(r geo.Rect) {
	b.bounds = r
}