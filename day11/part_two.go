package day11

import (
	"io"
)

func PartTwo(r io.Reader) any {
	return getNumberOfStones(r, partTwoGenerations)
}
