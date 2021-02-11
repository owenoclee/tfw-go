package canvas

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/geom"
)

type Screen struct {
	tcell.Screen
}

func (s *Screen) SetContent(pos geom.Vector, c rune, style tcell.Style) {
	s.Screen.SetContent(pos.X, pos.Y, c, nil, style)
}
