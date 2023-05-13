package main

import (
	"log"
)

// From Wikipedia:
// "This system was used in the 1988 BPS versions of Tetris, including the Famicom version.
// This was also used in Super Tetris 3's Standard mode."
func scorePoints(removedLinesCount int) int {
	switch removedLinesCount {
	case 0:
		return 0
	case 1:
		return 40
	case 2:
		return 100
	case 3:
		return 300
	case 4:
		return 1200
	default:
		printScoreWarning(removedLinesCount)
		return 0
	}
}

func printScoreWarning(removedLinesCount int) {
	if removedLinesCount < 0 {
		log.Printf("WARN probably something isn't right here. Erasing %d lines", removedLinesCount)
	} else if removedLinesCount > 4 {
		// That WTF moment:
		log.Fatal("How come did you erase more than 4 rows when the biggest block is 4 rows long???")
	}
}
