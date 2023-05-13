package main

import (
	"image/color"
	"log"
)

type board struct {
	matrix       [][]color.RGBA
	x            int
	y            int
	cellSize     int
	isGrayscaled bool
}

type direction int

const (
	down direction = iota
	up
	left
	right
)

func makeMatrix(width, height int) [][]color.RGBA {
	matrix := make([][]color.RGBA, height)
	for i := range matrix {
		matrix[i] = make([]color.RGBA, width)
	}
	return matrix
}

func newBoard(width, height, x, y, cellSize int) board {
	return board{
		makeMatrix(width, height),
		x, y,
		cellSize,
		false,
	}
}

func (b *board) clean() {
	b.matrix = makeMatrix(b.width(), b.height())
	b.isGrayscaled = false
}

func (b board) width() int {
	return len(b.matrix[0])
}

func (b board) height() int {
	return len(b.matrix)
}

func (b board) draw() {
	drawBoardBackground(b, gunMetal.color(), darkGunMetal.color())
	for y, row := range b.matrix {
		for x, cellColor := range row {
			if b.isGrayscaled {
				cellColor = toGrayscale(cellColor)
			}
			drawSquare(b.x+b.cellSize*x, b.y+b.cellSize*y, b.cellSize, cellColor)
		}
	}
}

func (b *board) inBounds(x, y int) bool {
	return y >= 0 && y < b.height() && x >= 0 && x < b.width()
}

func (b *board) isCellFree(x, y int) bool {
	return transparent.equals(b.matrix[y][x])
}

func (b *board) impressSprite(s sprite) {
	for y := 0; y < s.size(); y++ {
		for x := 0; x < s.size(); x++ {
			x1 := ((x * s.cellSize) + s.x - b.x) / s.cellSize
			y1 := ((y * s.cellSize) + s.y - b.y) / s.cellSize
			if b.inBounds(x1, y1) && s.has_block(x, y) && b.isCellFree(x1, y1) {
				b.matrix[y1][x1] = s.colour
			}
		}
	}
}

func (b board) isBlockingSprite(s sprite, d direction) bool {
	for y := 0; y < s.size(); y++ {
		for x := 0; x < s.size(); x++ {
			x1 := ((x * s.cellSize) + s.x - b.x) / s.cellSize
			y1 := ((y * s.cellSize) + s.y - b.y) / s.cellSize
			if b.inBounds(x1, y1) && s.has_block(x, y) {
				if b.isCellFree(x1, y1) {
					x2 := x1
					y2 := y1
					switch d {
					case down:
						y2 += 1
					case left:
						x2 -= 1
					case right:
						x2 += 1
					}
					if !b.inBounds(x2, y2) || !b.isCellFree(x2, y2) {
						return true
					}
				} else {
					return true
				}
			}
		}
	}
	return false
}

func (b board) findCompletedLines() []int {
	var lines []int
	for y := 0; y < b.height(); y++ {
		found := true
		for x := 0; x < b.width(); x++ {
			if b.isCellFree(x, y) {
				found = false
				break
			}
		}
		if found {
			lines = append(lines, y)
		}
	}
	return lines
}

func (b *board) removeLine(line int) {
	if line < 0 || line >= b.height() {
		log.Print("WARN ", "Attempting to remove an non existing line: ", line)
		return
	}
	for x := 0; x < b.width(); x++ {
		b.matrix[line][x] = transparent.color()
	}
	if line > 0 {
		for y := line - 1; y >= 0; y-- {
			for x := 0; x < b.width(); x++ {
				b.matrix[y+1][x] = b.matrix[y][x]
			}
		}
		for x := 0; x < b.width(); x++ {
			b.matrix[0][x] = transparent.color()
		}
	}

}

func (b *board) enabledGrayscale(isEnabled bool) {
	b.isGrayscaled = isEnabled
}
