package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type margin struct {
	bounds  geo.Rect
	visible bool
	child   tfw.Drawable
	top     int
	right   int
	left    int
	bottom  int
}

func NewMargin(top, bottom, left, right int, child tfw.Drawable) *margin {
	return &margin{
		visible: true,
		child:   child,
		top:     top,
		bottom:  bottom,
		left:    left,
		right:   right,
	}
}

var _ tfw.DrawableWithChild = &margin{}

func (m *margin) Draw(s tfw.Screen) tfw.KeyCallbacks {
	if !m.child.Visible() {
		return nil
	}
	m.child.SetBounds(geo.Rect{
		TopLeft:     m.bounds.TopLeft.Add(geo.Vector{m.left, m.top}),
		BottomRight: m.bounds.BottomRight.Add(geo.Vector{-m.right, -m.bottom}),
	})
	return m.child.Draw(s)
}

func (m *margin) SetBounds(b geo.Rect) {
	m.bounds = b
}

func (m *margin) SetVisible(visible bool) {
	m.visible = visible
}

func (m *margin) Visible() bool {
	return m.visible
}

func (m *margin) SetChild(child tfw.Drawable) {
	m.child = child
}

func (m *margin) Child() tfw.Drawable {
	return m.child
}
