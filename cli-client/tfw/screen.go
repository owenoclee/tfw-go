package tfw

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

// Screen is a wrapper around tcell.Screen to provide some convenience methods.
type Screen struct {
	tcell.Screen
}

// SetContent acts the same as tcell.Screen.SetContent but uses a Vector to
// specify the location on screen to set. It also does not support combining
// characters for the sake of brevity.
func (s *Screen) SetContent(pos geo.Vector, c rune, style tcell.Style) {
	s.Screen.SetContent(pos.X, pos.Y, c, nil, style)
}
