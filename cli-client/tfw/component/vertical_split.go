package component

import (
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type VerticalSplit struct {
	bounds   geo.Rect
	Children []tfw.Drawable
}

func (vs *VerticalSplit) Draw(s tfw.Screen) tfw.KeyCallbacks {
	n := len(vs.Children)
	if n == 0 {
		child := Blank{}
		child.SetBounds(vs.bounds)
		return child.Draw(s)
	}

	fullWidth := vs.bounds.HorizontalCells()
	splitWidth := fullWidth / n
	splitWidthRemainder := fullWidth % n

	callbacks := tfw.NewKeyCallbacks()
	topLeft := vs.bounds.TopLeft
	width := splitWidth + splitWidthRemainder
	for _, c := range vs.Children {
		bottomLeft := topLeft.SetY(vs.bounds.BottomRight.Y)
		bounds := geo.Rect{
			TopLeft: topLeft,
			BottomRight: bottomLeft.Add(geo.Vector{
				width - 1,
				0,
			}),
		}
		c.SetBounds(bounds)
		callbacks.Push(c.Draw(s))

		topLeft = topLeft.Add(geo.Vector{width, 0})
		width = splitWidth
	}
	return callbacks
}

func (vs *VerticalSplit) SetBounds(r geo.Rect) {
	vs.bounds = r
}
