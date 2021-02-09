package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Drawable interface {
	Draw(s *screen)
}

type box struct {
	bounds rect
	pieces map[piece]rune
}

func (b *box) Draw(s *screen) {
	for _, mc := range b.rectCells() {
		c := b.pieces[mc.kind]
		s.SetContent(mc.pos, c, tcell.StyleDefault)
	}
}

type piece int

// pieces
const (
	TopLeftPiece piece = iota
	BottomLeftPiece
	TopRightPiece
	BottomRightPiece
	HorizontalPiece
	VerticalPiece
)

type cellMeta struct {
	pos  vector
	kind piece
}

func (b *box) rectCells() []cellMeta {
	if !b.bounds.IsValid() {
		return nil
	}

	var metaCells []cellMeta
	for x := b.bounds.topLeft.x; x <= b.bounds.bottomRight.x; x++ {
		topCell := cellMeta{
			pos:  vector{x, b.bounds.topLeft.y},
			kind: HorizontalPiece,
		}
		bottomCell := cellMeta{
			pos:  vector{x, b.bounds.bottomRight.y},
			kind: HorizontalPiece,
		}
		switch x {
		case b.bounds.topLeft.x:
			topCell.kind = TopLeftPiece
			bottomCell.kind = BottomLeftPiece
		case b.bounds.bottomRight.x:
			topCell.kind = TopRightPiece
			bottomCell.kind = BottomRightPiece
		}
		metaCells = append(metaCells, topCell, bottomCell)
	}
	for y := b.bounds.topLeft.y + 1; y <= b.bounds.bottomRight.y-1; y++ {
		topCell := cellMeta{
			pos:  vector{b.bounds.topLeft.x, y},
			kind: VerticalPiece,
		}
		bottomCell := cellMeta{
			pos:  vector{b.bounds.bottomRight.x, y},
			kind: VerticalPiece,
		}
		metaCells = append(metaCells, topCell, bottomCell)
	}
	return metaCells
}

type wrappedText struct {
	bounds rect
	text   string
	style  tcell.Style
}

func (wt *wrappedText) Draw(s *screen) {
	words := strings.Split(wt.text, " ")
	cursor := wt.bounds.topLeft
	for _, word := range words {
		needsNewLine := false
		// request new line if the word wont fit on this line
		if !wt.bounds.VectorInBounds(cursor.add(vector{len(word) - 1, 0})) {
			needsNewLine = true
		}
		// if there are no new lines available or the word cannot fit on a new line,
		// replace the last written character with an ellipsis
		if (needsNewLine && cursor.y+1 > wt.bounds.bottomRight.y) || len(word) > wt.bounds.HorizontalCells() {
			s.SetContent(cursor.add(vector{-2, 0}), '…', wt.style)
			return
		}
		// start new line
		if needsNewLine {
			cursor.x = wt.bounds.topLeft.x
			cursor.y++
		}
		// write text
		for _, c := range word {
			s.SetContent(cursor, c, wt.style)
			cursor = cursor.add(vector{1, 0})
		}
		cursor = cursor.add(vector{1, 0})
	}
}

type listItem struct {
	bounds   rect
	shortcut rune
	text     string
}

var shortcutStyle = tcell.Style{}.Foreground(tcell.ColorOrange)

func (li *listItem) Draw(s *screen) {
	if li.bounds.HorizontalCells() < 7 {
		panic("must be wider than 7 cells")
	}

	// draw the shortcut key next to the list item
	shortcut := wrappedText{
		bounds: rect{
			topLeft:     li.bounds.topLeft,
			bottomRight: li.bounds.topLeft.add(vector{2, 0}),
		},
		text:  fmt.Sprintf("(%s)", string(li.shortcut)),
		style: shortcutStyle,
	}
	shortcut.Draw(s)

	description := wrappedText{
		bounds: rect{
			topLeft:     li.bounds.topLeft.add(vector{4, 0}),
			bottomRight: li.bounds.bottomRight,
		},
		text: li.text,
	}
	description.Draw(s)
}

type list struct {
	bounds rect
	items  []*listItem
}

func (l *list) Draw(s *screen) {
	// assert all list items are within the bounds of the list
	for i, li := range l.items {
		if !l.bounds.RectInBounds(li.bounds) {
			panic(fmt.Sprintf("list item %d is out of the screen bounds of the list", i))
		}
	}

	for _, li := range l.items {
		li.Draw(s)
	}
}
