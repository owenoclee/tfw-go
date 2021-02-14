package layout

import (
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type Columns struct {
	bounds      geo.Rect
	Children    []tfw.Drawable
	ColumnCells int
	ColumnGap   int
}

func (c *Columns) Draw(s tfw.Screen) tfw.KeyCallbacks {
	if c.ColumnCells < 1 {
		panic("ColumnCells must be greater than 0")
	}

	s.ClearRegion(c.bounds)

	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := c.bounds.TopLeft
	for _, child := range c.Children {
		boundsOfChild := geo.Rect{
			TopLeft:     topLeftCursor,
			BottomRight: topLeftCursor.SetY(c.bounds.BottomRight.Y).Add(geo.Vector{c.ColumnCells - 1, 0}),
		}
		if !c.bounds.RectInBounds(boundsOfChild) {
			break
		}
		child.SetBounds(boundsOfChild)
		callbacks.Push(child.Draw(s))

		topLeftCursor = topLeftCursor.Add(geo.Vector{c.ColumnCells + c.ColumnGap, 0})
	}
	return callbacks
}

func (c *Columns) SetBounds(b geo.Rect) {
	c.bounds = b
}
