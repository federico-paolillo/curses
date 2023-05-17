package conversion

import (
	"curses/pkg/hex"
)

type ScreenCoords struct {
	X, Y int
}

const pixelMultiplier = 50

const vDistance = 3 / 2
const hDistance = 1.75

func ConvertVertexToCoords(vertex hex.Vertex) ScreenCoords {
	xWithOffset := (vertex.X + (hDistance / 2))
	yWithOffset := (vertex.Y + (1))

	xWithTranslation := xWithOffset //+ yWithOffset*(hDistance/2)
	yWithTranslation := yWithOffset // * vDistance

	screenX := xWithTranslation * pixelMultiplier
	screenY := yWithTranslation * pixelMultiplier

	return ScreenCoords{int(screenX) + 1, int(screenY)}
}
