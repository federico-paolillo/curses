package hex

type Hex struct {
	x, y, z int
}

const Ne, Nw, E, Se, Sw, W = 0, 1, 2, 3, 4, 5

func New(x, y, z int) Hex {
	return Hex{x, y, z}
}

func (hex Hex) Distance(other Hex) int {
	return manatthanDistance(hex, other) / 2
}

func (hex Hex) Neighbors() [6]Hex {
	return [6]Hex{
		offset(hex, 0, -1, +1),
		offset(hex, +1, -1, 0),
		offset(hex, -1, 0, +1),
		offset(hex, -1, +1, 0),
		offset(hex, 0, +1, -1),
		offset(hex, +1, 0, -1),
	}
}

func (hex Hex) LineTo(other Hex) []Hex {
	distanceFromOther := diagonalDistance(hex, other)
	adjustedDistance := distanceFromOther + 1

	lineHexes := make([]Hex, adjustedDistance)

	for i := 0; i < adjustedDistance; i++ {
		step := float64(i) / float64(distanceFromOther)
		lineHexes[i] = interpolateHex(hex, other, step)
	}

	return lineHexes
}

func (hex Hex) IsValid() bool {
	return hex.x+hex.y+hex.z == 0
}

func (hex Hex) GetX() int {
	return hex.x
}

func (hex Hex) GetY() int {
	return hex.y
}

func (hex Hex) GetZ() int {
	return hex.z
}
