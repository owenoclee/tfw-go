package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

var _ tfw.Drawable = (*rows)(nil)
var _ tfw.HasChildren = (*rows)(nil)

type rows struct {
	bounds   geo.Rect
	visible  bool
	children []tfw.Drawable
	lines    int
	gap      int
}

func NewRows(lines, gap int, children ...tfw.Drawable) *rows {
	return &rows{
		visible:  true,
		children: children,
		lines:    lines,
		gap:      gap,
	}
}

func (r *rows) Draw(s tfw.Screen) tfw.KeyCallbacks {
	if r.lines < 1 {
		panic("RowLines must be greater than 0")
	}

	s.ClearRegion(r.bounds)

	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := r.bounds.TopLeft
	for _, child := range r.children {
		boundsOfChild := geo.Rect{
			TopLeft:     topLeftCursor,
			BottomRight: topLeftCursor.SetX(r.bounds.BottomRight.X).Add(geo.Vector{0, r.lines - 1}),
		}
		if !child.Visible() || !r.bounds.RectInBounds(boundsOfChild) {
			continue
		}
		child.SetBounds(boundsOfChild)
		callbacks.Push(child.Draw(s))

		topLeftCursor = topLeftCursor.Add(geo.Vector{0, r.lines + r.gap})
	}
	return callbacks
}

func (r *rows) SetBounds(b geo.Rect) {
	r.bounds = b
}

func (r *rows) SetVisible(visible bool) {
	r.visible = visible
}

func (r *rows) Visible() bool {
	return r.visible
}

func (r *rows) AppendChild(child tfw.Drawable) {
	r.children = append(r.children, child)
}

func (r *rows) PrependChild(child tfw.Drawable) {
	r.children = append([]tfw.Drawable{child}, r.children...)
}

func (r *rows) SetChildAt(index int, child tfw.Drawable) {
	r.children[index] = child
}

func (r *rows) ChildAt(index int) tfw.Drawable {
	if len(r.children) > index {
		return r.children[index]
	}
	return nil
}

func (r *rows) ChildrenLen() int {
	return len(r.children)
}
