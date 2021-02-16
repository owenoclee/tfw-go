package tfw

import "github.com/owenoclee/tfw-go/geo"

type Drawable interface {
	Draw(Screen) KeyCallbacks
	SetBounds(geo.Rect)
}

type MinBoundableDrawable interface {
	Drawable
	MinBounds(topLeft geo.Vector) geo.Rect
}