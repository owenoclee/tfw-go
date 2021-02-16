package component

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

type ShortcutOption struct {
	bounds        geo.Rect
	Shortcut      rune
	Text          string
	Callback      func()
	ShortcutStyle *tcell.Style
	TextStyle     *tcell.Style
}

var _ tfw.Drawable = &ShortcutOption{}

func (so *ShortcutOption) Draw(s tfw.Screen) tfw.KeyCallbacks {
	shortcutStyle := tcell.Style{}.Foreground(tcell.ColorOrange)
	if so.ShortcutStyle != nil {
		shortcutStyle = *so.ShortcutStyle
	}
	textStyle := tcell.StyleDefault
	if so.TextStyle != nil {
		textStyle = *so.TextStyle
	}

	shortcutText := WrappedText{
		Bounds: geo.Rect{
			TopLeft:     so.bounds.TopLeft,
			BottomRight: so.bounds.TopLeft.Add(geo.Vector{2, 0}),
		},
		Text:  fmt.Sprintf("(%s)", string(so.Shortcut)),
		Style: &shortcutStyle,
	}
	shortcutText.Draw(s)

	description := WrappedText{
		Bounds: geo.Rect{
			TopLeft:     so.bounds.TopLeft.Add(geo.Vector{4, 0}),
			BottomRight: so.bounds.BottomRight,
		},
		Text:  so.Text,
		Style: &textStyle,
	}
	description.Draw(s)

	callbacks := tfw.NewKeyCallbacks()
	callbacks.Register(so.Shortcut, so.Callback)
	return callbacks
}

func (so *ShortcutOption) SetBounds(r geo.Rect) {
	so.bounds = r
}

func (so *ShortcutOption) MinBounds(topLeft geo.Vector) geo.Rect {
	return geo.Rect{
		TopLeft:     topLeft,
		BottomRight: topLeft.Add(geo.Vector{3 + len(so.Text), 0}),
	}
}
