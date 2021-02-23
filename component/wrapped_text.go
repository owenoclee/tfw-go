package component

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/owenoclee/tfw-go"
	"github.com/owenoclee/tfw-go/geo"
)

var _ tfw.Drawable = (*wrappedText)(nil)
var _ tfw.HasText = (*wrappedText)(nil)

type wrappedText struct {
	bounds  geo.Rect
	visible bool
	text    string
	style   *tcell.Style
}

func NewWrappedText(text string) *wrappedText {
	return &wrappedText{
		visible: true,
		text:    text,
	}
}

func NewWrappedTextWithStyle(text string, style *tcell.Style) *wrappedText {
	return &wrappedText{
		visible: true,
		text:    text,
		style:   style,
	}
}

func (wt *wrappedText) Draw(s tfw.Screen) tfw.KeyCallbacks {
	style := tcell.StyleDefault
	if wt.style != nil {
		style = *wt.style
	}

	s.SetRegion(wt.bounds, ' ', style)

	words := strings.Split(wt.text, " ")
	cursor := wt.bounds.TopLeft
	for _, word := range words {
		needsNewLine := false
		// request new line if the word wont fit on this line
		if !wt.bounds.VectorInBounds(cursor.Add(geo.Vector{len(word) - 1, 0})) {
			needsNewLine = true
		}
		// if there are no new lines available or the word cannot fit on a new line,
		// replace the last written character with an ellipsis
		if (needsNewLine && cursor.Y+1 > wt.bounds.BottomRight.Y) || len(word) > wt.bounds.HorizontalCells() {
			s.SetContent(cursor.Add(geo.Vector{-2, 0}), '…', style)
			return nil
		}
		// start new line
		if needsNewLine {
			cursor.X = wt.bounds.TopLeft.X
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

func (wt *wrappedText) SetBounds(r geo.Rect) {
	wt.bounds = r
}

func (wt *wrappedText) SetVisible(visible bool) {
	wt.visible = visible
}

func (wt *wrappedText) Visible() bool {
	return wt.visible
}

func (wt *wrappedText) SetText(text string) {
	wt.text = text
}

func (wt *wrappedText) Text() string {
	return wt.text
}
