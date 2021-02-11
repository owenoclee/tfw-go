package drawable

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/canvas"
	"github.com/owenoclee/tfw-go/cli-client/geom"
)

var shortcutStyle = tcell.Style{}.Foreground(tcell.ColorOrange)

type ListItem struct {
	Bounds   geom.Rect
	Shortcut rune
	Text     string
}

var _ Drawable = &ListItem{}

func (li *ListItem) Draw(s canvas.Screen) {
	if li.Bounds.HorizontalCells() < 7 {
		panic("list item must be at least 7 cells wide")
	}

	shortcut := WrappedText{
		Bounds: geom.Rect{
			TopLeft:     li.Bounds.TopLeft,
			BottomRight: li.Bounds.TopLeft.Add(geom.Vector{2, 0}),
		},
		Text:  fmt.Sprintf("(%s)", string(li.Shortcut)),
		Style: shortcutStyle,
	}
	shortcut.Draw(s)

	description := WrappedText{
		Bounds: geom.Rect{
			TopLeft:     li.Bounds.TopLeft.Add(geom.Vector{4, 0}),
			BottomRight: li.Bounds.BottomRight,
		},
		Text: li.Text,
	}
	description.Draw(s)
}
