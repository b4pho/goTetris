package main

import (
	"strings"
	"testing"
)

const (
	tBLock int = iota
	sBlock
	zBlock
	jBlock
	iBlock
	oBlock
	lBlock
)

const cellSize int = 1

func TestOblockDownToZblock(t *testing.T) {
	// Given a 16x6 board with a Z-block piece on the bottom
	board := newBoard(10, 6, 0, 0, cellSize)
	board.fill([]string{
		"..........",
		"..........",
		"..........",
		"...*......",
		"..**......",
		"..*.......",
	})

	// When a "O-block" is picked
	piece := pickSprite(oBlock, cellSize)

	// Then piece starts at coordinates {0, -4}
	piece.assertPosition(t, 0, -4)

	// And when the piece is moved
	repeatMoves(1, func() { piece.goRight(board) })
	repeatMoves(5, func() { piece.goDown(board) })

	// Then the last movement is blocked ({1, 0} instead of {1, 1} position)
	piece.assertPosition(t, 1, 0)
	piece.assertBlocked(t, down, board)

	// And when board picture is updated
	(&board).impressSprite(piece)

	// Then O-block is placed on top of the Z-block
	board.assertState(t, []string{
		"..........",
		"..vv......",
		"..vv......",
		"...*......",
		"..**......",
		"..*.......",
	})
}

func TestIblockLeftToZblock(t *testing.T) {
	// Given a 16x6 board with a Z-block piece on the bottom
	board := newBoard(10, 6, 0, 0, cellSize)
	board.fill([]string{
		"..........",
		"..........",
		"..........",
		"...*......",
		"..**......",
		"..*.......",
	})

	// When a "I-block" is picked
	piece := pickSprite(iBlock, cellSize)

	// Then piece starts at coordinates {0, -4}
	piece.assertPosition(t, 0, -4)

	// And when piece is moved
	repeatMoves(1, func() { piece.rotateCounterClockwise() })
	repeatMoves(4, func() { piece.goRight(board) })
	repeatMoves(8, func() { piece.goDown(board) })
	repeatMoves(1, func() { piece.goLeft(board) })

	// The last movement is blocked
	piece.assertPosition(t, 3, 4)
	piece.assertBlocked(t, down, board)

	// And when board picture is updated
	(&board).impressSprite(piece)

	// Then I-block is placed next to the Z-block
	board.assertState(t, []string{
		"..........",
		"..........",
		"..........",
		"...*......",
		"..**......",
		"..*vvvv...",
	})
}

func repeatMoves(n int, callback func()) {
	for i := 0; i < n; i++ {
		callback()
	}
}

func (b board) fill(lines []string) {
	for y, line := range lines {
		for x, char := range line {
			if string(char) == "*" {
				b.matrix[y][x] = darkGunMetal.color()
			}
		}
	}
}

func (b board) assertState(t *testing.T, lines []string) {
	lines2 := make([]string, len(lines))
	for y, line := range lines {
		line2 := ""
		for x := range line {
			if darkGunMetal.equals(b.matrix[y][x]) {
				line2 += "*"
			} else if transparent.equals(b.matrix[y][x]) {
				line2 += "."
			} else {
				line2 += "v"
			}
		}
		lines2[y] = line2
	}
	for y, line := range lines {
		if line != lines2[y] {
			t.Fatalf("Assert failed\nExpected:\n%s\n\ngot:\n%s\n\n", strings.Join(lines, "\n"), strings.Join(lines2, "\n"))
		}
	}
}

func (s sprite) assertPosition(t *testing.T, x, y int) {
	if s.x != x || s.y != y {
		t.Fatalf("Expected sprite position (x: %d, y: %d), got: (x: %d, y: %d)", x, y, s.x, s.y)
	}
}

func (s sprite) assertBlocked(t *testing.T, d direction, b board) {
	if !b.isBlockingSprite(s, d) {
		t.Fatalf("Expected sprite blocked on board but wasn't")
	}
}
