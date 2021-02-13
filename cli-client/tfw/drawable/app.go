package drawable

import (
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type App struct {
	Child Drawable
}

var _ Drawable = &App{}

func (al *App) Draw(s tfw.Screen) KeyCallbacks {
	w, h := s.Size()
	al.Child.SetBounds(geo.Rect{
		TopLeft:     geo.Vector{0, 0},
		BottomRight: geo.Vector{w - 1, h - 1},
	})
	return al.Child.Draw(s)
}

func (_ *App) SetBounds(_ geo.Rect) {
	panic("SetBounds on App is not permitted")
}
