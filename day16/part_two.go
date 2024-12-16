package day16

import (
	"io"
)

func PartTwo(r io.Reader) any {
	_, bestSpots := traverseMaze(r)
	return bestSpots
}
