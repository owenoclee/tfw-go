package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/canvas"
	"github.com/owenoclee/tfw-go/cli-client/drawable"
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

	box1 := &drawable.Box{
		Pieces: drawable.DefaultBoxPieces,
	}
	box2 := &drawable.Box{
		Pieces: drawable.DefaultBoxPieces,
	}
	box3 := &drawable.Box{
		Pieces: drawable.DefaultBoxPieces,
	}
	splits := &drawable.VerticalSplit{
		Children: []drawable.Drawable{box1, box2, box3},
	}
	app := &drawable.App{
		Child: splits,
	}
	callbacks := app.Draw(s)

	hasResized := false
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
					f := callbacks.CallbackForKey(ev.Rune())
					if f != nil {
						f()
						s.Sync()
					}
				}
			case *tcell.EventResize:
				hasResized = true
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
			app.Draw(s)
			if hasResized == true {
				s.Sync()
				hasResized = false
			}
		}
	}
}
