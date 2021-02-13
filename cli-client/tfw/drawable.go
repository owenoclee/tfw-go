package tfw

import "github.com/owenoclee/tfw-go/cli-client/tfw/geo"

type Drawable interface {
	Draw(Screen) KeyCallbacks
	SetBounds(geo.Rect)
}
