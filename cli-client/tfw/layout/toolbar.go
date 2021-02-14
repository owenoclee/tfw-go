package layout

import (
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type BarElem struct {
	Shortcut rune
	Text     string
	Callback func()
}

type WithToolbar struct {
	bounds      geo.Rect
	Primary     tfw.Drawable
	BarElements []tfw.MinBoundableDrawable
	ElementGap  int
}

func (wt *WithToolbar) Draw(s tfw.Screen) tfw.KeyCallbacks {
	barBounds := geo.Rect{
		TopLeft:     wt.bounds.TopLeft.SetY(wt.bounds.BottomRight.Y),
		BottomRight: wt.bounds.BottomRight,
	}
	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := barBounds.TopLeft
	for _, elem := range wt.BarElements {
		bounds := elem.MinBounds(topLeftCursor)
		elem.SetBounds(bounds)
		callbacks.Push(elem.Draw(s))
		topLeftCursor = topLeftCursor.SetX(bounds.BottomRight.X + 1 + wt.ElementGap)
	}

	primaryBounds := geo.Rect{
		TopLeft:     wt.bounds.TopLeft,
		BottomRight: wt.bounds.BottomRight.Add(geo.Vector{0, -1}),
	}
	wt.Primary.SetBounds(primaryBounds)
	callbacks.Push(wt.Primary.Draw(s))
	return callbacks
}

func (wt *WithToolbar) SetBounds(b geo.Rect) {
	wt.bounds = b
}
