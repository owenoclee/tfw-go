package component

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

var _ tfw.Drawable = (*shortcutOption)(nil)
var _ tfw.MinBoundable = (*shortcutOption)(nil)

type shortcutOption struct {
	bounds        geo.Rect
	visible       bool
	shortcut      rune
	text          string
	callback      func()
	shortcutStyle *tcell.Style
	textStyle     *tcell.Style
}

func NewShortcutOption(shortcut rune, text string, callback func()) *shortcutOption {
	return &shortcutOption{
		visible:  true,
		shortcut: shortcut,
		text:     text,
		callback: callback,
	}
}

func NewShortcutOptionWithStyles(shortcut rune, text string, callback func(), shortcutStyle, textStyle *tcell.Style) *shortcutOption {
	return &shortcutOption{
		visible:       true,
		shortcut:      shortcut,
		text:          text,
		callback:      callback,
		shortcutStyle: shortcutStyle,
		textStyle:     textStyle,
	}
}

func (so *shortcutOption) Draw(s tfw.Screen) tfw.KeyCallbacks {
	shortcutStyle := tcell.Style{}.Foreground(tcell.ColorOrange)
	if so.shortcutStyle != nil {
		shortcutStyle = *so.shortcutStyle
	}
	textStyle := tcell.StyleDefault
	if so.textStyle != nil {
		textStyle = *so.textStyle
	}

	shortcutText := NewWrappedTextWithStyle(fmt.Sprintf("(%s)", string(so.shortcut)), &shortcutStyle)
	shortcutText.SetBounds(geo.Rect{
		TopLeft:     so.bounds.TopLeft,
		BottomRight: so.bounds.TopLeft.Add(geo.Vector{2, 0}),
	})
	shortcutText.Draw(s)

	description := NewWrappedTextWithStyle(so.text, &textStyle)
	description.SetBounds(geo.Rect{
		TopLeft:     so.bounds.TopLeft.Add(geo.Vector{4, 0}),
		BottomRight: so.bounds.BottomRight,
	})
	description.Draw(s)

	callbacks := tfw.NewKeyCallbacks()
	callbacks.Register(so.shortcut, so.callback)
	return callbacks
}

func (so *shortcutOption) SetBounds(r geo.Rect) {
	so.bounds = r
}

func (so *shortcutOption) SetVisible(visible bool) {
	so.visible = visible
}

func (so *shortcutOption) Visible() bool {
	return so.visible
}

func (so *shortcutOption) MinBounds(topLeft geo.Vector) geo.Rect {
	return geo.Rect{
		TopLeft:     topLeft,
		BottomRight: topLeft.Add(geo.Vector{3 + len(so.text), 0}),
	}
}
