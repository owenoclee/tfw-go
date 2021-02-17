package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type HorizontalSplit struct {
	bounds   geo.Rect
	visible  bool
	Children []tfw.Drawable
}

func (hs *HorizontalSplit) Draw(s tfw.Screen) tfw.KeyCallbacks {
	n := len(hs.Children)
	fullHeight := hs.bounds.VerticalCells()
	splitHeight := fullHeight / n
	splitHeightRemainder := fullHeight % n

	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := hs.bounds.TopLeft
	height := splitHeight + splitHeightRemainder
	for _, child := range hs.Children {
		if !child.Visible() {
			continue
		}
		topRight := topLeftCursor.SetX(hs.bounds.BottomRight.X)
		bounds := geo.Rect{
			TopLeft: topLeftCursor,
			BottomRight: topRight.Add(geo.Vector{
				0,
				height - 1,
			}),
		}
		child.SetBounds(bounds)
		callbacks.Push(child.Draw(s))

		topLeftCursor = topLeftCursor.Add(geo.Vector{0, height})
		height = splitHeight
	}
	return callbacks
}

func (hs *HorizontalSplit) SetBounds(r geo.Rect) {
	hs.bounds = r
}

func (hs *HorizontalSplit) SetVisible(visible bool) {
	hs.visible = visible
}

func (hs *HorizontalSplit) Visible() bool {
	return hs.visible
}
