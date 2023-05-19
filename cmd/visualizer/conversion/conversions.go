package conversion

import (
	"curses/pkg/hex"
)

type ScreenCoords struct {
	X, Y int
}

const pixelMultiplier = 100

const halfHDistance = 0.875

func ConvertVertexToCoords(vertex hex.Vertex) ScreenCoords {
	xWithOffset := vertex.X + halfHDistance
	yWithOffset := vertex.Y + 1

	screenX := xWithOffset * pixelMultiplier
	screenY := yWithOffset * pixelMultiplier

	return ScreenCoords{int(screenX) + 1, int(screenY)}
}
