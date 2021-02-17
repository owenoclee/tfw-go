package tfw

import (
	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/geo"
)

type App struct {
	child Drawable
	quit  chan struct{}
}

func NewApp(quit chan struct{}, child Drawable) *App {
	return &App{
		child: child,
		quit:  quit,
	}
}

func (a *App) Run() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	ts, err := tcell.NewScreen()
	s := Screen{ts}
	defer s.Fini()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	s.Clear()

	var callbacks KeyCallbacks
	redraw := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyRune:
					f := callbacks.CallbackForKey(ev.Rune())
					if f != nil {
						f()
						redraw <- struct{}{}
					}
				}
			case *tcell.EventResize:
				redraw <- struct{}{}
			}
		}
	}()

	s.Show()
	for {
		select {
		case <-a.quit:
			return
		case <-redraw:
			callbacks = a.draw(s)
			s.Sync()
		}
	}
}

func (a *App) draw(s Screen) KeyCallbacks {
	w, h := s.Size()
	a.child.SetBounds(geo.Rect{
		TopLeft:     geo.Vector{0, 0},
		BottomRight: geo.Vector{w - 1, h - 1},
	})
	return a.child.Draw(s)
}
