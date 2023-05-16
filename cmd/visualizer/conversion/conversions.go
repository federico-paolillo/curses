package conversion

import (
	"curses/pkg/hex"
)

type ScreenCoords struct {
	X, Y int
}

const hexSizeUnitToPixelConversionRate = 10
const hexSize = 2
const hexSizeMultiplier = hexSize * hexSizeUnitToPixelConversionRate

const heightDistance = 3 / 2
const widthDistance = 1.75

func ConvertCenterToCoords(hex hex.Hex) ScreenCoords {
	center := hex.Center()

	floatX := float64(center.X)
	floatY := float64(center.Y)

	screenX := ((floatX * widthDistance) + floatY*(widthDistance/2)) * hexSizeMultiplier
	screenY := (floatY * heightDistance) * hexSizeMultiplier

	return ScreenCoords{int(screenX) + 100, int(screenY) + 100}
}

func ConvertVertexToCoords(hex hex.Hex, vertex hex.Vertex) ScreenCoords {
	centerCoords := ConvertCenterToCoords(hex)

	screenX := centerCoords.X + int(vertex.X)
	screenY := centerCoords.Y + int(vertex.Y)

	return ScreenCoords{int(screenX), int(screenY)}
}
