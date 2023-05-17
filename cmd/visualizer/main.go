package main

import (
	"curses/cmd/visualizer/conversion"
	"curses/pkg/hex"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	hex1 := hex.New(-1, 0, 1)

	coords := make([]conversion.ScreenCoords, 0, 6)

	vertices := hex1.Vertices()

	for _, vertex := range vertices {
		coords = append(coords, conversion.ConvertVertexToCoords(vertex))
	}

	hex2 := hex.New(0, 0, 0)

	coords2 := make([]conversion.ScreenCoords, 0, 6)

	vertices2 := hex2.Vertices()

	for _, vertex := range vertices2 {
		coords2 = append(coords2, conversion.ConvertVertexToCoords(vertex))
	}

	// center1 := hex.Vertex{0, 0}
	// // center2 := hex.Vertex{0, 1}
	// center3 := hex.Vertex{1, 0}

	// center1Coords := conversion.ConvertVertexToCoords(center1)
	// // center2Coords := conversion.ConvertVertexToCoords(center2)
	// center3Coords := conversion.ConvertVertexToCoords(center3)

	rl.InitWindow(800, 450, "curses")

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawLine(int32(coords[0].X), int32(coords[0].Y), int32(coords[1].X), int32(coords[1].Y), rl.White)
		rl.DrawLine(int32(coords[1].X), int32(coords[1].Y), int32(coords[2].X), int32(coords[2].Y), rl.White)
		rl.DrawLine(int32(coords[2].X), int32(coords[2].Y), int32(coords[3].X), int32(coords[3].Y), rl.White)
		rl.DrawLine(int32(coords[3].X), int32(coords[3].Y), int32(coords[4].X), int32(coords[4].Y), rl.White)
		rl.DrawLine(int32(coords[4].X), int32(coords[4].Y), int32(coords[5].X), int32(coords[5].Y), rl.White)
		rl.DrawLine(int32(coords[5].X), int32(coords[5].Y), int32(coords[0].X), int32(coords[0].Y), rl.White)

		rl.DrawLine(int32(coords2[0].X), int32(coords2[0].Y), int32(coords2[1].X), int32(coords2[1].Y), rl.Red)
		rl.DrawLine(int32(coords2[1].X), int32(coords2[1].Y), int32(coords2[2].X), int32(coords2[2].Y), rl.Red)
		rl.DrawLine(int32(coords2[2].X), int32(coords2[2].Y), int32(coords2[3].X), int32(coords2[3].Y), rl.Red)
		rl.DrawLine(int32(coords2[3].X), int32(coords2[3].Y), int32(coords2[4].X), int32(coords2[4].Y), rl.Red)
		rl.DrawLine(int32(coords2[4].X), int32(coords2[4].Y), int32(coords2[5].X), int32(coords2[5].Y), rl.Red)
		rl.DrawLine(int32(coords2[5].X), int32(coords2[5].Y), int32(coords2[0].X), int32(coords2[0].Y), rl.Red)

		// rl.DrawLine(int32(center1Coords.X), int32(center1Coords.Y), int32(center3Coords.X), int32(center3Coords.Y), rl.Red)

		rl.EndDrawing()
	}

}
