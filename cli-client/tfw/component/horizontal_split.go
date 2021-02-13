package component

import (
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type HorizontalSplit struct {
	bounds   geo.Rect
	Children []tfw.Drawable
}

func (hs *HorizontalSplit) Draw(s tfw.Screen) tfw.KeyCallbacks {
	n := len(hs.Children)
	if n == 0 {
		child := Blank{}
		child.SetBounds(hs.bounds)
		return child.Draw(s)
	}

	fullHeight := hs.bounds.VerticalCells()
	splitHeight := fullHeight / n
	splitHeightRemainder := fullHeight % n

	callbacks := tfw.NewKeyCallbacks()
	topLeft := hs.bounds.TopLeft
	height := splitHeight + splitHeightRemainder
	for _, c := range hs.Children {
		topRight := topLeft.SetX(hs.bounds.BottomRight.X)
		bounds := geo.Rect{
			TopLeft: topLeft,
			BottomRight: topRight.Add(geo.Vector{
				0,
				height - 1,
			}),
		}
		c.SetBounds(bounds)
		callbacks.Push(c.Draw(s))

		topLeft = topLeft.Add(geo.Vector{0, height})
		height = splitHeight
	}
	return callbacks
}

func (hs *HorizontalSplit) SetBounds(r geo.Rect) {
	hs.bounds = r
}
