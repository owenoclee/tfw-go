package component

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type WrappedText struct {
	Bounds geo.Rect
	Text   string
	Style  *tcell.Style
}

var _ tfw.Drawable = &WrappedText{}

func (wt *WrappedText) Draw(s tfw.Screen) tfw.KeyCallbacks {
	style := tcell.StyleDefault
	if wt.Style != nil {
		style = *wt.Style
	}

	words := strings.Split(wt.Text, " ")
	cursor := wt.Bounds.TopLeft
	for _, word := range words {
		needsNewLine := false
		// request new line if the word wont fit on this line
		if !wt.Bounds.VectorInBounds(cursor.Add(geo.Vector{len(word) - 1, 0})) {
			needsNewLine = true
		}
		// if there are no new lines available or the word cannot fit on a new line,
		// replace the last written character with an ellipsis
		if (needsNewLine && cursor.Y+1 > wt.Bounds.BottomRight.Y) || len(word) > wt.Bounds.HorizontalCells() {
			s.SetContent(cursor.Add(geo.Vector{-2, 0}), '…', style)
			return nil
		}
		// start new line
		if needsNewLine {
			cursor.X = wt.Bounds.TopLeft.X
			cursor.Y++
		}
		// write text
		for _, c := range word {
			s.SetContent(cursor, c, style)
			cursor = cursor.Add(geo.Vector{1, 0})
		}
		cursor = cursor.Add(geo.Vector{1, 0})
	}
	return nil
}

func (wt *WrappedText) SetBounds(r geo.Rect) {
	wt.Bounds = r
}
