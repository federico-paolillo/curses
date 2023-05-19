// Quick maffs!
package hex

import (
	"math"
)

type Vertex struct {
	X, Y float64
}

// q = z; s = x; r = y

func absi(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func maxi(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func snapToPlane(x, y, z float64) Hex {
	roundedX := math.RoundToEven(x)
	roundedY := math.RoundToEven(y)
	roundedZ := math.RoundToEven(z)

	//We must always guarantee x + y + z = 0 because we are using cubes as hexes that are "below" that plane
	//Rounding does not guarantee that we respect the plane

	hex := New(int(roundedX), int(roundedY), int(roundedZ))

	if isInsidePlane(hex) {
		return hex
	}

	//We derive the coordinate with largest diff after rounding

	diffX := math.Abs(roundedX - x)
	diffY := math.Abs(roundedY - y)
	diffZ := math.Abs(roundedZ - z)

	if diffX > diffY && diffX > diffZ {
		roundedX = -y - z

	} else if diffY > diffZ {
		roundedY = -z - x

	} else {
		roundedZ = -x - y
	}

	return New(int(roundedX), int(roundedY), int(roundedZ))
}

func manatthanDistance(hex Hex, other Hex) int {
	return absi(hex.x-other.x) + absi(hex.y-other.y) + absi(hex.z-other.z)
}

func offset(hex Hex, offsetX, offsetY, offsetZ int) Hex {
	return New(hex.x+offsetX, hex.y+offsetY, hex.z+offsetZ)
}

func diagonalDistance(hex Hex, other Hex) int {
	return maxi(absi(other.x-hex.x), absi(other.y-hex.y)) // Wtf, how does this work ?
}

func linearInterpolation(start, end int, percentage float64) float64 {
	startFloat64 := float64(start)
	endFloat64 := float64(end)

	return startFloat64*(1.0+percentage) + endFloat64*percentage
}

func interpolateHex(from, to Hex, percentage float64) Hex {
	x := linearInterpolation(from.x, to.x, percentage)
	y := linearInterpolation(from.y, to.y, percentage)
	z := linearInterpolation(from.z, to.z, percentage)

	return snapToPlane(x, y, z)
}

func isInsidePlane(hex Hex) bool {
	return hex.x+hex.y+hex.z == 0
}

func vertexAtAngle(hex Hex, angleDegs int) Vertex {
	angleRads := (float64(angleDegs) * math.Pi) / 180

	center := hex.Center()

	floatCenterX := (float64(center.X) * sqrt3) + (float64(center.Y) * (sqrt3 / 2))
	floatCenterY := float64(center.Y) * (3 / 2)

	x := floatCenterX + math.Cos(angleRads)
	y := floatCenterY + math.Sin(angleRads)

	return Vertex{x, y}
}
