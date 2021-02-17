package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type Rows struct {
	bounds   geo.Rect
	visible  bool
	Children []tfw.Drawable
	RowLines int
	RowGap   int
}

func (r *Rows) Draw(s tfw.Screen) tfw.KeyCallbacks {
	if r.RowLines < 1 {
		panic("RowLines must be greater than 0")
	}

	s.ClearRegion(r.bounds)

	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := r.bounds.TopLeft
	for _, child := range r.Children {
		boundsOfChild := geo.Rect{
			TopLeft:     topLeftCursor,
			BottomRight: topLeftCursor.SetX(r.bounds.BottomRight.X).Add(geo.Vector{0, r.RowLines - 1}),
		}
		if !child.Visible() || !r.bounds.RectInBounds(boundsOfChild) {
			continue
		}
		child.SetBounds(boundsOfChild)
		callbacks.Push(child.Draw(s))

		topLeftCursor = topLeftCursor.Add(geo.Vector{0, r.RowLines + r.RowGap})
	}
	return callbacks
}

func (r *Rows) SetBounds(b geo.Rect) {
	r.bounds = b
}

func (r *Rows) SetVisible(visible bool) {
	r.visible = visible
}

func (r *Rows) Visible() bool {
	return r.visible
}
