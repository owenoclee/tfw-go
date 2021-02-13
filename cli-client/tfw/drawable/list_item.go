package drawable

import (
	"fmt"
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

var shortcutStyle = tcell.Style{}.Foreground(tcell.ColorOrange)

type ListItem struct {
	Bounds   geo.Rect
	Shortcut rune
	Text     string
}

var _ Drawable = &ListItem{}

func (li *ListItem) Draw(s tfw.Screen) KeyCallbacks {
	if li.Bounds.HorizontalCells() < 7 {
		panic("list item must be at least 7 cells wide")
	}

	shortcut := WrappedText{
		Bounds: geo.Rect{
			TopLeft:     li.Bounds.TopLeft,
			BottomRight: li.Bounds.TopLeft.Add(geo.Vector{2, 0}),
		},
		Text:  fmt.Sprintf("(%s)", string(li.Shortcut)),
		Style: shortcutStyle,
	}
	shortcut.Draw(s)

	description := WrappedText{
		Bounds: geo.Rect{
			TopLeft:     li.Bounds.TopLeft.Add(geo.Vector{4, 0}),
			BottomRight: li.Bounds.BottomRight,
		},
		Text: li.Text,
	}
	description.Draw(s)

	callbacks := NewKeyCallbacks()
	callbacks.Register(li.Shortcut, func() {
		s.SetContent(geo.Vector{48 + rand.Intn(20), 0}, 'A', tcell.StyleDefault)
	})
	return callbacks
}

func (li *ListItem) SetBounds(r geo.Rect) {
	li.Bounds = r
}
