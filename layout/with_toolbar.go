package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type withToolbar struct {
	bounds      geo.Rect
	visible     bool
	child       tfw.Drawable
	barElements []tfw.MinBoundableDrawable
	gap         int
}

func NewWithToolbar(gap int, child tfw.Drawable, barElements ...tfw.MinBoundableDrawable) *withToolbar {
	return &withToolbar{
		visible:     true,
		child:       child,
		barElements: barElements,
		gap:         gap,
	}
}

var _ tfw.Drawable = &withToolbar{}

func (wt *withToolbar) Draw(s tfw.Screen) tfw.KeyCallbacks {
	barBounds := geo.Rect{
		TopLeft:     wt.bounds.TopLeft.SetY(wt.bounds.BottomRight.Y),
		BottomRight: wt.bounds.BottomRight,
	}
	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := barBounds.TopLeft
	for _, elem := range wt.barElements {
		if !elem.Visible() {
			continue
		}
		bounds := elem.MinBounds(topLeftCursor)
		elem.SetBounds(bounds)
		callbacks.Push(elem.Draw(s))
		topLeftCursor = topLeftCursor.SetX(bounds.BottomRight.X + 1 + wt.gap)
	}

	childBounds := geo.Rect{
		TopLeft:     wt.bounds.TopLeft,
		BottomRight: wt.bounds.BottomRight.Add(geo.Vector{0, -1}),
	}
	wt.child.SetBounds(childBounds)
	callbacks.Push(wt.child.Draw(s))
	return callbacks
}

func (wt *withToolbar) SetBounds(b geo.Rect) {
	wt.bounds = b
}

func (wt *withToolbar) SetVisible(visible bool) {
	wt.visible = visible
}

func (wt *withToolbar) Visible() bool {
	return wt.visible
}
