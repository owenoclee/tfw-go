package main

import (
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
		// if there are no new lines available, replace the last written character with an ellipsis
		if needsNewLine && cursor.y+1 > wt.bounds.bottomRight.y {
			s.SetContent(cursor.add(vector{-2, 0}), '…', tcell.StyleDefault)
			return
		}
		// start new line
		if needsNewLine {
			cursor.x = wt.bounds.topLeft.x
			cursor.y++
			// if the word still wont fit, then there's nothing we can do so truncate and bail
			// TODO: consider adding mechanism to break a single word over a line?
			if len(word) > wt.bounds.HorizontalCells() {
				s.SetContent(cursor, '…', tcell.StyleDefault)
				return
			}
		}
		// write text
		for _, c := range word {
			s.SetContent(cursor, c, tcell.StyleDefault)
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

func (li *listItem) Draw(s *screen) {
	// stonks
}
