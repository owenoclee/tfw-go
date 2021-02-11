package drawable

import "github.com/owenoclee/tfw-go/cli-client/canvas"

type Drawable interface {
	Draw(canvas.Screen)
}
