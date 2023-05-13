package main

import (
	"log"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type gameState int

const (
	playing gameState = iota
	gameover
)

func randomizeSeed() *rand.Rand {
	seed := rand.NewSource(time.Now().UnixNano())
	return rand.New(seed)
}

func main() {
	state := playing

	random := randomizeSeed()

	rl.InitWindow(800, 600, "Tetris test")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	cellSize := 30
	piece := pickRandomSprite(random, cellSize)
	board := newBoard(10, 20, 0, 0, cellSize)

	timeCount := float32(0.0)
	blinking := false
	score := 0

	for !rl.WindowShouldClose() {

		if state == playing {
			if rl.IsKeyPressed(rl.KeyLeft) {
				piece.goLeft(board)
			} else if rl.IsKeyPressed(rl.KeyRight) {
				piece.goRight(board)
			} else if rl.IsKeyPressed(rl.KeyUp) {
				piece.rotateCounterClockwise()
				if board.isBlockingSprite(piece, right) {
					piece.goLeft(board)
				} else if board.isBlockingSprite(piece, left) {
					piece.goRight(board)
				}
			} else if rl.IsKeyDown(rl.KeyDown) {
				piece.goDown(board)
			}
		} else if state == gameover {
			if rl.IsKeyPressed(rl.KeySpace) {
				(&board).clean()
				state = playing
			}
		}

		frameTime := rl.GetFrameTime()
		timeCount += frameTime

		if timeCount > float32(0.5) {
			timeCount = float32(0.0)
			if state == playing {
				if !piece.goDown(board) {
					(&board).impressSprite(piece)
					piece = pickRandomSprite(random, cellSize)
				}
				completedLines := board.findCompletedLines()
				for _, line := range completedLines {
					board.removeLine(line)
				}
				score += scorePoints(len(completedLines))
				if board.isBlockingSprite(piece, up) {
					(&board).impressSprite(piece)
					state = gameover
					log.Print("INFO ", "Game over!")
				}
			} else if state == gameover {
				// (little) hack: half a second blinking period
				blinking = !blinking
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(outerSpace.color())
		board.draw()
		if state == playing {
			piece.draw()
		} else if state == gameover {
			board.enabledGrayscale(true)
			printCenteredText("GAME OVER!", 50, 0, lightMalachiteGreen.color())
			if blinking {
				printCenteredText("- Press SPACE bar to restart -", 30, 50, persianPink.color())
			}
		}
		printScoreText(score, 20, veryLightTangelo.color())

		rl.EndDrawing()
	}
}
