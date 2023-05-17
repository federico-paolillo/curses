package hex

type Orientation = int

type Hex struct {
	x, y, z     int
	orientation Orientation
}

type Center struct {
	X, Y int
}

const (
	Pointy = iota
	Flatty
)

const sqrt3 = 1.73
const unit = 1

func New(x, y, z int) Hex {
	return Hex{x, y, z, Pointy}
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

func (hex Hex) Width() float64 {
	if hex.orientation == Pointy {
		return sqrt3
	} else {
		return unit * 2
	}
}

func (hex Hex) Height() float64 {
	if hex.orientation == Pointy {
		return unit * 2
	} else {
		return sqrt3
	}
}

func (hex Hex) Vertices() [6]Vertex {
	return [6]Vertex{
		vertexAtAngle(hex, 30),
		vertexAtAngle(hex, 90),
		vertexAtAngle(hex, 150),
		vertexAtAngle(hex, 210),
		vertexAtAngle(hex, 270),
		vertexAtAngle(hex, 330),
	}
}

func (hex Hex) Center() Center {
	return Center{X: hex.z, Y: hex.y}
}
