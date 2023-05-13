package main

import(
    "log"
    "math/rand"
)

type sprite struct {
    shape
    x int
    y int
    cellSize int
}

func pickRandomSprite(random *rand.Rand, cellSize int) sprite {
    return pickSprite(random.Intn(7), cellSize)
}

func pickSprite(index, cellSize int) sprite {
    s, err := chooseShape(index)
    if err != nil {
        log.Fatal(err)
    }
    return sprite{*s, 0, -(cellSize * s.size()), cellSize}
}

func (s sprite) draw() {
    s.shape.draw(s.x, s.y, s.cellSize)
}

func (s *sprite) goDown(b board) bool {
    if !b.isBlockingSprite(*s, down) {
        s.y += s.cellSize
        return true
    }
    return false
}

func (s *sprite) goLeft(b board) bool {
    if !b.isBlockingSprite(*s, left) {
        s.x -= s.cellSize
        return true
    }
    return false
}

func (s *sprite) goRight(b board) bool {
    if !b.isBlockingSprite(*s, right) {
        s.x += s.cellSize
        return true
    }
    return false
}
