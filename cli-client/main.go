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
		Child: &drawable.HorizontalSplit{
			Children: []drawable.Drawable{
				&drawable.Box{},
				&drawable.Box{},
			},
		},
	}
	box2 := &drawable.Box{}
	box3 := &drawable.Box{}
	splits := &drawable.VerticalSplit{
		Children: []drawable.Drawable{box1, box2, box3},
	}
	container := &drawable.Box{
		Child: splits,
	}
	app := &drawable.App{
		Child: container,
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
