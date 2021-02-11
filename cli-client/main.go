package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/canvas"
	"github.com/owenoclee/tfw-go/cli-client/drawable"
	"github.com/owenoclee/tfw-go/cli-client/geom"
)

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	ts, err := tcell.NewScreen()
	s := canvas.Screen{ts}
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	s.Clear()

	b := drawable.Box{
		Pieces: drawable.DefaultBoxPieces,
		Bounds: geom.Rect{
			TopLeft:     geom.Vector{0, 0},
			BottomRight: geom.Vector{46, 10},
		},
	}
	b.Draw(s)
	wt := drawable.WrappedText{
		Bounds: geom.Rect{
			TopLeft:     geom.Vector{1, 1},
			BottomRight: geom.Vector{45, 9},
		},
		Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent ultrices sapien aca tortor posuere, semper iaculis lectus volutpat. Nulla facilisi. Donec varius aliquam efficitur. Sed dictum urna tellus, et lacinia erat hendrerit eget. Quisque dolor nunc, hendrerit a pharetra sed, pretium a magna. Mauris dignissim quam elit, non accumsan ipsum laoreet vitae. Nam dapibus venenatis sollicitudin. Ut tempor metus non vestibulum laoreet.",
	}
	wt.Draw(s)

	li := drawable.ListItem{
		Bounds: geom.Rect{
			TopLeft:     geom.Vector{0, 12},
			BottomRight: geom.Vector{46, 13},
		},
		Shortcut: 'a',
		Text:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent ultrices sapien aca tortor posuere, semper iaculis lectus volutpat. Nulla facilisi. Donec varius aliquam efficitur. Sed dictum urna tellus, et lacinia erat hendrerit eget. Quisque dolor nunc, hendrerit a pharetra sed, pretium a magna. Mauris dignissim quam elit, non accumsan ipsum laoreet vitae. Nam dapibus venenatis sollicitudin. Ut tempor metus non vestibulum laoreet.",
	}
	li2 := drawable.ListItem{
		Bounds: geom.Rect{
			TopLeft:     geom.Vector{0, 15},
			BottomRight: geom.Vector{46, 16},
		},
		Shortcut: 'b',
		Text:     "stonks!",
	}

	l := drawable.List{
		Bounds: geom.Rect{
			TopLeft:     geom.Vector{0, 12},
			BottomRight: geom.Vector{46, 16},
		},
		Items: []*drawable.ListItem{
			&li,
			&li2,
		},
	}
	l.Draw(s)

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyRune:
					if ev.Rune() == 'q' {
						close(quit)
						return
					}
				}
			}
		}
	}()

	s.Show()
	for {
		select {
		case <-quit:
			s.Fini()
			return
		case <-time.After(time.Millisecond * 50):
			l.Draw(s)
		}
	}
}
