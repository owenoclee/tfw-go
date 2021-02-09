package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

var rectPieces = map[piece]rune{
	TopLeftPiece:     '┌',
	BottomLeftPiece:  '└',
	TopRightPiece:    '┐',
	BottomRightPiece: '┘',
	HorizontalPiece:  '─',
	VerticalPiece:    '│',
}

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	ts, err := tcell.NewScreen()
	s := screen{ts}
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	s.Clear()

	for _, mc := range rectCells(vector{0, 0}, vector{46, 20}) {
		c := rectPieces[mc.kind]
		s.SetContent(mc.pos, c, tcell.StyleDefault)
	}
	for x, c := range "charlie" {
		for y := 1; y < 20; y++ {
			s.SetContent(vector{x + (y * 2), y}, c, tcell.Style{}.Foreground(tcell.Color(y)+tcell.ColorValid+32))

		}
	}

	s.Show()
	time.Sleep(time.Second * 5)
	s.Fini()
}
