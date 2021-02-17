package tfw

import "github.com/owenoclee/tfw-go/geo"

type Drawable interface {
	Draw(Screen) KeyCallbacks
	SetBounds(geo.Rect)
	SetVisible(bool)
	Visible() bool
}

type MinBoundableDrawable interface {
	Drawable
	MinBounds(topLeft geo.Vector) geo.Rect
}

type DrawableWithChild interface {
	Drawable
	SetChild(child Drawable)
	Child() Drawable
}

type DrawableWithChildren interface {
	Drawable
	SetChildAt(index int, child Drawable)
	ChildAt(index int) Drawable
}
