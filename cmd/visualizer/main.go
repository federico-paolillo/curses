package main

import (
	"curses/cmd/visualizer/conversion"
	"curses/pkg/hex"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	hex1 := hex.New(0, 0, 0)

	verticess := hex1.Vertices(20)

	vertices := [6]conversion.ScreenCoords{
		conversion.ConvertVertexToCoords(hex1, verticess[0]),
		conversion.ConvertVertexToCoords(hex1, verticess[1]),
		conversion.ConvertVertexToCoords(hex1, verticess[2]),
		conversion.ConvertVertexToCoords(hex1, verticess[3]),
		conversion.ConvertVertexToCoords(hex1, verticess[4]),
		conversion.ConvertVertexToCoords(hex1, verticess[5]),
	}

	fmt.Print(vertices)

	rl.InitWindow(800, 450, "curses")

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawLine(int32(vertices[0].X), int32(vertices[0].Y), int32(vertices[1].X), int32(vertices[1].Y), rl.White)
		rl.DrawLine(int32(vertices[1].X), int32(vertices[1].Y), int32(vertices[2].X), int32(vertices[2].Y), rl.White)
		rl.DrawLine(int32(vertices[2].X), int32(vertices[2].Y), int32(vertices[3].X), int32(vertices[3].Y), rl.White)
		rl.DrawLine(int32(vertices[3].X), int32(vertices[3].Y), int32(vertices[4].X), int32(vertices[4].Y), rl.White)
		rl.DrawLine(int32(vertices[4].X), int32(vertices[4].Y), int32(vertices[5].X), int32(vertices[5].Y), rl.White)
		rl.DrawLine(int32(vertices[5].X), int32(vertices[5].Y), int32(vertices[0].X), int32(vertices[0].Y), rl.White)

		rl.EndDrawing()
	}

}
