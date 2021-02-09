package main

import "github.com/gdamore/tcell/v2"

type screen struct {
	tcell.Screen
}

func (s *screen) SetContent(v vector, r rune, st tcell.Style) {
	s.Screen.SetContent(v.x, v.y, r, nil, st)
}
