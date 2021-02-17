package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type Margin struct {
	bounds  geo.Rect
	visible bool
	Child   tfw.Drawable
	Top     int
	Left    int
	Right   int
	Bottom  int
}

func (m *Margin) Draw(s tfw.Screen) tfw.KeyCallbacks {
	if !m.Child.Visible() {
		return nil
	}
	m.Child.SetBounds(geo.Rect{
		TopLeft:     m.bounds.TopLeft.Add(geo.Vector{m.Left, m.Top}),
		BottomRight: m.bounds.BottomRight.Add(geo.Vector{-m.Right, -m.Bottom}),
	})
	return m.Child.Draw(s)
}

func (m *Margin) SetBounds(b geo.Rect) {
	m.bounds = b
}

func (m *Margin) SetVisible(visible bool) {
	m.visible = visible
}

func (m *Margin) Visible() bool {
	return m.visible
}
