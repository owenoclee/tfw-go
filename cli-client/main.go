package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/component"
	"github.com/owenoclee/tfw-go/cli-client/tfw/layout"
)

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	ts, err := tcell.NewScreen()
	s := tfw.Screen{ts}
	defer s.Fini()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	s.Clear()

	box1 := &component.Box{
		Child: &layout.HorizontalSplit{
			Children: []tfw.Drawable{
				&component.Box{},
				&component.Box{},
			},
		},
	}
	box2 := &component.Box{
		Child: &layout.Rows{
			Children: []tfw.Drawable{
				&component.WrappedText{
					Text: "stonks!",
				},
				&component.WrappedText{
					Text: "The quick brown fox jumps over the lazy dog.",
				},
				&component.Box{},
				&component.Box{},
				&component.Box{},
				&component.Box{},
				&component.Box{},
				&component.Box{},
			},
			RowLines: 2,
		},
	}
	box3 := &component.Box{
		Child: &layout.Rows{
			Children: []tfw.Drawable{
				&component.ShortcutOption{
					Shortcut: 'a',
					Text:     "press me!",
					Callback: func() { println("ah yes") },
				},
				&layout.Columns{
					Children: []tfw.Drawable{
						&component.WrappedText{
							Text: "sweet!",
						},
						&component.WrappedText{
							Text: "cool!",
						},
						&component.WrappedText{
							Text: "nice!",
						},
					},
					ColumnCells: 7,
				},
			},
			RowLines: 2,
		},
	}
	splits := &layout.VerticalSplit{
		Children: []tfw.Drawable{box1, box2, box3},
	}
	container := &component.Box{
		Child: splits,
	}
	app := &tfw.App{
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
