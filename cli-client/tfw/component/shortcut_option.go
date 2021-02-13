package component

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

var shortcutStyle = tcell.Style{}.Foreground(tcell.ColorOrange)

type ShortcutOption struct {
	bounds   geo.Rect
	Shortcut rune
	Text     string
	Callback func()
}

var _ tfw.Drawable = &ShortcutOption{}

func (so *ShortcutOption) Draw(s tfw.Screen) tfw.KeyCallbacks {
	shortcutText := WrappedText{
		Bounds: geo.Rect{
			TopLeft:     so.bounds.TopLeft,
			BottomRight: so.bounds.TopLeft.Add(geo.Vector{2, 0}),
		},
		Text:  fmt.Sprintf("(%s)", string(so.Shortcut)),
		Style: shortcutStyle,
	}
	shortcutText.Draw(s)

	description := WrappedText{
		Bounds: geo.Rect{
			TopLeft:     so.bounds.TopLeft.Add(geo.Vector{4, 0}),
			BottomRight: so.bounds.BottomRight,
		},
		Text: so.Text,
	}
	description.Draw(s)

	callbacks := tfw.NewKeyCallbacks()
	callbacks.Register(so.Shortcut, so.Callback)
	return callbacks
}

func (so *ShortcutOption) SetBounds(r geo.Rect) {
	so.bounds = r
}
