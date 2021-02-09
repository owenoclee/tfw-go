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
	s := &screen{ts}
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	s.Clear()

	b := box{
		pieces: rectPieces,
		bounds: rect{
			topLeft:     vector{0, 0},
			bottomRight: vector{46, 10},
		},
	}
	b.Draw(s)
	wt := wrappedText{
		bounds: rect{
			topLeft:     vector{1, 1},
			bottomRight: vector{45, 9},
		},
		text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent ultrices sapien aca tortor posuere, semper iaculis lectus volutpat. Nulla facilisi. Donec varius aliquam efficitur. Sed dictum urna tellus, et lacinia erat hendrerit eget. Quisque dolor nunc, hendrerit a pharetra sed, pretium a magna. Mauris dignissim quam elit, non accumsan ipsum laoreet vitae. Nam dapibus venenatis sollicitudin. Ut tempor metus non vestibulum laoreet.",
	}
	wt.Draw(s)

	li := listItem{
		bounds: rect{
			topLeft:     vector{0, 12},
			bottomRight: vector{46, 13},
		},
		shortcut: 'a',
		text:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent ultrices sapien aca tortor posuere, semper iaculis lectus volutpat. Nulla facilisi. Donec varius aliquam efficitur. Sed dictum urna tellus, et lacinia erat hendrerit eget. Quisque dolor nunc, hendrerit a pharetra sed, pretium a magna. Mauris dignissim quam elit, non accumsan ipsum laoreet vitae. Nam dapibus venenatis sollicitudin. Ut tempor metus non vestibulum laoreet.",
	}
	li2 := listItem{
		bounds: rect{
			topLeft:     vector{0, 15},
			bottomRight: vector{46, 16},
		},
		shortcut: 'b',
		text:     "stonks!",
	}

	l := list{
		bounds: rect{
			topLeft:     vector{0, 12},
			bottomRight: vector{46, 16},
		},
		items: []*listItem{
			&li,
			&li2,
		},
	}
	l.Draw(s)

	s.Show()
	time.Sleep(time.Second * 5)
	s.Fini()
}
