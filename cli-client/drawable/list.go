package drawable

import (
	"fmt"

	"github.com/owenoclee/tfw-go/cli-client/canvas"
	"github.com/owenoclee/tfw-go/cli-client/geom"
)

type List struct {
	Bounds geom.Rect
	Items  []*ListItem
}

var _ Drawable = &List{}

func (l *List) Draw(s canvas.Screen) KeyCallbacks {
	// assert all list items are within the bounds of the list
	for i, li := range l.Items {
		if !l.Bounds.RectInBounds(li.Bounds) {
			panic(fmt.Sprintf("list item %d is out of the screen bounds of the list", i))
		}
	}

	callbacks := NewKeyCallbacks()
	for _, li := range l.Items {
		callbacks.Push(li.Draw(s))
	}

	return callbacks
}

func (l *List) SetBounds(r geom.Rect) {
	l.Bounds = r
}
