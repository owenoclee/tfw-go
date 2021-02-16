package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type VerticalSplit struct {
	bounds   geo.Rect
	Children []tfw.Drawable
}

func (vs *VerticalSplit) Draw(s tfw.Screen) tfw.KeyCallbacks {
	n := len(vs.Children)
	fullWidth := vs.bounds.HorizontalCells()
	splitWidth := fullWidth / n
	splitWidthRemainder := fullWidth % n

	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := vs.bounds.TopLeft
	width := splitWidth + splitWidthRemainder
	for _, c := range vs.Children {
		bottomLeft := topLeftCursor.SetY(vs.bounds.BottomRight.Y)
		bounds := geo.Rect{
			TopLeft: topLeftCursor,
			BottomRight: bottomLeft.Add(geo.Vector{
				width - 1,
				0,
			}),
		}
		c.SetBounds(bounds)
		callbacks.Push(c.Draw(s))

		topLeftCursor = topLeftCursor.Add(geo.Vector{width, 0})
		width = splitWidth
	}
	return callbacks
}

func (vs *VerticalSplit) SetBounds(r geo.Rect) {
	vs.bounds = r
}
