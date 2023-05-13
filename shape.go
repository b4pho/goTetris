package main

import (
	"errors"
	"image/color"
)

type shape struct {
	colour color.RGBA
	matrix [][]byte
}

func chooseShape(i int) (*shape, error) {
	switch i {
	case 0:
		return &shape{
			colour: paleViolet.color(),
			matrix: [][]byte{
				{0, 1, 0},
				{1, 1, 1},
				{0, 0, 0},
			},
		}, nil
	case 1:
		return &shape{
			colour: persianPink.color(),
			matrix: [][]byte{
				{0, 0, 0},
				{0, 1, 1},
				{1, 1, 0},
			},
		}, nil
	case 2:
		return &shape{
			colour: persianPink.color(),
			matrix: [][]byte{
				{0, 0, 0},
				{1, 1, 0},
				{0, 1, 1},
			},
		}, nil
	case 3:
		return &shape{
			colour: lightMalachiteGreen.color(),
			matrix: [][]byte{
				{0, 0, 0},
				{1, 1, 1},
				{0, 0, 1},
			},
		}, nil
	case 4:
		return &shape{
			colour: veryLightTangelo.color(),
			matrix: [][]byte{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 1, 1, 1},
				{0, 0, 0, 0},
			},
		}, nil
	case 5:
		return &shape{
			colour: keyLime.color(),
			matrix: [][]byte{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
		}, nil
	case 6:
		return &shape{
			colour: lightMalachiteGreen.color(),
			matrix: [][]byte{
				{0, 0, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
		}, nil
	default:
		return nil, errors.New("wrong shape index")
	}
}

func (s shape) size() int {
	return len(s.matrix)
}

func (s shape) has_block(x, y int) bool {
	if x >= 0 && x < s.size() && y >= 0 && y < s.size() {
		return s.matrix[x][y] == 1
	} else {
		return false
	}
}

func empty(n int) shape {
	m := make([][]byte, n)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			m[y] = append(m[y], 0)
		}
	}
	return shape{
		colour: color.RGBA{0, 0, 0, 0},
		matrix: m,
	}
}

func (s shape) copyFrom(s2 shape) {
	for i := 0; i < s.size(); i++ {
		for j := 0; j < s.size(); j++ {
			s.matrix[i][j] = s2.matrix[i][j]
		}
	}
}

func (s shape) transpose() {
	s2 := empty(s.size())
	for i := 0; i < s.size(); i++ {
		for j := 0; j < s.size(); j++ {
			s2.matrix[i][j] = s.matrix[j][i]
		}
	}
	s.copyFrom(s2)
}

func (s shape) flipVertically() {
	s2 := empty(s.size())
	for i := 0; i < s.size(); i++ {
		for j := 0; j < s.size(); j++ {
			s2.matrix[i][j] = s.matrix[i][s.size()-1-j]
		}
	}
	s.copyFrom(s2)
}

// See inverse rotation comment below...
// func (s shape) flipHorizontally() {
// 	s2 := empty(s.size())
// 	for i := 0; i < s.size(); i++ {
// 		for j := 0; j < s.size(); j++ {
// 			s2.matrix[i][j] = s.matrix[s.size()-1-i][j]
// 		}
// 	}
// 	s.copyFrom(s2)
// }

func (s shape) draw(x, y, cellSize int) {
	for i := 0; i < s.size(); i++ {
		for j := 0; j < s.size(); j++ {
			if s.matrix[i][j] == 1 {
				drawSquare(x+cellSize*i, y+cellSize*j, cellSize, s.colour)
			}
		}
	}
}

// Maybe could be used in future? inverse rotation? mmm...
// func (s shape) rotateClockwise() {
// 	s.transpose()
// 	s.flipHorizontally()
// }

func (s shape) rotateCounterClockwise() {
	s.transpose()
	s.flipVertically()
}
