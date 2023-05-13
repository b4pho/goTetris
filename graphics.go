package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawSquare(x, y int, cellSize int, colour color.RGBA) {
	rl.DrawRectangle(int32(x), int32(y), int32(cellSize), int32(cellSize), colour)
}

func drawBoardBackground(b board, bgColor color.RGBA, borderColor color.RGBA) {
	rl.DrawRectangle(
		int32(b.x), int32(b.y),
		int32(b.cellSize*b.width()),
		int32(b.cellSize*b.height()),
		bgColor,
	)
	rl.DrawRectangleLines(
		int32(b.x), int32(b.y),
		int32(b.cellSize*b.width()),
		int32(b.cellSize*b.height()),
		borderColor,
	)
}

func printCenteredText(text string, fontSize, offsetY int, colour color.RGBA) {
	screenWidth := rl.GetScreenWidth()
	screenHeight := rl.GetScreenHeight()
	textWidth := rl.MeasureText(text, int32(fontSize))
	textHeight := int32(fontSize)
	x := (int32(screenWidth) - textWidth) / 2
	y := (int32(screenHeight)-textHeight)/2 + int32(offsetY)
	rl.DrawText(text, x, y, int32(fontSize), colour)
}

func printScoreText(score, fontSize int, colour color.RGBA) {
	screenWidth := rl.GetScreenWidth()
	label := fmt.Sprintf("Score: %07d   \n", score)
	textWidth := rl.MeasureText(label, int32(fontSize))
	x := int32(screenWidth) - textWidth
	y := int32(fontSize)
	rl.DrawText(label, x, y, int32(fontSize), colour)
}
