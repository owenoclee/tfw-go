package layout

import (
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

var _ tfw.Drawable = (*horizontalSplit)(nil)
var _ tfw.HasChildren = (*horizontalSplit)(nil)

type horizontalSplit struct {
	bounds   geo.Rect
	visible  bool
	children []tfw.Drawable
}

func NewHorizontalSplit(children ...tfw.Drawable) *horizontalSplit {
	return &horizontalSplit{
		visible:  true,
		children: children,
	}
}

func (hs *horizontalSplit) Draw(s tfw.Screen) tfw.KeyCallbacks {
	n := len(hs.children)
	fullHeight := hs.bounds.VerticalCells()
	splitHeight := fullHeight / n
	splitHeightRemainder := fullHeight % n

	callbacks := tfw.NewKeyCallbacks()
	topLeftCursor := hs.bounds.TopLeft
	height := splitHeight + splitHeightRemainder
	for _, child := range hs.children {
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

func (hs *horizontalSplit) SetBounds(r geo.Rect) {
	hs.bounds = r
}

func (hs *horizontalSplit) SetVisible(visible bool) {
	hs.visible = visible
}

func (hs *horizontalSplit) Visible() bool {
	return hs.visible
}

func (hs *horizontalSplit) AppendChild(child tfw.Drawable) {
	hs.children = append(hs.children, child)
}

func (hs *horizontalSplit) PrependChild(child tfw.Drawable) {
	hs.children = append([]tfw.Drawable{child}, hs.children...)
}

func (hs *horizontalSplit) SetChildAt(index int, child tfw.Drawable) {
	hs.children[index] = child
}

func (hs *horizontalSplit) ChildAt(index int) tfw.Drawable {
	if len(hs.children) > index {
		return hs.children[index]
	}
	return nil
}

func (hs *horizontalSplit) ChildrenLen() int {
	return len(hs.children)
}
