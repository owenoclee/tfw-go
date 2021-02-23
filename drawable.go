package tfw

import "github.com/owenoclee/tfw-go/geo"

type Drawable interface {
	Draw(Screen) KeyCallbacks
	SetBounds(geo.Rect)
	SetVisible(bool)
	Visible() bool
}

type MinBoundable interface {
	MinBounds(topLeft geo.Vector) geo.Rect
}

type HasChild interface {
	SetChild(child Drawable)
	Child() Drawable
}

type HasChildren interface {
	SetChildAt(index int, child Drawable)
	ChildAt(index int) Drawable
}

type HasText interface {
	SetText(text string)
	Text() string
}
