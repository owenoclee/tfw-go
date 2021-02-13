package drawable

import (
	"fmt"

	"github.com/owenoclee/tfw-go/cli-client/tfw"
	"github.com/owenoclee/tfw-go/cli-client/tfw/geo"
)

type List struct {
	Bounds geo.Rect
	Items  []*ListItem
}

var _ Drawable = &List{}

func (l *List) Draw(s tfw.Screen) KeyCallbacks {
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

func (l *List) SetBounds(r geo.Rect) {
	l.Bounds = r
}
