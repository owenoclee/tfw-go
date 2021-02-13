package drawable

import (
	"github.com/owenoclee/tfw-go/cli-client/canvas"
	"github.com/owenoclee/tfw-go/cli-client/geom"
)

type HorizontalSplit struct {
	bounds   geom.Rect
	Children []Drawable
}

func (hs *HorizontalSplit) Draw(s canvas.Screen) KeyCallbacks {
	n := len(hs.Children)
	if n == 0 {
		child := Blank{}
		child.SetBounds(hs.bounds)
		return child.Draw(s)
	}

	fullHeight := hs.bounds.VerticalCells()
	splitHeight := fullHeight / n
	splitHeightRemainder := fullHeight % n

	callbacks := NewKeyCallbacks()
	topLeft := hs.bounds.TopLeft
	height := splitHeight + splitHeightRemainder
	for _, c := range hs.Children {
		topRight := topLeft.SetX(hs.bounds.BottomRight.X)
		bounds := geom.Rect{
			TopLeft: topLeft,
			BottomRight: topRight.Add(geom.Vector{
				0,
				height - 1,
			}),
		}
		c.SetBounds(bounds)
		callbacks.Push(c.Draw(s))

		topLeft = topLeft.Add(geom.Vector{0, height})
		height = splitHeight
	}
	return callbacks
}

func (hs *HorizontalSplit) SetBounds(r geom.Rect) {
	hs.bounds = r
}
