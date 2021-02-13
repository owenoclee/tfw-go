package drawable

import (
	"github.com/owenoclee/tfw-go/cli-client/canvas"
	"github.com/owenoclee/tfw-go/cli-client/geom"
)

type VerticalSplit struct {
	bounds   geom.Rect
	Children []Drawable
}

func (vs *VerticalSplit) Draw(s canvas.Screen) KeyCallbacks {
	n := len(vs.Children)
	if n == 0 {
		child := Blank{}
		child.SetBounds(vs.bounds)
		return child.Draw(s)
	}

	fullWidth := vs.bounds.HorizontalCells()
	splitWidth := fullWidth / n
	splitWidthRemainder := fullWidth % n

	callbacks := NewKeyCallbacks()
	topLeft := vs.bounds.TopLeft
	width := splitWidth + splitWidthRemainder
	for _, c := range vs.Children {
		bottomLeft := topLeft.SetY(vs.bounds.BottomRight.Y)
		bounds := geom.Rect{
			TopLeft: topLeft,
			BottomRight: bottomLeft.Add(geom.Vector{
				width - 1,
				0,
			}),
		}
		c.SetBounds(bounds)
		callbacks.Push(c.Draw(s))

		topLeft = topLeft.Add(geom.Vector{width, 0})
		width = splitWidth
	}
	return callbacks
}

func (vs *VerticalSplit) SetBounds(r geom.Rect) {
	vs.bounds = r
}
