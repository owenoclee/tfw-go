package drawable

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/canvas"
	"github.com/owenoclee/tfw-go/cli-client/geom"
)

type Blank struct {
	bounds geom.Rect
}

var _ Drawable = &Blank{}

func (b *Blank) Draw(s canvas.Screen) KeyCallbacks {
	for _, cell := range b.bounds.CellLocations() {
		s.SetContent(cell, ' ', tcell.StyleDefault)
	}
	return nil
}

func (b *Blank) SetBounds(r geom.Rect) {
	b.bounds = r
}
