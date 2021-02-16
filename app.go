package tfw

import "github.com/owenoclee/tfw-go/geo"

type App struct {
	Child Drawable
}

func (a *App) Draw(s Screen) KeyCallbacks {
	w, h := s.Size()
	a.Child.SetBounds(geo.Rect{
		TopLeft:     geo.Vector{0, 0},
		BottomRight: geo.Vector{w - 1, h - 1},
	})
	return a.Child.Draw(s)
}
