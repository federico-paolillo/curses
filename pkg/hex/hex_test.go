package hex

import (
	"testing"
)

func TestCompareHexes(t *testing.T) { // Does golang work ? Jokes aside, ensures that you do not change New() behavior

	hexA := New(2, 0, -2)
	hexB := New(2, 0, -2)

	if hexA != hexB {
		t.Errorf("%v is not equal to %v. New() made two different structs with the same args. Did you return a pointer ?", hexA, hexB)
	}

}

func TestNeighborDistanceIsAlwaysOne(t *testing.T) {

	const expectedDistance = 1

	hex := New(0, 0, 0)

	hexNeighbors := hex.Neighbors()

	for _, neighbor := range hexNeighbors {

		distance := hex.Distance(neighbor)
		isValid := hex.IsValid()

		if distance != expectedDistance {
			t.Errorf("Neighbor %v should be distant %d but was distant %d", neighbor, expectedDistance, distance)
		}

		if !isValid {
			t.Errorf("Neighbor %v is outside the correct cube half", neighbor)
		}
	}
}

func TestLineBetweenTwoHexes(t *testing.T) {
	start := New(0, 0, 0)
	end := New(4, -4, 0)

	expectedLinePieces := [5]Hex{
		New(0, 0, 0),
		New(1, -1, 0),
		New(2, -2, 0),
		New(3, -3, 0),
		New(4, -4, 0),
	}

	lineHexes := start.LineTo(end)

	if len(lineHexes) != len(expectedLinePieces) {
		t.Fatalf("Generated line between %v and %v is too short", start, end)
	}

	for i, lineHex := range lineHexes {
		if expectedLinePieces[i] != lineHex {
			t.Errorf("Generated line between %v and %v has an invalid segment at %d. Should be %v but was %v", start, end, i, expectedLinePieces[i], lineHex)
		}
	}
}
