package drawable

import "github.com/owenoclee/tfw-go/cli-client/canvas"

type KeyCallbacks map[rune]func()

func NewKeyCallbacks() KeyCallbacks {
	return make(map[rune]func())
}

func (kc KeyCallbacks) Register(key rune, callback func()) {
	kc[key] = callback
}

func (kc KeyCallbacks) Push(kc2 KeyCallbacks) {
	for k, c := range kc2 {
		kc[k] = c
	}
}

func (kc KeyCallbacks) CallbackForKey(key rune) func() {
	return kc[key]
}

type Drawable interface {
	Draw(canvas.Screen) KeyCallbacks
}
