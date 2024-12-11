package day11

import (
	"io"
)

func PartOne(r io.Reader) any {
	return getNumberOfStones(r, partOneGenerations)
}
