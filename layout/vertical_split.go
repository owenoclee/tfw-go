package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

var _ tfw.Drawable = (*verticalSplit)(nil)
var _ tfw.HasChildren = (*verticalSplit)(nil)

type verticalSplit struct {
	bounds   geo.Rect
	visible  bool
	children []tfw.Drawable
}

func NewVerticalSplit(children ...tfw.Drawable) *verticalSplit {
	return &verticalSplit{
		visible:  true,
		children: children,
	}
}

func (vs *verticalSplit) Draw(s tfw.Screen) tfw.KeyCallbacks {
	n := len(vs.children)
	fullWidth := vs.bounds.HorizontalCells()
	splitWidth := fullWidth / n
	splitWidthRemainder := fullWidth % n

	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := vs.bounds.TopLeft
	width := splitWidth + splitWidthRemainder
	for _, child := range vs.children {
		if !child.Visible() {
			continue
		}
		bottomLeft := topLeftCursor.SetY(vs.bounds.BottomRight.Y)
		bounds := geo.Rect{
			TopLeft: topLeftCursor,
			BottomRight: bottomLeft.Add(geo.Vector{
				width - 1,
				0,
			}),
		}
		child.SetBounds(bounds)
		callbacks.Push(child.Draw(s))

		topLeftCursor = topLeftCursor.Add(geo.Vector{width, 0})
		width = splitWidth
	}
	return callbacks
}

func (vs *verticalSplit) SetBounds(r geo.Rect) {
	vs.bounds = r
}

func (vs *verticalSplit) SetVisible(visible bool) {
	vs.visible = visible
}

func (vs *verticalSplit) Visible() bool {
	return vs.visible
}

func (vs *verticalSplit) AppendChild(child tfw.Drawable) {
	vs.children = append(vs.children, child)
}

func (vs *verticalSplit) PrependChild(child tfw.Drawable) {
	vs.children = append([]tfw.Drawable{child}, vs.children...)
}

func (vs *verticalSplit) SetChildAt(index int, child tfw.Drawable) {
	vs.children[index] = child
}

func (vs *verticalSplit) ChildAt(index int) tfw.Drawable {
	if len(vs.children) > index {
		return vs.children[index]
	}
	return nil
}

func (vs *verticalSplit) ChildrenLen() int {
	return len(vs.children)
}
