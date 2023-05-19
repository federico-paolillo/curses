package main

import (
	"curses/cmd/visualizer/conversion"
	"curses/pkg/hex"
	"fmt"

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

	hex3 := hex.New(-2, 0, 2)

	coords3 := make([]conversion.ScreenCoords, 0, 6)

	vertices3 := hex3.Vertices()

	for _, vertex := range vertices3 {
		coords3 = append(coords3, conversion.ConvertVertexToCoords(vertex))
	}

	fmt.Println(vertices2[0], vertices[2])

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

		rl.DrawLine(int32(coords3[0].X), int32(coords3[0].Y), int32(coords3[1].X), int32(coords3[1].Y), rl.Green)
		rl.DrawLine(int32(coords3[1].X), int32(coords3[1].Y), int32(coords3[2].X), int32(coords3[2].Y), rl.Green)
		rl.DrawLine(int32(coords3[2].X), int32(coords3[2].Y), int32(coords3[3].X), int32(coords3[3].Y), rl.Green)
		rl.DrawLine(int32(coords3[3].X), int32(coords3[3].Y), int32(coords3[4].X), int32(coords3[4].Y), rl.Green)
		rl.DrawLine(int32(coords3[4].X), int32(coords3[4].Y), int32(coords3[5].X), int32(coords3[5].Y), rl.Green)
		rl.DrawLine(int32(coords3[5].X), int32(coords3[5].Y), int32(coords3[0].X), int32(coords3[0].Y), rl.Green)

		rl.EndDrawing()
	}

}
