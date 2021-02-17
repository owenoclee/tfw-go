package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type columns struct {
	bounds   geo.Rect
	visible  bool
	children []tfw.Drawable
	cells    int
	gap      int
}

func NewColumns(cells, gap int, children ...tfw.Drawable) *columns {
	return &columns{
		visible:  true,
		children: children,
		cells:    cells,
		gap:      gap,
	}
}

var _ tfw.DrawableWithChildren = &columns{}

func (c *columns) Draw(s tfw.Screen) tfw.KeyCallbacks {
	if c.cells < 1 {
		panic("cells must be greater than 0")
	}

	s.ClearRegion(c.bounds)

	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := c.bounds.TopLeft
	for _, child := range c.children {
		boundsOfChild := geo.Rect{
			TopLeft:     topLeftCursor,
			BottomRight: topLeftCursor.SetY(c.bounds.BottomRight.Y).Add(geo.Vector{c.cells - 1, 0}),
		}
		if !child.Visible() || !c.bounds.RectInBounds(boundsOfChild) {
			continue
		}
		child.SetBounds(boundsOfChild)
		callbacks.Push(child.Draw(s))

		topLeftCursor = topLeftCursor.Add(geo.Vector{c.cells + c.gap, 0})
	}
	return callbacks
}

func (c *columns) SetBounds(b geo.Rect) {
	c.bounds = b
}

func (c *columns) SetVisible(visible bool) {
	c.visible = visible
}

func (c *columns) Visible() bool {
	return c.visible
}

func (c *columns) SetChildAt(index int, child tfw.Drawable) {
	c.children[index] = child
}

func (c *columns) ChildAt(index int) tfw.Drawable {
	if len(c.children) > index {
		return c.children[index]
	}
	return nil
}
