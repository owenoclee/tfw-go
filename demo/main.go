package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/component"
	"github.com/owenoclee/tfw-go/layout"
)

var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer eget nibh et risus porta semper id id diam. Nam blandit, arcu at eleifend facilisis, augue nunc pulvinar nisl, a finibus sapien orci quis mi. Donec malesuada, tellus vitae finibus ornare, nunc nibh consectetur risus, sed porttitor risus nisl a libero. Nullam odio lectus, pellentesque sed ipsum a, convallis sodales odio. Nam justo sapien, posuere eget felis at, accumsan varius augue. Integer pulvinar tempor consequat. Praesent in maximus massa. Fusce posuere sit amet ligula eu condimentum. Nullam porta augue maximus lectus euismod, porttitor placerat risus facilisis. Nunc interdum neque metus, quis suscipit dui pretium vitae."
var number = 0

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

	exampleStyle := tcell.Style{}.Foreground(tcell.ColorBlue)
	exampleStyle2 := tcell.Style{}.Foreground(tcell.ColorOrange)
	leftBoxes := &layout.HorizontalSplit{
		Children: []tfw.Drawable{
			&component.Box{
				Child: &component.WrappedText{
					Text:  loremIpsum,
					Style: &exampleStyle,
				},
			},
			&component.Box{
				Child: &component.WrappedText{
					Text:  loremIpsum,
					Style: &exampleStyle2,
				},
			},
		},
	}

	numberDisplay := &component.WrappedText{
		Text: fmt.Sprint(number),
	}
	middleBox := &component.Box{
		Child: &layout.Margin{Child: &layout.Rows{
			Children: []tfw.Drawable{
				numberDisplay,
				&component.ShortcutOption{
					Shortcut: 'k',
					Text:     "Increase number",
					Callback: func() { number++; numberDisplay.SetText(fmt.Sprint(number)) },
				},
				&component.ShortcutOption{
					Shortcut: 'j',
					Text:     "Decrease number",
					Callback: func() { number--; numberDisplay.SetText(fmt.Sprint(number)) },
				},
			},
			RowLines: 2,
		},
			Top:    1,
			Bottom: 1,
			Left:   2,
			Right:  2,
		},
	}

	rightBox := &component.Box{
		Child: &layout.Columns{
			Children: []tfw.Drawable{
				&component.WrappedText{
					Text: "column 1",
				},
				&component.WrappedText{
					Text: "column 2",
				},
				&component.WrappedText{
					Text: "column 3",
				},
			},
			ColumnCells: 12,
		},
	}

	app := &tfw.App{
		Child: &layout.WithToolbar{
			Primary: &component.Box{
				Child: &layout.VerticalSplit{
					Children: []tfw.Drawable{
						leftBoxes,
						middleBox,
						rightBox,
					},
				},
			},
			BarElements: []tfw.MinBoundableDrawable{
				&component.ShortcutOption{
					Shortcut: 'q',
					Text:     "quit",
				},
			},
			ElementGap: 1,
		},
	}

	var callbacks tfw.KeyCallbacks
	var needsRedraw bool
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
						needsRedraw = true
					}
				}
			case *tcell.EventResize:
				needsRedraw = true
			}
		}
	}()

	s.Show()
	for {
		select {
		case <-quit:
			return
		case <-time.After(time.Millisecond * 50):
			if needsRedraw == true {
				callbacks = app.Draw(s)
				s.Sync()
				needsRedraw = false
			}
		}
	}
}
