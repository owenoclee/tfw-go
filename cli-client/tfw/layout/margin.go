package layout

import (
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type Margin struct {
	bounds geo.Rect
	Child  tfw.Drawable
	Top    int
	Left   int
	Right  int
	Bottom int
}

func (m *Margin) Draw(s tfw.Screen) tfw.KeyCallbacks {
	m.Child.SetBounds(geo.Rect{
		TopLeft:     m.bounds.TopLeft.Add(geo.Vector{m.Left, m.Top}),
		BottomRight: m.bounds.BottomRight.Add(geo.Vector{-m.Right, -m.Bottom}),
	})
	return m.Child.Draw(s)
}

func (m *Margin) SetBounds(b geo.Rect) {
	m.bounds = b
}
