package tfw

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/geo"
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

// SetRegion sets all cells within a geo.Rect to the given rune and style.
func (s *Screen) SetRegion(r geo.Rect, c rune, style tcell.Style) {
	for _, cell := range r.CellLocations() {
		s.SetContent(cell, c, style)
	}
}

// ClearRegion sets all cells within a geo.Rect to blank.
func (s *Screen) ClearRegion(r geo.Rect) {
	s.SetRegion(r, ' ', tcell.StyleDefault)
}
