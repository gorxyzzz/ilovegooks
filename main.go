
package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const CIRCLE_SIZE = 300
const CIRCLE_BORDER_SIZE = CIRCLE_SIZE + 10
var CIRCLE_COLOR = rl.Color { 24, 24, 24, 255 };


type Slice struct {
	x rl.Vector2
	y rl.Vector2
	text string
}

func main() {
	rl.SetConfigFlags(
		rl.FlagWindowUndecorated |
			rl.FlagWindowTransparent |
			rl.FlagMsaa4xHint,
	)

	rl.InitWindow(CIRCLE_SIZE, CIRCLE_SIZE, "Sliced Circle")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	centerX := float32(CIRCLE_SIZE) / 2
	centerY := float32(CIRCLE_SIZE) / 2
	radius := float32(CIRCLE_SIZE) / 2
	borderRadius := float32(CIRCLE_BORDER_SIZE) / 2

	numLines := 3
	numSlices := numLines * 2

	sliceAngle := 2 * math.Pi / float64(numSlices)

	currentSelection := 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)

		rl.DrawCircle(
			int32(centerX),
			int32(centerY),
			borderRadius,
			rl.White,
		)


		rl.DrawCircle(
			int32(centerX),
			int32(centerY),
			radius,
			CIRCLE_COLOR,
		)

		for i := 0; i < numLines; i++ {
			angle := float64(i) * math.Pi / float64(numLines)

			offsetX := float32(math.Cos(angle)) * radius
			offsetY := float32(math.Sin(angle)) * radius

			startPoint := rl.NewVector2(
				centerX+offsetX,
				centerY+offsetY,
			)

			endPoint := rl.NewVector2(
				centerX-offsetX,
				centerY-offsetY,
			)

			rl.DrawLineEx(
				startPoint,
				endPoint,
				2,
				rl.White,
			)
		}

		textRadius := radius * 0.55

		for i := 0; i < numSlices; i++ {
			theta := (float64(i) + 0.5) * sliceAngle

			textX := centerX +
				float32(math.Cos(theta))*textRadius

			textY := centerY -
				float32(math.Sin(theta))*textRadius

			text := fmt.Sprintf("%d", i+1)

			fontSize := 24
			textWidth := rl.MeasureText(text, int32(fontSize))

			drawX := int32(textX) - textWidth/2
			drawY := int32(textY) - int32(fontSize)/2

			color := rl.RayWhite
			if currentSelection == i {
				color = rl.Red
			}

			rl.DrawText(
				text,
				drawX,
				drawY,
				int32(fontSize),
				color,
			)
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}

		if rl.IsKeyPressed(rl.KeyP) {
			currentSelection = (currentSelection + 1) % numSlices
		} else if rl.IsKeyPressed(rl.KeyN){ 
			currentSelection = (currentSelection - 1 + numSlices) % numSlices

		}

		rl.EndDrawing()
	}
}
