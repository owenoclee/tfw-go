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
		text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent ultrices sapien aca tortor posuere, semper iaculis lectus volutpat. Nulla facilisi. Donec varius aliquam efficitur. Sed dictum urna tellus, et laciniaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa erat hendrerit eget. Quisque dolor nunc, hendrerit a pharetra sed, pretium a magna. Mauris dignissim quam elit, non accumsan ipsum laoreet vitae. Nam dapibus venenatis sollicitudin. Ut tempor metus non vestibulum laoreet.",
	}
	wt.Draw(s)

	wt2 := wrappedText{
		bounds: rect{
			topLeft:     vector{50, 50},
			bottomRight: vector{50, 50},
		},
		text: "G",
	}
	wt2.Draw(s)

	s.Show()
	time.Sleep(time.Second * 5)
	s.Fini()
}
