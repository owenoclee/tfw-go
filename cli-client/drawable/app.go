package drawable

import (
	"github.com/owenoclee/tfw-go/cli-client/canvas"
	"github.com/owenoclee/tfw-go/cli-client/geom"
)

type App struct {
	Child Drawable
}

var _ Drawable = &App{}

func (al *App) Draw(s canvas.Screen) KeyCallbacks {
	w, h := s.Size()
	al.Child.SetBounds(geom.Rect{
		TopLeft:     geom.Vector{0, 0},
		BottomRight: geom.Vector{w - 1, h - 1},
	})
	return al.Child.Draw(s)
}

func (_ *App) SetBounds(_ geom.Rect) {
	panic("SetBounds on App is not permitted")
}
