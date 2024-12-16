package day16

import (
	"io"
)

func PartOne(r io.Reader) any {
	bestScore, _ := traverseMaze(r)
	return bestScore
}
