package main

import (
	"encoding/json"
	_ "fmt"
	"math"
	"os"

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

type Hanzi struct {
	Entries map[string][]string `json:"entries"`
}

func loadPinyin() Hanzi {
	var pinyin Hanzi
	content, _ := os.ReadFile("pinyin.json")
	json.Unmarshal(content, &pinyin)

	return pinyin
}

func main() {
	pinyin := loadPinyin()
	wo := []string{"我", "喔", "窝"}
	
	numLines := len(wo) - 1

	rl.SetConfigFlags(
		rl.FlagWindowUndecorated |
		rl.FlagWindowTransparent |
		rl.FlagMsaa4xHint,
	)

	rl.InitWindow(CIRCLE_SIZE, CIRCLE_SIZE, "Sliced Circle")
	defer rl.CloseWindow()

	var allChars []rune
	// Add default characters just in case
	allChars = append(allChars, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ")...)
	// Add every character found in your JSON map
	for _, words := range pinyin.Entries {
		for _, word := range words {
			for _, char := range word {
				allChars = append(allChars, char)
			}
		}
	}

	// Convert your unique Go runes to an int32 slice for Raylib
	var codepoints []int32
	encountered := map[int32]bool{}
	for _, c := range allChars {
		val := int32(c)
		if !encountered[val] {
			encountered[val] = true
			codepoints = append(codepoints, val)
		}
	}

	// Load font with explicit character maps (using 48px base size for clarity)
	font := rl.LoadFontEx("./SimHei.ttf", 48, codepoints, int32(len(codepoints)))
	defer rl.UnloadFont(font)

	rl.SetTargetFPS(60)

	centerX := float32(CIRCLE_SIZE) / 2
	centerY := float32(CIRCLE_SIZE) / 2
	radius := float32(CIRCLE_SIZE) / 2
	borderRadius := float32(CIRCLE_BORDER_SIZE) / 2


	currentSelection := 0
	for !rl.WindowShouldClose() {
		numSlices := numLines * 2
		sliceAngle := 2 * math.Pi / float64(numSlices)

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
		
		for i := range int32(numLines) {
			angle := float64(i) * math.Pi / float64(numLines)

			offsetX := float32(math.Cos(angle)) * radius
			offsetY := float32(math.Sin(angle)) * radius

			startPoint := rl.NewVector2(centerX+offsetX, centerY+offsetY)
			endPoint := rl.NewVector2(centerX-offsetX, centerY-offsetY)

			rl.DrawLineEx(startPoint, endPoint, 2, rl.White)
		}

		textRadius := radius * 0.55

		for i := 0; i < numSlices; i++ {
			theta := (float64(i) + 0.5) * sliceAngle

			textX := centerX + float32(math.Cos(theta))*textRadius
			textY := centerY - float32(math.Sin(theta))*textRadius

			// Map the current slice to a word from the pinyin JSON array
			text := wo[i%len(wo)]

			fontSize := float32(24)
			spacing := float32(2)

			// FIX 1: Convert the UTF-8 string to a slice of codepoints (runes)
			codepoints := []rune(text)
			var intCodepoints []int32
			for _, r := range codepoints {
				intCodepoints = append(intCodepoints, int32(r))
			}

			// FIX 2: Correctly measure Chinese text size using your custom TTF font
			textSize := rl.MeasureTextEx(font, text, fontSize, spacing)

			drawX := textX - textSize.X/2
			drawY := textY - textSize.Y/2

			color := rl.RayWhite
			if currentSelection == i {
				color = rl.Red
			}

			textPosition := rl.Vector2{
				X: drawX,
				Y: drawY,
			}

			// FIX 3: Draw the text using codepoints so Raylib displays Chinese characters
			rl.DrawTextCodepoints(
				font,
				intCodepoints,
				textPosition,
				fontSize,
				spacing,
				color,
			)
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}

		if rl.IsKeyPressed(rl.KeyP) {
			currentSelection = (currentSelection + 1) % numSlices
		} else if rl.IsKeyPressed(rl.KeyN) || rl.IsKeyPressed(rl.KeyTab) { 
			currentSelection = (currentSelection - 1 + numSlices) % numSlices
		}

		rl.EndDrawing()
	}
}

