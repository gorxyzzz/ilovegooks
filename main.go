package main

import (
	_ "fmt"
	"math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_SIZE = 300

func main() {
	rl.SetConfigFlags(rl.FlagWindowUndecorated | rl.FlagWindowTransparent)
	rl.InitWindow(WINDOW_SIZE, WINDOW_SIZE, "Sliced Circle")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Set configuration variables
	centerX := float32(WINDOW_SIZE) / 2
	centerY := float32(WINDOW_SIZE) / 2
	radius := float32(WINDOW_SIZE) / 2
	
	// Change this number to slice it with more or fewer lines!
	numLines := 6 

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)

		// 1. Draw base circle background
		rl.DrawCircle(int32(centerX), int32(centerY), radius, rl.Maroon)

		// 2. Calculate and draw the slicing lines
		for i := 0; i < numLines; i++ {
			// Calculate angle in radians for this specific line
			angle := float64(i) * (math.Pi / float64(numLines))

			// Calculate the offsets from center using standard trigonometry
			offsetX := float32(math.Cos(angle)) * radius
			offsetY := float32(math.Sin(angle)) * radius

			// Line goes from one edge of the circle, through the center, to the opposite edge
			startPoint := rl.NewVector2(centerX + offsetX, centerY + offsetY)
			endPoint := rl.NewVector2(centerX - offsetX, centerY - offsetY)

			// Draw the slicing line (using Raywhite for high contrast)
			rl.DrawLineV(startPoint, endPoint, rl.RayWhite)
		}

		rl.EndDrawing()
	}
}
